package mysql

import (
	"github.com/wjpxxx/letgo/lib"
	"database/sql"
	"fmt"
	"time"
	"github.com/wjpxxx/letgo/log"
	"sync"
)
//configLock
var configLock sync.Mutex
//MysqlPooler 全局连接池接口
type MysqlPooler interface {
	GetDB(connectName string) *sql.DB
	GetIncludeReadDB(connectName string) *sql.DB
	SetTx(connectName string, tx *sql.Tx)
	GetTx(connectName string) *sql.Tx
	BeginTx(connectName string)
	EndTx(connectName string)
	IsTransaction(connectName string)bool
	Close()
	Init(connect []MysqlConnect)
	AddConnect(connect MysqlConnect)
	AddConnects(connects []MysqlConnect)
}

type poolDB struct {
	master *sql.DB
	slave []*sql.DB
	masterTx *sql.Tx
	isTransaction bool
}

//MysqlPool 全局连接池
type MysqlPool struct{
	pool map[string]*poolDB
}
//GetDB 取出数据库连接
func(m *MysqlPool) GetDB(connectName string) *sql.DB{
	return m.pool[connectName].master
}
//SetTx 设置事务连接
func(m *MysqlPool) SetTx(connectName string,tx *sql.Tx){
	m.pool[connectName].masterTx=tx
}
//GetTx 获得事务连接
func(m *MysqlPool) GetTx(connectName string) *sql.Tx{
	return m.pool[connectName].masterTx
}
//BeginTx 开始事务
func(m *MysqlPool) BeginTx(connectName string){
	m.pool[connectName].isTransaction=true
}
//EndTx 获得事务连接
func(m *MysqlPool) EndTx(connectName string){
	m.pool[connectName].isTransaction=false
}
//GetIncludeReadDB 取出只读数据库连接
func(m *MysqlPool) GetIncludeReadDB(connectName string) *sql.DB{
	rand:=lib.Rand(0,5,lib.Time())
	if rand==1 {
		return m.pool[connectName].master
	}else{
		slaveCount:=len(m.pool[connectName].slave)
		if slaveCount>0{
			slaveIndex:=lib.Rand(0,slaveCount-1,lib.Time())
			return m.pool[connectName].slave[slaveIndex];
		}else{
			return m.pool[connectName].master
		}
		
	}
}

//Close 关闭连接
func(m *MysqlPool) Close(){
	for _,db:=range m.pool{
		db.master.Close();
		for _,slaveDB:=range db.slave{
			slaveDB.Close();
		}
	}
}

//Init 初始化连接池
func(m *MysqlPool) Init(connects []MysqlConnect){
	//log.DebugPrint("初始化pool")
	if m.pool==nil{
		m.pool=make(map[string]*poolDB)
	}
	for _,connect:=range connects{
		m.AddConnect(connect)
	}
}
//AddConnect 添加连接
func(m *MysqlPool)AddConnect(connect MysqlConnect) {
	configLock.Lock()
	defer configLock.Unlock()
	if _,ok:=m.pool[connect.Master.Name];ok{
		//log.PanicPrint("Mysql data connection already exists")
		return 
	}
	//log.DebugPrint("初始化数据库%s",connect.Master.Name)
	master:=m.open(connect.Master)
	if master!=nil{
		//log.DebugPrint("初始化数据库%s",connect.Master.Name)
		m.pool[connect.Master.Name]=&poolDB{
			master:master,
		}
		if connect.Slave!=nil{
			for _,connectSlave:=range connect.Slave{
				if connectSlave.Host==""{
					continue
				}
				slave:=m.open(connectSlave)
				if slave!=nil{
					m.pool[connect.Master.Name].slave=append(m.pool[connect.Master.Name].slave,slave)
				}
			}
		}
		
	}
	//fmt.Println("=====================pl",m.pool)
}
//AddConnects 添加多个连接
func(m *MysqlPool)AddConnects(connects []MysqlConnect) {
	for _,c:=range connects{
		m.AddConnect(c)
	}
}
//open 打开数据库连接
func(m *MysqlPool)open(connect SlaveDB) *sql.DB{
	connectStr:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",connect.UserName,connect.Password,connect.Host,connect.Port,connect.DatabaseName,connect.Charset)
	db, err:=sql.Open("mysql",connectStr)
	if err!=nil{
		log.PanicPrint("open connect fail %s",err.Error())
	}
	log.DebugPrint("=================加载数据源:%s 成功",connect.DatabaseName)
	if connect.MaxIdleConns>0{
		db.SetMaxIdleConns(connect.MaxIdleConns)
	}
	if connect.MaxOpenConns>0{
		db.SetMaxOpenConns(connect.MaxOpenConns)
	}
	if connect.MaxLifetime>0{
		var timeLife time.Duration =time.Duration(connect.MaxLifetime)*time.Second
		db.SetConnMaxLifetime(timeLife)
	}
	return db
}

//IsTransaction 是否开启事务
func(m *MysqlPool) IsTransaction(connectName string) bool{
	//fmt.Println("=============",m.pool)
	return m.pool[connectName].isTransaction
}