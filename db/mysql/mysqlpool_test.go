package mysql

import (
    "fmt"
    "testing"
)
func TestInit(t *testing.T){
	var configs []MysqlConnect
	var slaves []SlaveDB
	master:=SlaveDB{
			Name:"xingtool_base",
			DatabaseName:"xingtool_base",
			UserName:"wjp",
			Password:"wjp",
			Host:"127.0.0.1",
			Port:"3306",
			Charset:"utf8mb4",
			MaxOpenConns:10,
			MaxIdleConns:5,
	}
	slaves=append(slaves,SlaveDB{
			Name:"xingtool_base",
			DatabaseName:"xingtool_base",
			UserName:"wjp",
			Password:"wjp",
			Host:"127.0.0.1",
			Port:"3306",
			Charset:"utf8mb4",
			MaxOpenConns:10,
			MaxIdleConns:5,
	})
	configs=append(configs, MysqlConnect{
		Master:master,
		Slave:slaves,
	})
	//file.PutContent("db.config",fmt.Sprintf("%v",configs))
	db:=MysqlPool{}
	db.Init(configs)
	db.GetIncludeReadDB("xingtool_base")
	fmt.Println("test pool")

}