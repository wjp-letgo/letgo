package mysql

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wjp-letgo/letgo/file"
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/log"
)

//全局实现者
var pool MysqlPooler
var poolLock sync.Mutex

//NewPool 初始化数据库连接
func NewPool(config MysqlConnect) MysqlPooler{
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool==nil{
		var configs []MysqlConnect
		configs=append(configs, config)
		pool=&MysqlPool{}
		pool.Init(configs)
	}
	return pool;
}
//NewPools 初始化多数据库连接
func NewPools(configs []MysqlConnect) MysqlPooler{
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool==nil{
		pool=&MysqlPool{}
		pool.Init(configs)
	}
	return pool;
}
//DBer 数据库接口
type DBer interface{
	BeginTransaction()
	Commit()
	Rollback()
	Exec(sql string)int64
	Query(sql string,whereParams ...interface{})lib.SqlRows
	Desc(tableName string)lib.Columns
	IsExist(tableName string) bool
	ShowTables()[]string
	ShowCreateTable(tableName string)string
}
//DBPool 连接池接口
type DBPool interface{
	SetDB(connectName,databaseName string)DBer
}
//DB 数据库和连接
type DB struct{
	databaseName string
	connectName string
	dbPool MysqlPooler
}

//SetPool 设置连接池
func (db *DB)SetPool(pool MysqlPooler)DBPool{
	db.dbPool=pool
	return db
}
//SetDB 设置连接名和数据库名称
func (db *DB)SetDB(connectName,databaseName string)DBer{
	db.connectName=connectName
	db.databaseName=databaseName
	return db
}
//Exec 执行原生sql
func (db *DB)Exec(sql string)int64 {
	smt,err:=db.prepare(sql)
	if err!=nil{
		return -1
	}
	defer smt.Close()
	result, err :=smt.Exec()
	if err!=nil{
		return -2
	}
	effects, err := result.RowsAffected()
	if err!=nil{
		return -3
	}
	return effects
}
//BeginTransaction 开启事务
func (db *DB)BeginTransaction(){
	tx,err:=db.dbPool.GetDB(db.connectName).Begin()
	if err!=nil{
		return
	}
	db.dbPool.BeginTx(db.connectName)
	db.dbPool.SetTx(db.connectName,tx)
}
//Commit 提交事务
func (db *DB)Commit(){
	db.dbPool.GetTx(db.connectName).Commit()
	db.dbPool.EndTx(db.connectName)
}
//Rollback 事务回滚
func (db *DB)Rollback(){
	db.dbPool.GetTx(db.connectName).Rollback()
	db.dbPool.EndTx(db.connectName)
}
//prepare 执行操作
func (db *DB)prepare(sql string)(*sql.Stmt, error){
	if db.dbPool.IsTransaction(db.connectName) {
		//开启事务
		return db.dbPool.GetTx(db.connectName).Prepare(sql)
	}else{
		//未使用事务
		return db.dbPool.GetIncludeReadDB(db.connectName).Prepare(sql)
	}
}

func (db *DB)ShowCreateTable(tableName string)string{
	sql:=fmt.Sprintf("show create table `%s`.`%s`", db.databaseName,tableName)
	lst:= db.Query(sql)
	if len(lst)>0{
		return lst[0]["Create Table"].String()
	}
	return ""
}

//Desc 查询表结构
func (db *DB)Desc(tableName string)lib.Columns{
	sql:=fmt.Sprintf("select * from information_schema.columns where table_schema = '%s' and table_name = '%s'", db.databaseName,tableName)
	lst:= db.Query(sql)
	var cls lib.Columns
	for _,c:=range lst{
		cc:= lib.Column{
			Name: c["COLUMN_NAME"].String(),
			Type: c["COLUMN_TYPE"].String(),
			DataType: c["DATA_TYPE"].String(),
			Scale: c["NUMERIC_SCALE"].Int(),
			Extra: c["EXTRA"].String(),
			Key: c["COLUMN_KEY"].String(),
			IsNull: c["IS_NULLABLE"].String(),
			Default: c["COLUMN_DEFAULT"].String(),
			Comment:c["COLUMN_COMMENT"].String(),
			CharacterSetName:c["CHARACTER_SET_NAME"].String(),
			CollationName:c["COLLATION_NAME"].String(),
		}
		nums:=[]string{"int","bigint","bit","tinyint","decimal","double","float","integer","mediumint","numeric","smallint"}
		dates:=[]string{"datetime","timestamp","time"}
		if lib.InStringArray(c["DATA_TYPE"].String(),nums){
			cc.Length=c["NUMERIC_PRECISION"].Int()
		}else if lib.InStringArray(c["DATA_TYPE"].String(),dates){
			cc.Length=c["DATETIME_PRECISION"].Int()
		}else{
			cc.Length=c["CHARACTER_MAXIMUM_LENGTH"].Int()
		}
		cls=append(cls, cc)
	}
	return cls
}
//ShowTables 显示数据库的所有表名称
func (db *DB)ShowTables()[]string{
	var tables []string
	tbs:=db.Query("show tables")
	k:=fmt.Sprintf("Tables_in_%s",db.databaseName)
	for _,t:=range tbs{
		tables=append(tables, t[k].String())
	}
	return tables
}
//IsExist 检测表是否存在
func (db *DB)IsExist(tableName string) bool{
	sql:=fmt.Sprintf("select * from information_schema.tables where table_schema = '%s' and table_name ='%s';",db.databaseName,tableName)
	tb:= db.Query(sql)
	if tb!=nil{
		return true
	}
	return false
}

//Query
func (db *DB)Query(sql string,whereParams ...interface{})lib.SqlRows{
	rows,err:=db.query(sql,whereParams...)
	if err!=nil{
		return nil
	}
	defer rows.Close()
	return lib.RowsToSqlRows(rows)
}
//query
func (db *DB)query(sql string,whereParams ...interface{})(*sql.Rows, error){
	if db.dbPool.IsTransaction(db.connectName) {
		//开启事务
		if len(whereParams)>0{
			return db.dbPool.GetTx(db.connectName).Query(sql,whereParams...)
		}else{
			return db.dbPool.GetTx(db.connectName).Query(sql)
		}
	}else{
		//未使用事务
		if len(whereParams)>0{
			return db.dbPool.GetIncludeReadDB(db.connectName).Query(sql,whereParams...)
		}else{
			return db.dbPool.GetIncludeReadDB(db.connectName).Query(sql)
		}
	}
}
//Tabler 表操作接口
type Tabler interface{
	SetDB(db *DB)
	GetDB()*DB
	Insert(row lib.SqlIn) int64
	Replace(row lib.SqlIn) int64
	InsertOnDuplicate(row lib.SqlIn,updateRow lib.SqlIn) int64
	Drop() int64
	Truncate() int64
	Optimize() int64
	Delete(onParams []interface{},where string,whereParams ...interface{}) int64
	Update(row lib.SqlIn,onParams []interface{},where string,whereParams ...interface{})int64
	SelectByHasWhere(fields string, where string,hasWhere bool, whereParams ...interface{})lib.SqlRows
	Select(fields string, where string, whereParams ...interface{})lib.SqlRows
	GetLastSql()string
	GetSqlInfo()(string,[]interface{})
}
//Table 操作表
type Table struct{
	tableName string
	db *DB
	lastSql string
	preSql string
	preParams []interface{}
}
//获得最后执行的sql
func (t *Table) GetLastSql()string{
	return t.lastSql
}
//获得最后执行的sql
func (t *Table) GetSqlInfo()(string,[]interface{}){
	return t.preSql,t.preParams
}
//SetDB 设置数据库
func (t *Table) SetDB(db *DB){
	t.db=db
}
//GetDB 获得数据库
func (t *Table) GetDB()*DB{
	return t.db
}
//Insert 插入操作
func (t *Table) Insert(row lib.SqlIn) int64{
	return t.add(row, "insert into", "")
}
//Replace 插入操作
func (t *Table) Replace(row lib.SqlIn) int64{
	return t.add(row, "replace into", "")
}
//InsertOnDuplicate 如果你插入的记录导致一个UNIQUE索引或者primary key(主键)出现重复，那么就会认为该条记录存在，则执行update语句而不是insert语句，反之，则执行insert语句而不是更新语句。
func (t *Table) InsertOnDuplicate(row lib.SqlIn,updateRow lib.SqlIn) int64{
	var feildsArray []string
	var valuesArray []string
	var setsArray []string
	var vars []interface{}
	for k,value:=range row{
		feildsArray=append(feildsArray,fmt.Sprintf("`%s`",k))
		if v, ok := value.(lib.SqlRaw); ok {
			valuesArray=append(valuesArray, fmt.Sprintf("%s",v))
		}else{
			valuesArray=append(valuesArray,"?")
			vars=append(vars,value)
		}
	}
	for k,value:=range updateRow{
		if v, ok := value.(lib.SqlRaw); ok {
			setsArray=append(setsArray, fmt.Sprintf("`%s`=%s",k,v))
		}else{
			setsArray=append(setsArray, fmt.Sprintf("`%s`=?",k))
			vars=append(vars,value)
		}
	}
	sql:=fmt.Sprintf("insert into %s(%s) values(%s) ON DUPLICATE KEY UPDATE %s", t.tableName, strings.Join(feildsArray,","),strings.Join(valuesArray,","),strings.Join(setsArray,","))
	smt,err:=t.db.prepare(sql)
	t.sql(sql,vars...)
	if err!=nil{
		return -1
	}
	defer smt.Close()
	result, err :=smt.Exec(vars...)
	if err!=nil{
		return -2
	}
	effects, err := result.LastInsertId()
	if err!=nil{
		return -3
	}
	return effects
}
//add 增加数据通过操作符
func (t *Table)add(row lib.SqlIn,opBefore,opAfter string)int64{
	var feildsArray []string
	var valuesArray []string
	var vars []interface{}
	for k,value:=range row{
		if strings.Index(k,"`")==-1{
			feildsArray=append(feildsArray,fmt.Sprintf("`%s`",k))
		}else{
			feildsArray=append(feildsArray,fmt.Sprintf("%s",k))
		}
		valuesArray=append(valuesArray,"?")
		vars=append(vars,value)
	}
	sql:=fmt.Sprintf("%s %s(%s) values(%s) %s",opBefore, t.tableName, strings.Join(feildsArray,","),strings.Join(valuesArray,","),opAfter)
	smt,err:=t.db.prepare(sql)
	t.sql(sql,vars...)
	if err!=nil{
		log.DebugPrint("===========执行TransSql 转原生 sql失败:%s",err.Error())
		log.DebugPrint("===========错误sql:%s",t.GetLastSql())
		return -1
	}
	defer smt.Close()
	result, err :=smt.Exec(vars...)
	if err!=nil{
		log.DebugPrint("===========执行sql失败:%s",err.Error())
		log.DebugPrint("===========错误sql:%s",t.GetLastSql())
		return -2
	}
	effects, err := result.LastInsertId()
	if err!=nil{
		log.DebugPrint("===========获取插入ID失败:%s",err.Error())
		log.DebugPrint("===========错误sql:%s",t.GetLastSql())
		return -3
	}
	return effects
}
//Update 更新操作
func (t *Table) Update(row lib.SqlIn,onParams []interface{},where string,whereParams ...interface{})int64{
	var setsArray []string
	var vars []interface{}
	tbn:=t.tableName
	if len(onParams)>0{
		/*
		for i,v:=range onParams{
			s:=lib.InterfaceToString(v)
			regex,_:=regexp.Compile("\\.`[\\s\\S]+?`")
			if regex.MatchString(s){
				tbn=lib.ReplaceIndex(tbn,"?", s,i)
			}else{
				vars=append(vars, v)
			}
			//tbn=strings.Replace(tbn,"?", s,1)
		}
		*/
		//log.DebugPrint(tbn)
		vars=append(vars, onParams...)
	}
	for key,value:=range row {
		if v, ok := value.(lib.SqlRaw); ok {
			setsArray=append(setsArray, fmt.Sprintf("%s=%s",t.getSetField(key),v))
		}else{
			setsArray=append(setsArray, fmt.Sprintf("%s=?",t.getSetField(key)))
			vars=append(vars, value)
		}
		
	}
	vars=append(vars,whereParams...)
	sets:=strings.Join(setsArray,",")
	sql:=fmt.Sprintf("update %s set %s where %s", tbn,sets,where)
	t.sql(sql,vars...)
	smt,err:=t.db.prepare(sql)
	if err!=nil{
		log.DebugPrint("update error1 %v", err)
		log.DebugPrint("===========错误sql:%s",t.GetLastSql())
		return -1
	}
	defer smt.Close()
	result, err :=smt.Exec(vars...)
	if err!=nil{
		log.DebugPrint("update error2 %v", err)
		log.DebugPrint("===========错误sql:%s",t.GetLastSql())
		return -2
	}
	effects, err := result.RowsAffected()
	if err!=nil{
		log.DebugPrint("update error3 %v", err)
		log.DebugPrint("===========错误sql:%s",t.GetLastSql())
		return -3
	}
	return effects
}
//getSetField
func (t *Table) getSetField(field string)string{
	farr:=strings.Split(field,".")
	if len(farr)==1{
		if strings.Index(farr[0],"`")==-1{
			return "`"+farr[0]+"`"
		}
	}else if len(farr)==2{
		if strings.Index(farr[1],"`")==-1{
			return farr[0]+".`"+farr[1]+"`"
		}
	}
	return field
}
//Select 查询
func (t *Table) Select(fields string, where string, whereParams ...interface{})lib.SqlRows{
	return t.SelectByHasWhere(fields,where, true,whereParams...)
}
//SelectByHasWhere 查询
func (t *Table) SelectByHasWhere(fields string, where string,hasWhere bool, whereParams ...interface{})lib.SqlRows{
	var sql string
	if where!=""{
		if hasWhere&&len(whereParams)>0{
			sql=fmt.Sprintf("select %s from %s where %s", fields,t.tableName, where)
		}else{
			if len(whereParams)==0{
				sql=fmt.Sprintf("select %s from %s %s", fields,t.tableName, where)
			}else{
				sql=fmt.Sprintf("select %s from %s where %s", fields,t.tableName, where)
			}
		}
	}else{
		sql=fmt.Sprintf("select %s from %s", fields,t.tableName)
	}
	rows,err:=t.query(sql,whereParams...)
	if err!=nil{
		return nil
	}
	defer rows.Close()
	return lib.RowsToSqlRows(rows)
}
//Drop 删除表
func (t *Table)Drop() int64 {
	sql:=fmt.Sprintf("drop table %s",t.tableName)
	smt,err:=t.db.prepare(sql)
	t.sql(sql)
	if err!=nil{
		return -1
	}
	defer smt.Close()
	result, err :=smt.Exec()
	if err!=nil{
		return -2
	}
	effects, err := result.RowsAffected()
	if err!=nil{
		return -3
	}
	return effects
}
//Optimize
func (t *Table)Optimize() int64{
	sql:=fmt.Sprintf("optimize table %s",t.tableName)
	smt,err:=t.db.prepare(sql)
	t.sql(sql)
	if err!=nil{
		return -1
	}
	defer smt.Close()
	result, err :=smt.Exec()
	if err!=nil{
		return -2
	}
	effects, err := result.RowsAffected()
	if err!=nil{
		return -3
	}
	return effects
}

//Truncate 清空表
func (t *Table)Truncate() int64 {
	sql:=fmt.Sprintf("truncate %s",t.tableName)
	smt,err:=t.db.prepare(sql)
	t.sql(sql)
	if err!=nil{
		return -1
	}
	defer smt.Close()
	result, err :=smt.Exec()
	if err!=nil{
		return -2
	}
	effects, err := result.RowsAffected()
	if err!=nil{
		return -3
	}
	return effects
}
//Delete 删除表
func (t *Table)Delete(onParams []interface{},where string,whereParams ...interface{}) int64 {
	var sql string
	if where!="" {
		sql=fmt.Sprintf("delete from %s where %s",t.tableName, where)
	} else {
		sql=fmt.Sprintf("delete from %s",t.tableName)
	}
	var vars []interface{}
	if len(onParams)>0{
		vars=append(vars, onParams...)
	}
	if len(whereParams)>0{
		vars=append(vars, whereParams...)
	}
	smt,err:=t.db.prepare(sql)
	t.sql(sql,vars...)
	if err!=nil{
		return -1
	}
	
	defer smt.Close()
	var result driver.Result
	var err2 error
	if where!="" {
		result, err2 =smt.Exec(vars...)
	}else{
		result, err2 =smt.Exec()
	}
	
	if err2!=nil{
		return -2
	}
	effects, err := result.RowsAffected()
	if err!=nil{
		return -3
	}
	return effects
}
//query 查询
func (t *Table)query(sql string,whereParams ...interface{})(*sql.Rows, error){
	t.sql(sql,whereParams...)
	return t.db.query(sql, whereParams...)
}

//sql
func (t *Table)sql(sql string,whereParams ...interface{}){
	t.preSql=sql
	t.preParams=whereParams
	t.lastSql=TransSql(sql,whereParams...)
}
//TransSql 转原生 sql
func TransSql(sql string,whereParams ...interface{}) string{
	if len(whereParams) > 0 {
		var params []interface{}
		for _,v:=range whereParams{
			params=append(params, lib.InterfaceToString(v))
		}
		sql=strings.ReplaceAll(sql,"?","%s")
		sql=fmt.Sprintf(sql,params...)
		return sql
	}else{
		return sql
	}
}
//NewTable 初始化表
func NewTable(db *DB,tableName string) Tabler{
	var table *Table=&Table{}
	table.tableName=tableName
	table.SetDB(db)
	return table
}

//Connect 连接到数据库
func Connect(connectName,databaseName string)*DB{
	db:=&DB{}
	db.SetPool(pool)
	db.SetDB(connectName, databaseName)
	return db
}

//NewDB 新建数据库连接包括连接池
func NewDB()*DB{
	dbFile:="config/db.config"
	cfgFile:=file.GetContent(dbFile)
	var configs []MysqlConnect
	if cfgFile==""{
		var slaves []SlaveDB=make([]SlaveDB, 1)
		master:=SlaveDB{
				Name:"connectName",
				DatabaseName:"databaseName",
				UserName:"userName",
				Password:"password",
				Host:"127.0.0.1",
				Port:"3306",
				Charset:"utf8mb4",
				MaxOpenConns:20,
				MaxIdleConns:10,
		}
		configs=append(configs, MysqlConnect{
			Master:master,
			Slave:slaves,
		})
		file.PutContent(dbFile,fmt.Sprintf("%v",configs))
		log.PanicPrint("please setting database config in config/db.config file")
	}
	lib.StringToObject(cfgFile, &configs)
	var db DB
	db.SetPool(NewPools(configs))
	return &db
}
