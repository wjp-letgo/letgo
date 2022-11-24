package mssql

import (
    "testing"
    "fmt"
)

func TestDDL(t *testing.T){
	var db DB
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
	db.SetPool(NewPools(configs))
	db.SetDB("wjp","wjp")
	fmt.Println(db.Desc("wjp"))
}