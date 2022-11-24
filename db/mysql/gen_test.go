package mysql
import(
	"testing"
)

func TestGen(t *testing.T){
	GenModelAndEntityByTableName("hyinx","cd_area",
		SlaveDB{
		Name:"bdsy",
		DatabaseName:"bdsy",
		UserName:"wjp",
		Password:"wjp",
		Host:"127.0.0.1",
		Port:"3306",
		Charset:"utf8mb4",
		Prefix: "cd",
	},false)
	/*
	GenModelAndEntity("hyinx",
		SlaveDB{
		Name:"bdsy",
		DatabaseName:"bdsy",
		UserName:"wjp",
		Password:"wjp",
		Host:"127.0.0.1",
		Port:"3306",
		Charset:"utf8mb4",
		Prefix: "cd",
	},false)
	*/
}