package mongo

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)


func TestDB(t *testing.T){
	//tb:=NewTable(NewDB().SetDB("xingtool_orders_001", "xingtool_orders_001"), "orders_001")
	tb:=NewModel("xingtool_orders_001", "orders_002")
	/*
	for i:=0;i<10000; i++ {
		rs:=tb.InsertOne(bson.D{
			{Key:"name",Value:"pi"},
			{Key:"value",Value:3.14},
		})
		fmt.Println(rs.Hex())
	}
	*/
	var data []interface{}
	data=append(data,bson.D{
		{Key:"name",Value:"pi"},
		{Key:"value",Value:3.77777},
	})
	data=append(data,bson.D{
		{Key:"name",Value:"pi"},
		{Key:"value",Value:3.14},
	})
	rx:=tb.InsertMany(data)
	xx:=tb.UpdateByID(rx[0],bson.M{
		"$set":bson.M{
			"name":"fffffff",
		},
	})
	rs:=tb.UpdateMany(bson.M{
		"name":"pi",
	},bson.M{
		"$set":bson.M{
			"name":"wjpkk",
		},
	})
	var r []bson.D
	//tb.Find(bson.D{{Key:"name",Value:"fffffff"}}, &r)
	//tb.Find(bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}}, &r)
	tb.Pager(bson.D{},&r,1,10)
	fmt.Println(rx,rs,xx,r)
}