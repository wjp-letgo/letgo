package mysql

import (
	"fmt"
	"testing"

	"github.com/wjp-letgo/letgo/lib"
)

//TestModel
func TestModel(t *testing.T) {
	model:=NewModel("xingtool_base","sys_user_master")
	model2:=NewModel("xingtool_base","sys_user_master")
	model2.Fields("*").Where("id",1)
	ids:=[]int64{1,2,3,4,5}
	model.Fields("*").
	Alias("m").
	Join("sys_shopee_shop as s").
	On("m.`id`","s.`master_id`").
	OrOn("s.master_id",1).
	AndOnRaw("m.id=1 or m.id=2").
	LeftJoin("sys_lazada_shop as l").
	On("m.id", "l.master_id").
	WhereRaw("m.id=1").
	AndWhere("m.id",2).
	OrWhereIn("m.id",ids).
	GroupBy("m.id").
	Having("m.id",1).
	AndHaving("m.id",1).
	OrderBy("m.id desc").Find()
	fmt.Println(model.GetLastSql())
}

func TestPage(t *testing.T){
	model:=NewModel("xingtool_base","sys_user_master")
	model.Fields("*").Limit(0,10).Get()
	//fmt.Println(model.GetLastSql())
}

func TestUnionModel(t *testing.T) {
	model:=NewModel("xingtool_base","sys_user_master")
	model2:=NewModel("xingtool_base","sys_user_master")
	model2.Fields("*")
	model.Fields("*").Union(model2).Find()
	//fmt.Println(model.GetLastSql())
}

func TestPagerModel(t *testing.T) {
	model:=NewModel("xingtool_base","sys_user_master")
	a:=model.Fields("*").WhereSymbol("id", ">", 0).Find()
	x:=lib.SqlRow{}
	s:=lib.Serialize(a)
	lib.UnSerialize(s,&x)
	//fmt.Println("序列化1",x)
}

func TestUpdateModel(t *testing.T) {
	model:=NewModel("xingtool_base","sys_user_master")
	model.Alias("m").Join("sys_shopee_shop as s").On("m.`id`","s.`master_id`").Where("m.id",2).Update(lib.SqlIn{
		"db_code":"300",
	})

	model.Alias("m").Join("sys_shopee_shop as s").On("m.`id`","s.`master_id`").Where("m.id",2).Update(lib.SqlIn{
		"db_code":"500",
	})
	//fmt.Println("======================xxxxx==============",model.GetLastSql())
}

func TestInsertModel(t *testing.T) {
	model:=NewModel("xingtool_base","sys_user_master")
	model.Replace(lib.SqlIn{
		"db_code":"300",
		"table_code":"300",
		"open_id":"xxxxxxxxxxx",
		"mobile":"15860541821",
	})
}

