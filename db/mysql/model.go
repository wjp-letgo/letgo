package mysql

import (
	"github.com/wjp-letgo/letgo/lib"
	//"github.com/wjp-letgo/letgo/log"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// db 全局变量
var db *DB

// Model 模型
type Model struct {
	tableName      string
	oldTableName   string
	aliasName      string
	otherTableName []joinCond
	dbName         string
	fields         []string
	where          []cond
	groupBy        []string
	having         []cond
	orderBy        string
	offset         int
	limit          int
	lastSql        string
	preSql         string
	preParams      []interface{}
	unionModel     Modeler
	db             *DB
	SoftDelete     bool
	DeleteName     string
	tmpSoftDelete  int
}

// cond 操作
type cond struct {
	field  string
	symbol string
	logic  string
	value  interface{}
}

// joinCond 连接条件
type joinCond struct {
	tableName string
	on        []cond
}

// Modeler 模型接口
type Modeler interface {
	Fields(fields ...string) Fielder
	Count() int64
	GetLastSql() string
	GetSqlInfo() (string, []interface{})
	GetModelSql() (string, []interface{})
	Alias(name string) Aliaser
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	Union(model Modeler) Unioner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value interface{}) Wherer
	WhereNotIn(field string, value interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Get() lib.SqlRows
	Find() lib.SqlRow
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
	Update(data lib.SqlIn) int64
	Insert(row lib.SqlIn) int64
	Create(row lib.SqlIn) int64
	Replace(row lib.SqlIn) int64
	InsertOnDuplicate(row lib.SqlIn, updateRow lib.SqlIn) int64
	Drop() int64
	Truncate() int64
	Optimize() int64
	Delete() int64
	SoftDel() int64
	IgnoreSoftDel()
	DB() DBer
}

// Fielder 字段暴露出去的接口
type Fielder interface {
	Alias(name string) Aliaser
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	Union(model Modeler) Unioner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value interface{}) Wherer
	WhereNotIn(field string, value interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Count() int64
	Get() lib.SqlRows
	Find() lib.SqlRow
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
}

// Aliaser 另外取名接口
type Aliaser interface {
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	Union(model Modeler) Unioner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value interface{}) Wherer
	WhereNotIn(field string, value interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Count() int64
	Get() lib.SqlRows
	Find() lib.SqlRow
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
	Update(data lib.SqlIn) int64
}

// Joiner 连接接口
type Joiner interface {
	On(field string, value interface{}) Oner
	OnRaw(on string) Oner
	OnSymbol(field, symbol string, value interface{}) Oner
	OnIn(field string, value interface{}) Oner
	OnNotIn(field string, value interface{}) Oner
}

// Oner on连接条件
type Oner interface {
	OrOn(field string, value interface{}) Oner
	OrOnRaw(on string) Oner
	OrOnSymbol(field, symbol string, value interface{}) Oner
	OrOnIn(field string, value interface{}) Oner
	OrOnNotIn(field string, value interface{}) Oner
	AndOn(field string, value interface{}) Oner
	AndOnRaw(on string) Oner
	AndOnSymbol(field, symbol string, value interface{}) Oner
	AndOnIn(field string, value interface{}) Oner
	AndOnNotIn(field string, value interface{}) Oner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value interface{}) Wherer
	WhereNotIn(field string, value interface{}) Wherer
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	GroupBy(field ...string) GroupByer
	Count() int64
	Get() lib.SqlRows
	Find() lib.SqlRow
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
	Update(data lib.SqlIn) int64
}

// Wherer 条件接口
type Wherer interface {
	AndWhere(field string, value interface{}) Wherer
	AndWhereRaw(where string) Wherer
	AndWhereSymbol(field, symbol string, value interface{}) Wherer
	AndWhereIn(field string, value interface{}) Wherer
	AndWhereNotIn(field string, value interface{}) Wherer
	OrWhere(field string, value interface{}) Wherer
	OrWhereRaw(where string) Wherer
	OrWhereSymbol(field, symbol string, value interface{}) Wherer
	OrWhereIn(field string, value interface{}) Wherer
	OrWhereNotIn(field string, value interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Get() lib.SqlRows
	Find() lib.SqlRow
	Count() int64
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
	Update(data lib.SqlIn) int64
	Delete() int64
	SoftDel() int64
}

// GroupByer 分组接口
type GroupByer interface {
	Having(field string, value interface{}) Havinger
	HavingRaw(having string) Havinger
	HavingSymbol(field, symbol string, value interface{}) Havinger
	HavingIn(field string, value interface{}) Havinger
	HavingNotIn(field string, value interface{}) Havinger
	Get() lib.SqlRows
	Find() lib.SqlRow
	Count() int64
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
}

// Havinger having条件
type Havinger interface {
	AndHaving(field string, value interface{}) Havinger
	AndHavingRaw(having string) Havinger
	AndHavingSymbol(field, symbol string, value interface{}) Havinger
	AndHavingIn(field string, value interface{}) Havinger
	AndHavingNotIn(field string, value interface{}) Havinger
	OrHaving(field string, value interface{}) Havinger
	OrHavingRaw(having string) Havinger
	OrHavingSymbol(field, symbol string, value interface{}) Havinger
	OrHavingIn(field string, value interface{}) Havinger
	OrHavingNotIn(field string, value interface{}) Havinger
	Get() lib.SqlRows
	Find() lib.SqlRow
	Count() int64
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
}

// OrderByer 排序
type OrderByer interface {
	Get() lib.SqlRows
	Find() lib.SqlRow
	Limit(offset ...int) Limiter
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
}

// Limiter
type Limiter interface {
	Get() lib.SqlRows
	Find() lib.SqlRow
	Update(data lib.SqlIn) int64
	Delete() int64
}

// Unioner
type Unioner interface {
	Count() int64
	Get() lib.SqlRows
	Find() lib.SqlRow
	OrderBy(orderBy string) OrderByer
	Limit(offset ...int) Limiter
	Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow)
}

// Init 初始化
func (m *Model) Init(dbName, tableName string) Modeler {
	m.tableName = tableName
	m.oldTableName = tableName
	m.dbName = dbName
	m.db = Connect(m.dbName, m.dbName)
	return m
}

// Init 初始化
func (m *Model) InitByConnectName(connectName, dbName, tableName string) Modeler {
	m.tableName = tableName
	m.oldTableName = tableName
	m.dbName = dbName
	m.db = Connect(connectName, m.dbName)
	return m
}

// DBer 获得数据库接口
func (m *Model) DB() DBer {
	return m.db
}

// Fields 查询字段
func (m *Model) Fields(fields ...string) Fielder {
	m.fields = fields
	return m
}

// 忽略软删除条件
func (m *Model) IgnoreSoftDel() {
	if m.SoftDelete {
		m.tmpSoftDelete = 1
	} else {
		m.tmpSoftDelete = 2
	}
	m.SoftDelete = false
}

// 还原软删除
func (m *Model) resetSoftDel() {
	if m.tmpSoftDelete == 1 {
		m.SoftDelete = true
		m.tmpSoftDelete = 0
	} else if m.tmpSoftDelete == 2 {
		m.SoftDelete = false
		m.tmpSoftDelete = 0
	}
}

// GetLastSql 获得最后执行的sql
func (m *Model) GetLastSql() string {
	return m.lastSql
}

// GetSqlInfo 获得最后执行的sql
func (m *Model) GetSqlInfo() (string, []interface{}) {
	return m.preSql, m.preParams
}

// Alias 命名
func (m *Model) Alias(name string) Aliaser {
	m.aliasName = name
	if !strings.Contains(m.tableName, " as ") {
		m.tableName = fmt.Sprintf("%s as %s", m.tableName, name)
	}
	return m
}

// Join 连接查询
func (m *Model) Join(tableName string) Joiner {
	m.otherTableName = append(
		m.otherTableName,
		joinCond{
			tableName: fmt.Sprintf("INNER JOIN %s", tableName),
		},
	)
	return m
}

// LeftJoin 连接查询
func (m *Model) LeftJoin(tableName string) Joiner {
	m.otherTableName = append(
		m.otherTableName,
		joinCond{
			tableName: fmt.Sprintf("LEFT JOIN %s", tableName),
		},
	)
	return m
}

// RightJoin 连接查询
func (m *Model) RightJoin(tableName string) Joiner {
	m.otherTableName = append(
		m.otherTableName,
		joinCond{
			tableName: fmt.Sprintf("RIGHT JOIN %s", tableName),
		},
	)
	return m
}

// setJoinCond 设置连接条件
func (m *Model) setJoinCond(field, symbol, logic string, value interface{}) {
	index := len(m.otherTableName) - 1
	if index >= 0 {
		var icond cond
		if len(m.otherTableName[index].on) == 0 {
			icond = cond{
				field:  field,
				symbol: symbol,
				logic:  "",
				value:  value,
			}
		} else {
			icond = cond{
				field:  field,
				symbol: symbol,
				logic:  logic,
				value:  value,
			}
		}
		m.otherTableName[index].on = append(
			m.otherTableName[index].on,
			icond,
		)
	}
}

// setWhereCond 设置条件
func (m *Model) setWhereCond(field, symbol, logic string, value interface{}) {
	var icond cond
	if len(m.where) == 0 {
		icond = cond{
			field:  field,
			symbol: symbol,
			logic:  "",
			value:  value,
		}
	} else {
		icond = cond{
			field:  field,
			symbol: symbol,
			logic:  logic,
			value:  value,
		}
	}
	m.where = append(m.where, icond)
}

// setHavingCond 设置条件
func (m *Model) setHavingCond(field, symbol, logic string, value interface{}) {
	var icond cond
	if len(m.having) == 0 {
		icond = cond{
			field:  field,
			symbol: symbol,
			logic:  "",
			value:  value,
		}
	} else {
		icond = cond{
			field:  field,
			symbol: symbol,
			logic:  logic,
			value:  value,
		}
	}
	m.having = append(m.having, icond)
}

// On Join 条件on
func (m *Model) On(field string, value interface{}) Oner {
	m.setJoinCond(field, "=", "and", value)
	return m
}

// OnRaw Join 条件on
func (m *Model) OnRaw(on string) Oner {
	m.setJoinCond(fmt.Sprintf("(%s)", on), "", "and", nil)
	return m
}

// OnSymbol Join 条件on
func (m *Model) OnSymbol(field, symbol string, value interface{}) Oner {
	m.setJoinCond(field, symbol, "and", value)
	return m
}

// OnIn Join 条件on
func (m *Model) OnIn(field string, value interface{}) Oner {
	m.setJoinCond(field, "in", "and", value)
	return m
}

// OnNotIn Join 条件on
func (m *Model) OnNotIn(field string, value interface{}) Oner {
	m.setJoinCond(field, "not in", "and", value)
	return m
}

// OrOn Join 条件on
func (m *Model) OrOn(field string, value interface{}) Oner {
	m.setJoinCond(field, "=", "or", value)
	return m
}

// OrOnRaw Join 条件on
func (m *Model) OrOnRaw(on string) Oner {
	m.setJoinCond(fmt.Sprintf("(%s)", on), "", "or", nil)
	return m
}

// OrOnSymbol Join 条件on
func (m *Model) OrOnSymbol(field, symbol string, value interface{}) Oner {
	m.setJoinCond(field, symbol, "or", value)
	return m
}

// OrOnIn Join 条件on
func (m *Model) OrOnIn(field string, value interface{}) Oner {
	m.setJoinCond(field, "in", "or", value)
	return m
}

// OrOnNotIn Join 条件on
func (m *Model) OrOnNotIn(field string, value interface{}) Oner {
	m.setJoinCond(field, "not in", "or", value)
	return m
}

// AndOn Join 条件on
func (m *Model) AndOn(field string, value interface{}) Oner {
	m.setJoinCond(field, "=", "and", value)
	return m
}

// AndOnRaw Join 条件on
func (m *Model) AndOnRaw(on string) Oner {
	m.setJoinCond(fmt.Sprintf("(%s)", on), "", "and", nil)
	return m
}

// AndOnSymbol Join 条件on
func (m *Model) AndOnSymbol(field, symbol string, value interface{}) Oner {
	m.setJoinCond(field, symbol, "and", value)
	return m
}

// AndOnIn Join 条件on
func (m *Model) AndOnIn(field string, value interface{}) Oner {
	m.setJoinCond(field, "in", "and", value)
	return m
}

// AndOnNotIn Join 条件on
func (m *Model) AndOnNotIn(field string, value interface{}) Oner {
	m.setJoinCond(field, "not in", "and", value)
	return m
}

// Union 联合查询
func (m *Model) Union(model Modeler) Unioner {
	m.unionModel = model
	return m
}

// Where where条件
func (m *Model) Where(field string, value interface{}) Wherer {
	m.setWhereCond(field, "=", "and", value)
	return m
}

// WhereRaw where条件
func (m *Model) WhereRaw(where string) Wherer {
	m.setWhereCond(fmt.Sprintf("(%s)", where), "", "and", nil)
	return m
}

// WhereSymbol where条件
func (m *Model) WhereSymbol(field, symbol string, value interface{}) Wherer {
	m.setWhereCond(field, symbol, "and", value)
	return m
}

// WhereIn where条件
func (m *Model) WhereIn(field string, value interface{}) Wherer {
	m.setWhereCond(field, "in", "and", value)
	return m
}

// WhereNotIn where条件
func (m *Model) WhereNotIn(field string, value interface{}) Wherer {
	m.setWhereCond(field, "not in", "and", value)
	return m
}

// AndWhere where条件
func (m *Model) AndWhere(field string, value interface{}) Wherer {
	m.setWhereCond(field, "=", "and", value)
	return m
}

// AndWhereRaw where条件
func (m *Model) AndWhereRaw(where string) Wherer {
	m.setWhereCond(fmt.Sprintf("(%s)", where), "", "and", nil)
	return m
}

// AndWhereSymbol where条件
func (m *Model) AndWhereSymbol(field, symbol string, value interface{}) Wherer {
	m.setWhereCond(field, symbol, "and", value)
	return m
}

// AndWhereIn where条件
func (m *Model) AndWhereIn(field string, value interface{}) Wherer {
	m.setWhereCond(field, "in", "and", value)
	return m
}

// AndWhereNotIn where条件
func (m *Model) AndWhereNotIn(field string, value interface{}) Wherer {
	m.setWhereCond(field, "not in", "and", value)
	return m
}

// OrWhere where条件
func (m *Model) OrWhere(field string, value interface{}) Wherer {
	m.setWhereCond(field, "=", "or", value)
	return m
}

// OrWhereRaw where条件
func (m *Model) OrWhereRaw(where string) Wherer {
	m.setWhereCond(fmt.Sprintf("(%s)", where), "", "or", nil)
	return m
}

// OrWhereSymbol where条件
func (m *Model) OrWhereSymbol(field, symbol string, value interface{}) Wherer {
	m.setWhereCond(field, symbol, "or", value)
	return m
}

// OrWhereIn where条件
func (m *Model) OrWhereIn(field string, value interface{}) Wherer {
	m.setWhereCond(field, "in", "or", value)
	return m
}

// OrWhereNotIn where条件
func (m *Model) OrWhereNotIn(field string, value interface{}) Wherer {
	m.setWhereCond(field, "not in", "or", value)
	return m
}

// GroupBy 分组
func (m *Model) GroupBy(field ...string) GroupByer {
	var newField []string
	for _, v := range field {
		newField = append(newField, m.getGroupBy(v))
	}
	m.groupBy = newField
	return m
}

func (m *Model) getGroupBy(field string) string {
	farr := strings.Split(field, ".")
	if len(farr) == 1 {
		if !strings.Contains(farr[0], "`") {
			return "`" + farr[0] + "`"
		}
	} else if len(farr) == 2 {
		if !strings.Contains(farr[1], "`") {
			return farr[0] + ".`" + farr[1] + "`"
		}
	}
	return field
}

// Having having条件
func (m *Model) Having(field string, value interface{}) Havinger {
	m.setHavingCond(field, "=", "and", value)
	return m
}

// HavingRaw having条件
func (m *Model) HavingRaw(having string) Havinger {
	m.setHavingCond(fmt.Sprintf("(%s)", having), "", "and", nil)
	return m
}

// HavingSymbol having条件
func (m *Model) HavingSymbol(field, symbol string, value interface{}) Havinger {
	m.setHavingCond(field, symbol, "and", value)
	return m
}

// HavingIn having条件
func (m *Model) HavingIn(field string, value interface{}) Havinger {
	m.setHavingCond(field, "in", "and", value)
	return m
}

// HavingNotIn having条件
func (m *Model) HavingNotIn(field string, value interface{}) Havinger {
	m.setHavingCond(field, "not in", "and", value)
	return m
}

// AndHaving having条件
func (m *Model) AndHaving(field string, value interface{}) Havinger {
	m.setHavingCond(field, "=", "and", value)
	return m
}

// AndHavingRaw having条件
func (m *Model) AndHavingRaw(having string) Havinger {
	m.setHavingCond(fmt.Sprintf("(%s)", having), "", "and", nil)
	return m
}

// AndHavingSymbol having条件
func (m *Model) AndHavingSymbol(field, symbol string, value interface{}) Havinger {
	m.setHavingCond(field, symbol, "and", value)
	return m
}

// AndHavingIn having条件
func (m *Model) AndHavingIn(field string, value interface{}) Havinger {
	m.setHavingCond(field, "in", "and", value)
	return m
}

// AndHavingNotIn having条件
func (m *Model) AndHavingNotIn(field string, value interface{}) Havinger {
	m.setHavingCond(field, "not in", "and", value)
	return m
}

// OrHaving having条件
func (m *Model) OrHaving(field string, value interface{}) Havinger {
	m.setHavingCond(field, "=", "or", value)
	return m
}

// OrHavingRaw having条件
func (m *Model) OrHavingRaw(having string) Havinger {
	m.setHavingCond(fmt.Sprintf("(%s)", having), "", "or", nil)
	return m
}

// OrHavingSymbol having条件
func (m *Model) OrHavingSymbol(field, symbol string, value interface{}) Havinger {
	m.setHavingCond(field, symbol, "or", value)
	return m
}

// OrHavingIn having条件
func (m *Model) OrHavingIn(field string, value interface{}) Havinger {
	m.setHavingCond(field, "in", "or", value)
	return m
}

// OrHavingNotIn having条件
func (m *Model) OrHavingNotIn(field string, value interface{}) Havinger {
	m.setHavingCond(field, "not in", "or", value)
	return m
}

// getTables 表名
func (m *Model) getTables() (string, []interface{}) {
	var table string = m.tableName
	var values []interface{}
	for _, joinTable := range m.otherTableName {
		ons, v := m.getOn(joinTable.on)
		if len(v) > 0 {
			values = append(values, v...)
		}
		//fmt.Println(ons)
		table += fmt.Sprintf(" %s %s", joinTable.tableName, ons)

	}
	return table, values
}

// getOn 获得on条件
func (m *Model) getOn(ons []cond) (string, []interface{}) {
	on := "on"
	var values []interface{}
	for i, o := range ons {
		var q string = ""
		if o.value != nil {
			s := lib.InterfaceToString(o.value)
			regex, _ := regexp.Compile("\\.`[\\s\\S]+?`")
			if regex.MatchString(s) {
				q = s
			} else {
				if o.symbol == "in" || o.symbol == "not in" {
					var mq []string
					switch o.value.(type) {
					case []int:
						vs := o.value.([]int)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []int64:
						vs := o.value.([]int64)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []string:
						vs := o.value.([]string)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []float32:
						vs := o.value.([]float32)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []float64:
						vs := o.value.([]float64)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					}
					q = "(" + strings.Join(mq, ",") + ")"
				} else {
					if ok, _ := m.ValueContaionField(o.value); !ok {
						values = append(values, o.value)
						q = "?"
					}
				}
			}
			//values=append(values, o.value)
			//q="?"
		}
		if i == 0 {
			if ok, str := m.ValueContaionField(o.value); ok {
				on += fmt.Sprintf("%s %s %s %s", o.logic, m.getField(o.field, o.symbol), o.symbol, str)
			} else {
				on += fmt.Sprintf("%s %s %s "+q, o.logic, m.getField(o.field, o.symbol), o.symbol)
			}
		} else {
			if ok, str := m.ValueContaionField(o.value); ok {
				on += fmt.Sprintf(" %s %s %s %s", o.logic, m.getField(o.field, o.symbol), o.symbol, str)
			} else {
				on += fmt.Sprintf(" %s %s %s "+q, o.logic, m.getField(o.field, o.symbol), o.symbol)
			}
		}
	}
	//fmt.Println(on,values)
	return on, values
}

// getDeleteName 获得删除字段名称
func (m *Model) getDeleteName() string {
	if m.DeleteName == "" {
		return "`delete_time`"
	}
	return "`" + m.DeleteName + "`"
}

// getAddDeleteName 获得删除字段名称,在增改的时候使用
func (m *Model) getAddDeleteName() string {
	if m.DeleteName == "" {
		if m.aliasName != "" {
			return m.aliasName + ".`delete_time`"
		}
		return "delete_time"
	}
	if m.aliasName != "" {
		return m.aliasName + ".`" + m.DeleteName + "`"
	}
	return m.DeleteName
}

// 值里面是否含有字段
func (m *Model) ValueContaionField(value interface{}) (bool, string) {
	if str, ok := value.(string); ok {
		regex, _ := regexp.Compile("\\.`[\\s\\S]+?`")
		if regex.MatchString(str) {
			return true, str
		}
	}
	return false, ""
}

// getWhere
func (m *Model) getWhere() (string, []interface{}) {
	where := ""
	var values []interface{}
	for i, w := range m.where {
		var q string = ""
		if w.value != nil {
			s := lib.InterfaceToString(w.value)
			regex, _ := regexp.Compile("`[\\s\\S]+?`")
			arr := regex.FindStringSubmatchIndex(s)
			if len(arr) > 0 && arr[0] == 0 {
				q = s
			} else {
				if w.symbol == "in" || w.symbol == "not in" {
					var mq []string
					switch w.value.(type) {
					case []int:
						vs := w.value.([]int)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []int64:
						vs := w.value.([]int64)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []string:
						vs := w.value.([]string)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []float32:
						vs := w.value.([]float32)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					case []float64:
						vs := w.value.([]float64)
						for _, v := range vs {
							values = append(values, v)
							mq = append(mq, "?")
						}
					}
					q = "(" + strings.Join(mq, ",") + ")"
				} else {
					if ok, _ := m.ValueContaionField(w.value); !ok {
						//值里面不含字段
						values = append(values, w.value)
						q = "?"
					}
				}
			}
		}
		if i == 0 {
			if ok, str := m.ValueContaionField(w.value); ok {
				where += fmt.Sprintf("%s %s %s %s", w.logic, m.getField(w.field, w.symbol), w.symbol, str)
			} else {
				where += fmt.Sprintf("%s %s %s "+q, w.logic, m.getField(w.field, w.symbol), w.symbol)
			}

		} else {
			if ok, str := m.ValueContaionField(w.value); ok {
				where += fmt.Sprintf(" %s %s %s %s", w.logic, m.getField(w.field, w.symbol), w.symbol, str)
			} else {
				where += fmt.Sprintf(" %s %s %s "+q, w.logic, m.getField(w.field, w.symbol), w.symbol)
			}
		}
	}
	if m.SoftDelete {
		logic := "and"
		if len(m.where) == 0 {
			logic = ""
		}
		if m.aliasName != "" {
			where += fmt.Sprintf(" %s %s %s ", logic, m.aliasName+"."+m.getDeleteName(), "=?")
			values = append(values, -1)
		} else {
			where += fmt.Sprintf(" %s %s %s ", logic, m.getDeleteName(), "=?")
			values = append(values, -1)
		}
		//fmt.Println(where,values);
	}
	return where, values
}

// getGroup
func (m *Model) getGroup() (string, []interface{}) {
	if len(m.groupBy) == 0 {
		return "", nil
	}
	group := fmt.Sprintf(" GROUP BY %s ", strings.Join(m.groupBy, ","))
	if len(m.having) > 0 {
		group = group + " HAVING "
	}
	var values []interface{}
	for i, w := range m.having {
		var q string = ""
		if w.value != nil {
			if w.symbol == "in" || w.symbol == "not in" {
				var mq []string
				switch w.value.(type) {
				case []int:
					vs := w.value.([]int)
					for _, v := range vs {
						values = append(values, v)
						mq = append(mq, "?")
					}
				case []int64:
					vs := w.value.([]int64)
					for _, v := range vs {
						values = append(values, v)
						mq = append(mq, "?")
					}
				case []string:
					vs := w.value.([]string)
					for _, v := range vs {
						values = append(values, v)
						mq = append(mq, "?")
					}
				case []float32:
					vs := w.value.([]float32)
					for _, v := range vs {
						values = append(values, v)
						mq = append(mq, "?")
					}
				case []float64:
					vs := w.value.([]float64)
					for _, v := range vs {
						values = append(values, v)
						mq = append(mq, "?")
					}
				}
				q = "(" + strings.Join(mq, ",") + ")"
			} else {
				if ok, _ := m.ValueContaionField(w.value); !ok {
					values = append(values, w.value)
					q = "?"
				}
			}
		}
		if i == 0 {
			if ok, str := m.ValueContaionField(w.value); ok {
				group += fmt.Sprintf("%s %s %s %s", w.logic, m.getField(w.field, w.symbol), w.symbol,str)
			} else {
				group += fmt.Sprintf("%s %s %s "+q, w.logic, m.getField(w.field, w.symbol), w.symbol)
			}

		} else {
			if ok, str := m.ValueContaionField(w.value); ok {
				group += fmt.Sprintf(" %s %s %s %s", w.logic, m.getField(w.field, w.symbol), w.symbol,str)
			} else {
				group += fmt.Sprintf(" %s %s %s "+q, w.logic, m.getField(w.field, w.symbol), w.symbol)
			}
		}
	}
	return group, values
}

// getField
func (m *Model) getField(field, symbol string) string {
	if symbol == "" {
		return field
	}
	farr := strings.Split(field, ".")
	if len(farr) == 1 {
		if !strings.Contains(farr[0], "`") {
			return "`" + farr[0] + "`"
		}
	} else if len(farr) == 2 {
		if !strings.Contains(farr[1], "`") {
			return farr[0] + ".`" + farr[1] + "`"
		}
	}
	return field
}

// getOrderBy 获得排序
func (m *Model) getOrderBy() string {
	if m.orderBy != "" {
		return fmt.Sprintf(" ORDER BY %s", m.getOrderByField(m.orderBy))
	}
	return ""
}

// getOrderByField
func (m *Model) getOrderByField(field string) string {
	field = strings.ReplaceAll(field, "	", " ")
	farr := strings.Split(field, ",")
	for k, v := range farr {
		nfarr := strings.Split(v, ".")
		if len(nfarr) == 1 {
			if !strings.Contains(nfarr[0], "`") {
				nn := strings.Split(nfarr[0], " ")
				if len(nn) >= 2 {
					farr[k] = "`" + nn[0] + "` " + nn[len(nn)-1]
				} else {
					farr[k] = "`" + nfarr[0] + "`"
				}
			}
		} else if len(nfarr) == 2 {
			if !strings.Contains(nfarr[1], "`") {
				nn := strings.Split(nfarr[1], " ")
				//fmt.Println(nn)
				if len(nn) >= 2 {
					farr[k] = nfarr[0] + ".`" + nn[0] + "` " + nn[len(nn)-1]
				} else {
					farr[k] = nfarr[0] + ".`" + nfarr[1] + "`"
				}
			}
		}
	}
	return strings.Join(farr, ",")
}

// getLimit 分页
func (m *Model) getLimit() string {
	if m.limit == 0 && m.offset == 0 {
		return ""
	}
	return fmt.Sprintf(" LIMIT %d,%d", m.offset, m.limit)
}

// getUpdateAndDeleteLimit
func (m *Model) getUpdateAndDeleteLimit() string {
	if m.limit == 0 {
		return ""
	}
	return fmt.Sprintf(" LIMIT %d", m.limit)
}

// getFields 获得查询字段
func (m *Model) getFields() string {
	if len(m.fields) == 0 {
		return "*"
	} else {
		return strings.Join(m.fields, ",")
	}

}

// GetModelSql 模型sql
func (m *Model) GetModelSql() (string, []interface{}) {
	tableName, values := m.getTables()
	where, whereValues := m.getWhere()
	if len(whereValues) > 0 {
		values = append(values, whereValues...)
	}
	group, groupValues := m.getGroup()
	if group != "" {
		where = where + group
		if len(groupValues) > 0 {
			values = append(values, groupValues...)
		}
	}
	fields := m.getFields()
	var sql string
	if where != "" {
		sql = fmt.Sprintf("select %s from %s where %s", fields, tableName, where)
	} else {
		sql = fmt.Sprintf("select %s from %s", fields, tableName)
	}
	return sql, values
}

// Count 统计数量
func (m *Model) Count() int64 {
	defer m.resetSoftDel()
	return m.CountByClear(true)
}

// CountByClear 统计数量
func (m *Model) CountByClear(isClear bool) int64 {
	tableName, values := m.getTables()
	where, whereValues := m.getWhere()
	if len(whereValues) > 0 {
		values = append(values, whereValues...)
	}
	group, groupValues := m.getGroup()
	if group != "" {
		where = where + group
		if len(groupValues) > 0 {
			values = append(values, groupValues...)
		}
	}
	if m.unionModel != nil {
		union, unionValues := m.unionModel.GetModelSql()
		if where != "" {
			where = where + " UNION " + union
		} else {
			where = "1=1 UNION " + union
		}
		if len(unionValues) > 0 {
			values = append(values, unionValues...)
		}
	}
	table := NewTable(m.db, tableName)
	var rows lib.SqlRows
	if group != "" {
		rows = table.Select("count(1) as c from (select 1 ", where+") as mm LIMIT 0,1", values...)
	} else {
		rows = table.Select("count(1) as c", where+" LIMIT 0,1", values...)
	}
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	if isClear {
		m.clear()
	}
	if len(rows) == 1 {
		return rows[0]["c"].Int64()
	}
	return -1
}

// Get 获得列表
func (m *Model) Get() lib.SqlRows {
	defer m.resetSoftDel()
	return m.GetByClear(true)
}

// GetByClear 获得列表
func (m *Model) GetByClear(isClear bool) lib.SqlRows {
	tableName, values := m.getTables()
	where, whereValues := m.getWhere()
	//log.DebugPrint("%v",whereValues)
	hasWhere := false
	if len(whereValues) > 0 {
		values = append(values, whereValues...)
	}
	if where != "" {
		hasWhere = true
	}
	group, groupValues := m.getGroup()
	if group != "" {
		where = where + group
		if len(groupValues) > 0 {
			values = append(values, groupValues...)
		}
	}
	if m.unionModel != nil {
		union, unionValues := m.unionModel.GetModelSql()
		if where != "" {
			where = where + " UNION " + union
		} else {
			where = "1=1 UNION " + union
		}
		if len(unionValues) > 0 {
			values = append(values, unionValues...)
		}
	}
	where = where + m.getOrderBy()
	where = where + m.getLimit()
	table := NewTable(m.db, tableName)
	fields := m.getFields()
	rows := table.SelectByHasWhere(fields, where, hasWhere, values...)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	if isClear {
		m.clear()
	}
	return rows
}

// Find 获得单条数据
func (m *Model) Find() lib.SqlRow {
	defer m.resetSoftDel()
	tableName, values := m.getTables()
	where, whereValues := m.getWhere()
	if len(whereValues) > 0 {
		values = append(values, whereValues...)
	}
	group, groupValues := m.getGroup()
	if group != "" {
		where = where + group
		if len(groupValues) > 0 {
			values = append(values, groupValues...)
		}
	}
	if m.unionModel != nil {
		union, unionValues := m.unionModel.GetModelSql()
		if where != "" {
			where = where + " UNION " + union
		} else {
			where = "1=1 UNION " + union
		}
		if len(unionValues) > 0 {
			values = append(values, unionValues...)
		}
	}
	where = where + m.getOrderBy()
	where = where + " LIMIT 0,1"
	table := NewTable(m.db, tableName)
	fields := m.getFields()
	rows := table.Select(fields, where, values...)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	if len(rows) == 1 {
		return rows[0]
	}
	return nil
}

// clear 清理变量
func (m *Model) clear() {
	m.otherTableName = m.otherTableName[:0]
	m.fields = m.fields[:0]
	m.where = m.where[:0]
	m.groupBy = m.groupBy[:0]
	m.having = m.having[:0]
	m.orderBy = ""
	m.offset = 0
	m.limit = 0
	m.aliasName = ""
	m.tableName = m.oldTableName
	m.preParams = m.preParams[:0]
}

// Pager 分页查询
func (m *Model) Pager(page, pageSize int) (lib.SqlRows, lib.SqlRow) {
	defer m.resetSoftDel()
	offset := pageSize * (page - 1)
	m.offset = offset
	m.limit = pageSize
	var pageInfo lib.SqlRow = lib.SqlRow{}
	list := m.GetByClear(false)
	total := m.CountByClear(false)
	pageInfo["total"] = (&lib.Data{}).Set(total)
	pageInfo["pageCount"] = (&lib.Data{}).Set(int(math.Ceil(float64(total) / float64(pageSize))))
	pageInfo["page"] = (&lib.Data{}).Set(page)
	pageInfo["pageSize"] = (&lib.Data{}).Set(pageSize)
	m.clear()
	return list, pageInfo
}

// OrderBy 排序
func (m *Model) OrderBy(orderBy string) OrderByer {
	m.orderBy = orderBy
	return m
}

// Limit 排序
func (m *Model) Limit(offset ...int) Limiter {
	if len(offset) == 2 {
		m.limit = offset[1]
		m.offset = offset[0]
	} else {
		m.offset = 0
		m.limit = offset[0]
	}
	return m
}

// Update 更新操作
func (m *Model) Update(data lib.SqlIn) int64 {
	defer m.resetSoftDel()
	tableName, values := m.getTables()
	where, whereValues := m.getWhere()
	table := NewTable(m.db, tableName)
	data = m.addDeleteTime(data)
	where = where + m.getUpdateAndDeleteLimit()
	i := table.Update(data, values, where, whereValues...)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return i
}

// Insert 插入操作
func (m *Model) Insert(row lib.SqlIn) int64 {
	defer m.resetSoftDel()
	table := NewTable(m.db, m.tableName)
	row = m.addDeleteTime(row)
	id := table.Insert(row)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return id
}

// Create 插入数据
func (m *Model) Create(row lib.SqlIn) int64 {
	return m.Insert(row)
}

// 判断是否添加软删除
func (m *Model) addDeleteTime(row lib.SqlIn) lib.SqlIn {
	if m.SoftDelete {
		if _, ok := row[m.getAddDeleteName()]; !ok {
			row[m.getAddDeleteName()] = -1
		} else {
			d := lib.Data{Value: row[m.getAddDeleteName()]}
			if d.Int() == 0 {
				row[m.getAddDeleteName()] = -1
			}
		}
	}
	return row
}

// Replace 插入操作
func (m *Model) Replace(row lib.SqlIn) int64 {
	defer m.resetSoftDel()
	table := NewTable(m.db, m.tableName)
	row = m.addDeleteTime(row)
	effects := table.Replace(row)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return effects
}

// InsertOnDuplicate 如果你插入的记录导致一个UNIQUE索引或者primary key(主键)出现重复，那么就会认为该条记录存在，则执行update语句而不是insert语句，反之，则执行insert语句而不是更新语句。
func (m *Model) InsertOnDuplicate(row lib.SqlIn, updateRow lib.SqlIn) int64 {
	defer m.resetSoftDel()
	table := NewTable(m.db, m.tableName)
	row = m.addDeleteTime(row)
	effects := table.InsertOnDuplicate(row, updateRow)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return effects
}

// Drop 删除表
func (m *Model) Drop() int64 {
	table := NewTable(m.db, m.tableName)
	effects := table.Drop()
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return effects
}

// Truncate 清空表
func (m *Model) Truncate() int64 {
	table := NewTable(m.db, m.tableName)
	effects := table.Truncate()
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return effects
}

// Optimize 清空表
func (m *Model) Optimize() int64 {
	table := NewTable(m.db, m.tableName)
	effects := table.Optimize()
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return effects
}

// Delete 删除
func (m *Model) Delete() int64 {
	defer m.resetSoftDel()
	tableName, values := m.getTables()
	where, whereValues := m.getWhere()
	table := NewTable(m.db, tableName)
	where = where + m.getUpdateAndDeleteLimit()
	effects := table.Delete(values, where, whereValues...)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return effects
}

// SoftDel 删除
func (m *Model) SoftDel() int64 {
	tableName, values := m.getTables()
	where, whereValues := m.getWhere()
	table := NewTable(m.db, tableName)
	i := table.Update(lib.SqlIn{
		m.getAddDeleteName(): lib.Time(),
	}, values, where, whereValues...)
	m.lastSql = table.GetLastSql()
	m.preSql, m.preParams = table.GetSqlInfo()
	m.clear()
	return i
}

// init 初始化连接池
func init() {
	db = NewDB()
}

// GetDB 获得数据库对象
func GetDB() *DB {
	return db
}

// NewModel 新建一个模型
func NewModel(dbName, tableName string) Modeler {
	model := Model{}
	return model.Init(dbName, tableName)
}

// NewModelByConnectName 新建一个模型
func NewModelByConnectName(connectName, dbName, tableName string) Modeler {
	model := Model{}
	return model.InitByConnectName(connectName, dbName, tableName)
}

// CreateConnectFunc 创建连接
type CreateConnectFunc func(*DB) []MysqlConnect

// InjectCreatePool 注入连接池的创建过程
// connect 连接mysql配置
// fun 回调函数返回mysql连接配置
func InjectCreatePool(fun CreateConnectFunc) {
	if fun != nil {
		configs := fun(db)
		if len(configs) > 0 {
			db.dbPool.AddConnects(configs)
		}
	}
}
