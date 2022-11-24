package mssql


import (
    "fmt"
    "testing"
)
func TestInit(t *testing.T){
	var configs []MsSqlConnect
	master:=DBConfig{
			Name:"wjp",
			DatabaseName:"wjp",
			UserName:"sa",
			Password:"wjp",
			Host:"192.168.31.223",
			Port:"1433",
			Charset:"utf8mb4",
			MaxOpenConns:10,
			MaxIdleConns:5,
	}
	configs=append(configs, MsSqlConnect{
		Master:master,
		Slave:nil,
	})
	//file.PutContent("db.config",fmt.Sprintf("%v",configs))
	db:=MsSqlPool{}
	db.Init(configs)
	fmt.Println("test pool")
}