package mongo

import (
	"context"
	"fmt"
	"github.com/wjpxxx/letgo/log"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//configLock
var configLock sync.Mutex
//DBInfo
type DBInfo struct{
	Client *mongo.Client
	Database *mongo.Database
	Config MongoConnect
}
//Close
func (d *DBInfo)Close(){
	if d.Client!=nil{
		ctx,cancel:=context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		if err:=d.Client.Disconnect(ctx);err!=nil{
			log.PanicPrint("mongodb close %v",err)
		}
	}
}

//MongoPooler
type MongoPooler interface{
	GetDB(connectName string)*DBInfo
	Init(connect []MongoConnect)
	AddConnect(connect MongoConnect)
	AddConnects(connects []MongoConnect)
}

//MongoPool
type MongoPool struct{
	pool map[string]*poolDB
}

//GetDB
func (m *MongoPool)GetDB(connectName string)*DBInfo{
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(m.pool[connectName].config.ConnectTimeout)*time.Second)
	defer cancel()
	client,err:=mongo.Connect(ctx,options.Client().ApplyURI(m.pool[connectName].connect))
	if err !=nil{
		log.PanicPrint("mongodb GetDB error %v", err)
	}
	database:=client.Database(m.pool[connectName].config.Database)
	return &DBInfo{
		Client: client,
		Database: database,
		Config: m.pool[connectName].config,
	}
}


//Init
func (m *MongoPool) Init(connects []MongoConnect){
	m.pool=make(map[string]*poolDB)
	for _,connect:=range connects{
		m.AddConnect(connect)
	}
}
//AddConnect 添加连接
func(m *MongoPool)AddConnect(connect MongoConnect) {
	configLock.Lock()
	defer configLock.Unlock()
	if _,ok:=m.pool[connect.Name];ok{
		log.PanicPrint("Mongo data connection already exists")
	}
	connectStr:=m.open(connect)
	if connectStr!=""{
		m.pool[connect.Name]=&poolDB{
			config: connect,
			connect: connectStr,
		}
	}
}
//AddConnects 添加多个连接
func(m *MongoPool)AddConnects(connects []MongoConnect) {
	for _,c:=range connects{
		m.AddConnect(c)
	}
}
//open
func (m *MongoPool)open(connect MongoConnect)string{
	var userPassword string
	var hosts,options []string
	if connect.UserName!="" && connect.Password!=""{
		userPassword=fmt.Sprintf("%s:%s@",connect.UserName,connect.Password)
	}
	for _,h:=range connect.Hosts{
		if h.Port!=""{
			hosts=append(hosts, fmt.Sprintf("%s:%s",h.Hst,h.Port))
		}else{
			hosts=append(hosts, fmt.Sprintf("%s",h.Hst))
		}
		
	}
	if connect.Option.ConnectTimeoutMS>0{
		options=append(options, fmt.Sprintf("connectTimeoutMS=%d",connect.Option.ConnectTimeoutMS))
	}
	if connect.Option.MaxIdleTimeMS>0{
		options=append(options, fmt.Sprintf("maxIdleTimeMS=%d",connect.Option.MaxIdleTimeMS))
	}
	if connect.Option.MaxPoolSize>0{
		options=append(options, fmt.Sprintf("maxPoolSize=%d",connect.Option.MaxPoolSize))
	}
	if connect.Option.MinPoolSize>0{
		options=append(options, fmt.Sprintf("minPoolSize=%d",connect.Option.MinPoolSize))
	}
	if connect.Option.SocketTimeoutMS>0{
		options=append(options, fmt.Sprintf("socketTimeoutMS=%d",connect.Option.SocketTimeoutMS))
	}
	if connect.Option.WtimeoutMS>0{
		options=append(options, fmt.Sprintf("wtimeoutMS=%d",connect.Option.WtimeoutMS))
	}
	if connect.Option.ReplicaSet!=""{
		options=append(options, fmt.Sprintf("replicaSet=%s",connect.Option.ReplicaSet))
	}
	options=append(options, fmt.Sprintf("safe=%t",connect.Option.Safe))
	options=append(options, fmt.Sprintf("slaveOk=%t",connect.Option.SlaveOk))
	return fmt.Sprintf("mongodb://%s%s/%s?%s",userPassword,strings.Join(hosts,","),connect.Database,strings.Join(options,","))
}

//poolDB
type poolDB struct {
	config MongoConnect
	connect string
}

//全局实现者
var pool MongoPooler
var poolLock sync.Mutex

//NewPools 初始化多数据库连接
func NewPools(configs []MongoConnect) MongoPooler{
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool==nil{
		pool=&MongoPool{}
		pool.Init(configs)
	}
	return pool;
}

//NewPool 初始化数据库连接
func NewPool(config MongoConnect) MongoPooler{
	poolLock.Lock()
	defer poolLock.Unlock()
	if pool==nil{
		var configs []MongoConnect
		configs=append(configs, config)
		pool=&MongoPool{}
		pool.Init(configs)
	}
	return pool;
}

