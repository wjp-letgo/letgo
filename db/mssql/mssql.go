package mssql

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/log"
	"regexp"
	"strings"
	"sync"
)

//全局实现者
var pool MsSqlPooler
var poolLock sync.Mutex

//NewPool 初始化数据库连接
func NewPool(config MsSqlConnect) MsSqlPooler {
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool == nil {
		var configs []MsSqlConnect
		configs = append(configs, config)
		pool = &MsSqlPool{}
		pool.Init(configs)
	}
	return pool
}

//NewPools 初始化多数据库连接
func NewPools(configs []MsSqlConnect) MsSqlPooler {
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool == nil {
		pool = &MsSqlPool{}
		pool.Init(configs)
	}
	return pool
}

//DBer 数据库接口
type DBer interface {
	BeginTransaction()
	Commit()
	Rollback()
	Exec(sql string) int64
	Query(sql string, whereParams ...interface{}) lib.SqlRows
	Desc(tableName string) lib.Columns
	IsExist(tableName string) bool
	ShowTables() []string
	ShowDatabases() []string
}

//DBPool 连接池接口
type DBPool interface {
	SetDB(connectName, databaseName string) DBer
}

//DB 数据库和连接
type DB struct {
	databaseName string
	connectName  string
	dbPool       MsSqlPooler
}

//SetPool 设置连接池
func (db *DB) SetPool(pool MsSqlPooler) DBPool {
	db.dbPool = pool
	return db
}

//SetDB 设置连接名和数据库名称
func (db *DB) SetDB(connectName, databaseName string) DBer {
	db.connectName = connectName
	db.databaseName = databaseName
	return db
}

//Exec 执行原生sql
func (db *DB) Exec(sql string) int64 {
	smt, err := db.prepare(sql)
	if err != nil {
		return -1
	}
	defer smt.Close()
	result, err := smt.Exec()
	if err != nil {
		return -2
	}
	effects, err := result.RowsAffected()
	if err != nil {
		return -3
	}
	return effects
}

//BeginTransaction 开启事务
func (db *DB) BeginTransaction() {
	tx, err := db.dbPool.GetDB(db.connectName).Begin()
	if err != nil {
		return
	}
	db.dbPool.BeginTx(db.connectName)
	db.dbPool.SetTx(db.connectName, tx)
}

//Commit 提交事务
func (db *DB) Commit() {
	db.dbPool.GetTx(db.connectName).Commit()
	db.dbPool.EndTx(db.connectName)
}

//Rollback 事务回滚
func (db *DB) Rollback() {
	db.dbPool.GetTx(db.connectName).Rollback()
	db.dbPool.EndTx(db.connectName)
}

//prepare 执行操作
func (db *DB) prepare(sql string) (*sql.Stmt, error) {
	if db.dbPool.IsTransaction(db.connectName) {
		//开启事务
		return db.dbPool.GetTx(db.connectName).Prepare(sql)
	} else {
		//未使用事务
		return db.dbPool.GetIncludeReadDB(db.connectName).Prepare(sql)
	}
}

//Desc 查询表结构
func (db *DB) Desc(tableName string) lib.Columns {
	sql := fmt.Sprintf(`SELECT 序号 = a.colorder,COLUMN_NAME = a.name,COLUMN_COMMENT = f.value,
标识 = case when COLUMNPROPERTY( a.id,a.name,'IsIdentity') = 1 then '1' else '0' end,
COLUMN_KEY = case when exists(SELECT 1 FROM sysobjects where xtype = 'PK' and parent_obj = a.id and name in (
SELECT name FROM sysindexes WHERE indid in(
SELECT indid FROM sysindexkeys WHERE id = a.id AND colid = a.colid
))) then 'PRI' else '' end,
DATA_TYPE = b.name,
NUMERIC_PRECISION = COLUMNPROPERTY(a.id,a.name,'PRECISION'),
IS_NULLABLE = case when a.isnullable = 1 then 'YES' else 'NO' end,
COLUMN_DEFAULT = isnull(e.text,''),
COLLATION_NAME=a.collation,
NUMERIC_SCALE=a.scale
FROM syscolumns a
left join systypes b on a.xusertype = b.xusertype
inner join sysobjects d on a.id = d.id and d.xtype = 'U' and d.name <> 'dtproperties'
left join syscomments e on a.cdefault = e.id
left join sys.extended_properties f on f.major_id = d.id and f.minor_id = a.colorder
where d.name='%s' `, tableName)
	lst := db.Query(sql)
	var cls lib.Columns
	for _, c := range lst {
		cc := lib.Column{
			Name:          c["COLUMN_NAME"].String(),
			Type:          fmt.Sprintf("%s(%d)", c["DATA_TYPE"].String(), c["NUMERIC_PRECISION"].Int()),
			DataType:      c["DATA_TYPE"].String(),
			Scale:         c["NUMERIC_SCALE"].Int(),
			Key:           c["COLUMN_KEY"].String(),
			IsNull:        c["IS_NULLABLE"].String(),
			Default:       c["COLUMN_DEFAULT"].String(),
			Comment:       c["COLUMN_COMMENT"].String(),
			CollationName: c["COLLATION_NAME"].String(),
		}
		cc.Length = c["NUMERIC_PRECISION"].Int()
		cls = append(cls, cc)
	}
	return cls
}

//ShowDatabases 显示所有数据库
func (db *DB) ShowDatabases() []string {
	var databases []string
	db.Exec("use master")
	tbs := db.Query("select * from sysdatabases where dbid>4")
	for _, t := range tbs {
		databases = append(databases, t["name"].String())
	}
	db.Exec("use " + db.databaseName)
	return databases
}

//ShowTables 显示数据库的所有表名称
func (db *DB) ShowTables() []string {
	var tables []string
	tbs := db.Query("select * from sysobjects where xtype='u'")
	for _, t := range tbs {
		tables = append(tables, t["name"].String())
	}
	return tables
}

//IsExist 检测表是否存在
func (db *DB) IsExist(tableName string) bool {
	sql := fmt.Sprintf("select top 1 * from sysObjects where Id=OBJECT_ID(N'%s') and xtype='U'", tableName)
	tb := db.Query(sql)
	if tb != nil {
		return true
	}
	return false
}

//Query
func (db *DB) Query(sql string, whereParams ...interface{}) lib.SqlRows {
	rows, err := db.query(sql, whereParams...)
	if err != nil {
		return nil
	}
	defer rows.Close()
	return lib.RowsToSqlRows(rows)
}

//query
func (db *DB) query(sql string, whereParams ...interface{}) (*sql.Rows, error) {
	if db.dbPool.IsTransaction(db.connectName) {
		//开启事务
		if len(whereParams) > 0 {
			return db.dbPool.GetTx(db.connectName).Query(sql, whereParams...)
		} else {
			return db.dbPool.GetTx(db.connectName).Query(sql)
		}
	} else {
		//未使用事务
		if len(whereParams) > 0 {
			return db.dbPool.GetIncludeReadDB(db.connectName).Query(sql, whereParams...)
		} else {
			return db.dbPool.GetIncludeReadDB(db.connectName).Query(sql)
		}
	}
}

type Limit struct {
	offset int
	limit  int
}

//Tabler 表操作接口
type Tabler interface {
	SetDB(db *DB)
	GetDB() *DB
	Insert(row lib.SqlIn) int64
	Drop() int64
	Truncate() int64
	Delete(onParams []interface{}, where string, whereParams ...interface{}) int64
	Update(row lib.SqlIn, onParams []interface{}, where string, whereParams ...interface{}) int64
	SelectByHasWhere(fields string, where string, hasWhere bool, page *Limit, whereParams ...interface{}) lib.SqlRows
	Select(fields string, where string, whereParams ...interface{}) lib.SqlRows
	GetLastSql() string
	GetSqlInfo() (string, []interface{})
}

//Table 操作表
type Table struct {
	pri            string
	tableName      string
	firstTableName string
	db             *DB
	lastSql        string
	preSql         string
	preParams      []interface{}
}

//获得最后执行的sql
func (t *Table) GetLastSql() string {
	return t.lastSql
}

//获得最后执行的sql
func (t *Table) GetSqlInfo() (string, []interface{}) {
	return t.preSql, t.preParams
}

//SetDB 设置数据库
func (t *Table) SetDB(db *DB) {
	t.db = db
}

//GetDB 获得数据库
func (t *Table) GetDB() *DB {
	return t.db
}

//Insert 插入操作
func (t *Table) Insert(row lib.SqlIn) int64 {
	return t.add(row, "insert into", "")
}

//add 增加数据通过操作符
func (t *Table) add(row lib.SqlIn, opBefore, opAfter string) int64 {
	var feildsArray []string
	var oldFeildsArray []string
	var valuesArray []string
	var vars []interface{}
	var oldvars []interface{}
	i := 0
	for k, value := range row {
		if strings.Index(k, "[") == -1 {
			feildsArray = append(feildsArray, fmt.Sprintf("[%s]", k))
		} else {
			feildsArray = append(feildsArray, fmt.Sprintf("%s", k))
		}
		oldFeildsArray = append(oldFeildsArray, k)
		valuesArray = append(valuesArray, fmt.Sprintf("@pp%d", i))
		vars = append(vars, sql.Named(fmt.Sprintf("pp%d", i), value))
		oldvars = append(oldvars, value)
		i++
	}
	qsql := fmt.Sprintf("%s %s(%s) values(%s) %s;select ID = convert(bigint, SCOPE_IDENTITY());", opBefore, t.tableName, strings.Join(feildsArray, ","), strings.Join(valuesArray, ","), opAfter)
	smt, err := t.db.prepare(qsql)
	t.sql(qsql, oldvars...)
	if err != nil {
		log.DebugPrint("===========执行TransSql 转原生 sql失败:%s,语句:%s", err.Error(), qsql)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -1
	}
	defer smt.Close()
	result := smt.QueryRow(vars...)
	if result == nil {
		log.DebugPrint("===========执行sql失败,语句:%s", qsql)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -2
	}
	var newID int64
	err = result.Scan(&newID)
	if err != nil {
		log.DebugPrint("===========获取插入ID失败:%s", err.Error())
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -3
	}
	return newID
}

//Update 更新操作
func (t *Table) Update(row lib.SqlIn, onParams []interface{}, where string, whereParams ...interface{}) int64 {
	var setsArray []string
	var vars []interface{}
	var oldvars []interface{}
	tbn := t.tableName
	i := 0
	for key, value := range row {
		if v, ok := value.(lib.SqlRaw); ok {
			setsArray = append(setsArray, fmt.Sprintf("%s=%s", t.getSetField(key), v))
		} else {
			setsArray = append(setsArray, fmt.Sprintf("%s=@pp%d", t.getSetField(key), i))
			vars = append(vars, sql.Named(fmt.Sprintf("pp%d", i), value))
			oldvars = append(oldvars, value)
		}
		i++
	}
	if len(onParams) > 0 {
		reg, _ := regexp.Compile(`@\w+`)
		fs := reg.FindAllString(tbn, -1)
		for k, v := range onParams {
			sk := strings.ReplaceAll(fs[k], "@", "")
			vars = append(vars, sql.Named(sk, v))
			oldvars = append(oldvars, v)
		}
	}
	if len(whereParams) > 0 {
		reg, _ := regexp.Compile(`@\w+`)
		fs := reg.FindAllString(where, -1)
		for k, v := range whereParams {
			sk := strings.ReplaceAll(fs[k], "@", "")
			vars = append(vars, sql.Named(sk, v))
			oldvars = append(oldvars, v)
		}
	}
	sets := strings.Join(setsArray, ",")
	qsql := ""
	if len(onParams) > 0 {
		qsql = fmt.Sprintf("update %s set %s from %s where %s", t.firstTableName, sets, tbn, where)
	} else {
		qsql = fmt.Sprintf("update %s set %s where %s", tbn, sets, where)
	}

	t.sql(qsql, oldvars...)
	smt, err := t.db.prepare(qsql)
	if err != nil {
		log.DebugPrint("update error1 %v", err)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -1
	}
	defer smt.Close()
	result, err := smt.Exec(vars...)
	if err != nil {
		log.DebugPrint("update error2 %v,语句:%s", err, qsql)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -2
	}
	effects, err := result.RowsAffected()
	if err != nil {
		log.DebugPrint("update error3 %v", err)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -3
	}
	return effects
}
func (t *Table) getField(field string) string {
	farr := strings.Split(field, ".")
	if len(farr) == 1 {
		return farr[0]
	} else if len(farr) == 2 {
		return farr[1]
	}
	return field
}

//getSetField
func (t *Table) getSetField(field string) string {
	farr := strings.Split(field, ".")
	if len(farr) == 1 {
		if strings.Index(farr[0], "[") == -1 {
			return "[" + farr[0] + "]"
		}
	} else if len(farr) == 2 {
		if strings.Index(farr[1], "[") == -1 {
			return farr[0] + ".[" + farr[1] + "]"
		}
	}
	return field
}

//Select 查询
func (t *Table) Select(fields string, where string, whereParams ...interface{}) lib.SqlRows {
	return t.SelectByHasWhere(fields, where, true, nil, whereParams...)
}

//SelectByHasWhere 查询
func (t *Table) SelectByHasWhere(fields string, where string, hasWhere bool, page *Limit, whereParams ...interface{}) lib.SqlRows {
	var qsql string
	if where != "" {
		sortstr := "asc"
		if strings.Index(strings.ToLower(where), " desc") == len(where)-5 {
			sortstr = "desc"
		}
		if hasWhere && len(whereParams) > 0 {
			if page != nil {
				tbl := t.tableName
				newWhere := where
				if strings.Index(tbl, " as ") == -1 {
					tbl = t.tableName + " w1"
					if fields == "*" {
						fields = "w1.*"
					}
					newWhere = strings.ReplaceAll(where, "[", "w1.[")
				}

				qsql = fmt.Sprintf(`select %s%s from %s,(
					SELECT TOP %d row_number() OVER (ORDER BY %s %s) n,%s FROM %s where %s
				) w2 where w1.[%s] = w2.[%s] AND w2.n > %d and %s`, "w2.n,",
					fields,
					tbl,
					page.offset+page.limit,
					t.pri,
					sortstr,
					t.pri,
					t.tableName,
					strings.ReplaceAll(where, "@P", "@PPP"),
					t.pri,
					t.pri,
					page.offset,
					newWhere,
				)
				whereParams = append(whereParams, whereParams...)
			} else {
				qsql = fmt.Sprintf("select %s from %s where %s", fields, t.tableName, where)
			}
		} else {
			if page != nil {
				tbl := t.tableName
				newWhere := where
				if strings.Index(tbl, " as ") == -1 {
					tbl = t.tableName + " w1"
					if fields == "*" {
						fields = "w1.*"
					}
					newWhere = strings.ReplaceAll(where, "[", "w1.[")
				}
				qsql = fmt.Sprintf(`select %s from %s,(
				SELECT TOP %d row_number() OVER (ORDER BY %s %s) n,%s FROM %s
				) w2 where w1.[%s] = w2.[%s] AND w2.n > %d and %s`,
					fields,
					tbl,
					page.offset+page.limit,
					t.pri,
					sortstr,
					t.pri,
					t.tableName,
					t.pri,
					t.pri,
					page.offset,
					newWhere,
				)
			} else {
				if len(whereParams) == 0 {
					qsql = fmt.Sprintf("select %s from %s %s", fields, t.tableName, where)
				} else {
					qsql = fmt.Sprintf("select %s from %s where %s", fields, t.tableName, where)
				}

			}
		}
	} else {
		if page != nil {
			tbl := t.tableName
			if strings.Index(tbl, " as ") == -1 {
				tbl = t.tableName + " w1"
				if fields == "*" {
					fields = "w1.*"
				}
			}
			qsql = fmt.Sprintf(`select %s from %s,(
				SELECT TOP %d row_number() OVER (ORDER BY %s asc) n,%s FROM %s
				) w2 where w1.[%s] = w2.[%s] AND w2.n > %d`,
				fields,
				tbl,
				page.offset+page.limit,
				t.pri,
				t.pri,
				t.tableName,
				t.pri,
				t.pri,
				page.offset,
			)
		} else {
			qsql = fmt.Sprintf("select %s from %s", fields, t.tableName)
		}

	}
	rows, err := t.query(qsql, whereParams...)
	if err != nil {
		log.DebugPrint("delete error %v,语句:%s", err, qsql)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return nil
	}
	defer rows.Close()
	return lib.RowsToSqlRows(rows)
}

//Drop 删除表
func (t *Table) Drop() int64 {
	sql := fmt.Sprintf("drop table %s", t.tableName)
	smt, err := t.db.prepare(sql)
	t.sql(sql)
	if err != nil {
		return -1
	}
	defer smt.Close()
	result, err := smt.Exec()
	if err != nil {
		return -2
	}
	effects, err := result.RowsAffected()
	if err != nil {
		return -3
	}
	return effects
}

//Truncate 清空表
func (t *Table) Truncate() int64 {
	sql := fmt.Sprintf("truncate table %s", t.tableName)
	smt, err := t.db.prepare(sql)
	t.sql(sql)
	if err != nil {
		return -1
	}
	defer smt.Close()
	result, err := smt.Exec()
	if err != nil {
		return -2
	}
	effects, err := result.RowsAffected()
	if err != nil {
		return -3
	}
	return effects
}

//Delete 删除表
func (t *Table) Delete(onParams []interface{}, where string, whereParams ...interface{}) int64 {
	var qsql string
	if where != "" {
		qsql = fmt.Sprintf("delete from %s where %s", t.tableName, where)
	} else {
		qsql = fmt.Sprintf("delete from %s", t.tableName)
	}
	var vars []interface{}
	var oldvars []interface{}
	if len(onParams) > 0 {
		reg, _ := regexp.Compile(`@\w+`)
		fs := reg.FindAllString(t.tableName, -1)
		for k, v := range onParams {
			sk := strings.ReplaceAll(fs[k], "@", "")
			vars = append(vars, sql.Named(sk, v))
			oldvars = append(oldvars, v)
		}
	}
	if len(whereParams) > 0 {
		reg, _ := regexp.Compile(`@\w+`)
		fs := reg.FindAllString(where, -1)
		for k, v := range whereParams {
			sk := strings.ReplaceAll(fs[k], "@", "")
			vars = append(vars, sql.Named(sk, v))
			oldvars = append(oldvars, v)
		}
	}
	smt, err := t.db.prepare(qsql)
	t.sql(qsql, oldvars...)
	if err != nil {
		return -1
	}

	defer smt.Close()
	var result driver.Result
	var err2 error
	if where != "" {
		result, err2 = smt.Exec(vars...)
	} else {
		result, err2 = smt.Exec()
	}

	if err2 != nil {
		log.DebugPrint("delete error2 %v,语句:%s", err2, qsql)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -2
	}
	effects, err := result.RowsAffected()
	if err != nil {
		log.DebugPrint("delete error3 %v,语句:%s", err, qsql)
		log.DebugPrint("===========错误sql:%s", t.GetLastSql())
		return -3
	}
	return effects
}

//query 查询
func (t *Table) query(qsql string, whereParams ...interface{}) (*sql.Rows, error) {
	var vars []interface{}
	var oldvars []interface{}
	if len(whereParams) > 0 {
		reg, _ := regexp.Compile(`@\w+`)
		fs := reg.FindAllString(qsql, -1)
		for k, v := range whereParams {
			sk := strings.ReplaceAll(fs[k], "@", "")
			vars = append(vars, sql.Named(sk, v))
			oldvars = append(oldvars, v)
		}
	}
	t.sql(qsql, oldvars...)
	return t.db.query(qsql, vars...)
}

//sql
func (t *Table) sql(sql string, whereParams ...interface{}) {
	t.preSql = sql
	t.preParams = whereParams
	t.lastSql = TransSql(sql, whereParams...)
}

//TransSql 转原生 sql
func TransSql(sql string, whereParams ...interface{}) string {
	if len(whereParams) > 0 {
		var params []interface{}
		for _, v := range whereParams {
			params = append(params, lib.InterfaceToString(v))
		}
		reg, _ := regexp.Compile(`@\w+`)
		fs := reg.FindAllString(sql, -1)
		for _, v := range fs {
			sql = strings.ReplaceAll(sql, v, "%s")
		}
		sql = fmt.Sprintf(sql, params...)
		return sql
	} else {
		return sql
	}
}

//NewTable 初始化表
func NewTable(db *DB, tableName, firstTableName, pri string) Tabler {
	var table *Table = &Table{}
	table.tableName = tableName
	table.firstTableName = firstTableName
	table.SetDB(db)
	if pri == "" {
		cols := table.db.Desc(tableName)
		for _, v := range cols {
			if v.Key == "PRI" {
				table.pri = v.Name
			}
		}
	} else {
		table.pri = pri
	}
	return table
}

//Connect 连接到数据库
func Connect(connectName, databaseName string) *DB {
	db := &DB{}
	db.SetPool(pool)
	db.SetDB(connectName, databaseName)
	return db
}

//NewDB 新建数据库连接包括连接池
func NewDB() *DB {
	dbFile := "config/db.config"
	cfgFile := file.GetContent(dbFile)
	var configs []MsSqlConnect
	if cfgFile == "" {
		var slaves []DBConfig = make([]DBConfig, 1)
		master := DBConfig{
			Name:         "connectName",
			DatabaseName: "databaseName",
			UserName:     "userName",
			Password:     "password",
			Host:         "127.0.0.1",
			Port:         "1433",
			Charset:      "utf8mb4",
			MaxOpenConns: 20,
			MaxIdleConns: 10,
		}
		configs = append(configs, MsSqlConnect{
			Master: master,
			Slave:  slaves,
		})
		file.PutContent(dbFile, fmt.Sprintf("%v", configs))
		log.PanicPrint("please setting database config in config/db.config file")
	}
	lib.StringToObject(cfgFile, &configs)
	var db DB
	db.SetPool(NewPools(configs))
	return &db
}
