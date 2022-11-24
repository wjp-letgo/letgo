package mssql

import (
	"github.com/wjpxxx/letgo/lib"
	"fmt"
	"testing"
)

func TestInsertModel(t *testing.T) {
	model:=NewModel("wjp","wjp")
	model.Insert(lib.SqlIn{
		"name":"600",
	})
	model.Where("id",1).Update(lib.SqlIn{
		"name":"700",
	})
	d:=model.WhereRaw("1=1").Count()
	fmt.Println(d,model.GetLastSql())
}

