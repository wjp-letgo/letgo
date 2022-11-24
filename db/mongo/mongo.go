package mongo

import (
	"context"
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/wjpxxx/letgo/log"
)

//MongoDB
type MongoDB struct {
	databaseName string
	connectName string
	dbPool MongoPooler
}

//SetPool 设置连接池
func (db *MongoDB)SetPool(pool MongoPooler)*MongoDB{
	db.dbPool=pool
	return db
}

//SetDB 设置连接名和数据库名称
func (db *MongoDB)SetDB(connectName,databaseName string)*MongoDB{
	db.connectName=connectName
	db.databaseName=databaseName
	return db
}

//CreateConnectFunc 创建连接
type CreateConnectFunc func(*MongoDB)[]MongoConnect

//InjectCreatePool 注入连接池的创建过程
//connect 连接mongo配置 可以为空,如果为空需要自己创建数据连接查询配置后返回
//fun 回调函数返回mongo连接配置
func InjectCreatePool(fun CreateConnectFunc){
	if fun!=nil {
		configs:= fun(db)
		if len(configs)>0 {
			db.dbPool.AddConnects(configs)
		}
	}
}

//NewDB 
func NewDB()*MongoDB{
	dbFile:="config/mongo_db.config"
	cfgFile:=file.GetContent(dbFile)
	var configs []MongoConnect
	if cfgFile==""{
		db:=MongoConnect{
			Name:"connectName",
			Database:"databaseName",
			UserName:"userName",
			Password:"password",
			Hosts:[]Host{
				Host{
					Hst:"127.0.0.1",
					Port:"27017",
				},
			},
			ConnectTimeout: 10,
			ExecuteTimeout: 10,
		}
		configs=append(configs,db)
		file.PutContent(dbFile,fmt.Sprintf("%v",configs))
		log.PanicPrint("please setting mongo database config in config/mongo_db.config file")
	}
	lib.StringToObject(cfgFile, &configs)
	var db MongoDB
	return db.SetPool(NewPools(configs))
}

//Tabler
type Tabler interface{
	SetDB(db *MongoDB)
	InsertOne(document interface{})primitive.ObjectID
	InsertMany(documents []interface{})[]primitive.ObjectID
	UpdateOne(filter interface{}, update interface{}) *mongo.UpdateResult
	UpdateMany(filter interface{}, update interface{}) *mongo.UpdateResult
	UpdateByID(id interface{}, update interface{}) *mongo.UpdateResult
	ReplaceOne(filter interface{}, update interface{}) *mongo.UpdateResult
	DeleteOne(filter interface{}) *mongo.DeleteResult
	DeleteMany(filter interface{}) *mongo.DeleteResult
	FindOne(filter interface{},result interface{})
	Find(filter interface{},result interface{})
	Pager(filter interface{},result interface{},page,pageSize int64)
	Aggregate(pipeline interface{},result interface{})
}

//Table 操作表
type Table struct{
	tableName string
	db *MongoDB
}

//NewTable 初始化表
func NewTable(db *MongoDB,tableName string) Tabler{
	var table *Table=&Table{}
	table.tableName=tableName
	table.SetDB(db)
	table.CreateCollection()
	return table
}
//CreateCollection 创建集合
func (t *Table)CreateCollection(){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
	}()
	db.Database.CreateCollection(ctx, t.tableName)
}

//SetDB 设置数据库
func (t *Table) SetDB(db *MongoDB){
	t.db=db
}

//getDB
func (t *Table)getDB() *DBInfo{
	return t.db.dbPool.GetDB(t.db.connectName)
}

//InsertOne
func (t *Table)InsertOne(document interface{})primitive.ObjectID{
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).InsertOne(ctx,document)
	if err!=nil{
		log.PanicPrint("mongo InsertOne error %v", err)
	}
	return res.InsertedID.(primitive.ObjectID)
}

//InsertMany
func (t *Table)InsertMany(documents []interface{})[]primitive.ObjectID{
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).InsertMany(ctx,documents)
	if err!=nil{
		log.PanicPrint("mongo InsertMany error %v", err)
	}
	var result []primitive.ObjectID
	for _,v:=range res.InsertedIDs{
		result=append(result, v.(primitive.ObjectID))
	}
	return result
}

//UpdateOne
func (t *Table)UpdateOne(filter interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).UpdateOne(ctx,filter,update)
	if err!=nil{
		log.PanicPrint("mongo UpdateOne error %v", err)
	}
	return res
}

//UpdateMany
func (t *Table)UpdateMany(filter interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).UpdateMany(ctx,filter,update)
	if err!=nil{
		log.PanicPrint("mongo UpdateMany error %v", err)
	}
	return res
}

//UpdateByID
func (t *Table)UpdateByID(id interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).UpdateByID(ctx,id,update)
	if err!=nil{
		log.PanicPrint("mongo UpdateByID error %v", err)
	}
	return res
}

//ReplaceOne
func (t *Table)ReplaceOne(filter interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).ReplaceOne(ctx,filter,update)
	if err!=nil{
		log.PanicPrint("mongo ReplaceOne error %v", err)
	}
	return res
}

//DeleteOne
func (t *Table)DeleteOne(filter interface{}) *mongo.DeleteResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).DeleteOne(ctx,filter)
	if err!=nil{
		log.PanicPrint("mongo DeleteOne error %v", err)
	}
	return res
}

//DeleteMany
func (t *Table)DeleteMany(filter interface{}) *mongo.DeleteResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).DeleteMany(ctx,filter)
	if err!=nil{
		log.PanicPrint("mongo DeleteMany error %v", err)
	}
	return res
}

//FindOne
func (t *Table)FindOne(filter interface{},result interface{}){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	err:=db.Database.Collection(t.tableName).FindOne(ctx,filter).Decode(result)
	if err!=nil{
		log.PanicPrint("mongo FindOne error %v", err)
	}
}

//Find
func (t *Table)Find(filter interface{},result interface{}){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	cur,err:=db.Database.Collection(t.tableName).Find(ctx,filter)
	if err!=nil{
		log.PanicPrint("mongo Find error %v", err)
	}
	defer cur.Close(context.Background())
	err=cur.All(context.Background(), result)
	if err!=nil{
		log.PanicPrint("mongo Find cur error %v", err)
	}
}

//Pager
func (t *Table)Pager(filter interface{},result interface{},page,pageSize int64){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	option:=options.Find()
	option.SetLimit(pageSize)
	option.SetSkip(pageSize * (page-1))
	cur,err:=db.Database.Collection(t.tableName).Find(ctx,filter,option)
	if err!=nil{
		log.PanicPrint("mongo Pager error %v", err)
	}
	defer cur.Close(context.Background())
	err=cur.All(context.Background(), result)
	if err!=nil{
		log.PanicPrint("mongo Pager cur error %v", err)
	}
}

//Aggregate
func (t *Table)Aggregate(pipeline interface{},result interface{}){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	cur,err:=db.Database.Collection(t.tableName).Aggregate(ctx,pipeline)
	if err!=nil{
		log.PanicPrint("mongo Aggregate error %v", err)
	}
	defer cur.Close(context.Background())
	err=cur.All(context.Background(), result)
	if err!=nil{
		log.PanicPrint("mongo Aggregate cur error %v", err)
	}
}

//NewModel
func NewModel(dbName,tableName string)Tabler{
	return NewTable(db.SetDB(dbName, dbName), tableName)
}

//NewModelByConnectName 新建一个模型
func NewModelByConnectName(connectName,dbName,tableName string) Tabler{
	return NewTable(db.SetDB(connectName, dbName), tableName)
}
//db 全局变量
var db *MongoDB
//init 初始化连接池
func init(){
	db=NewDB()
}