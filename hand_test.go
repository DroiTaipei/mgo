package mgo

import (
	"fmt"
	"testing"

	"github.com/DroiTaipei/mgo/bson"
)

func TestBulkUpdateOver1000(t *testing.T) {
	session, err := Dial("localhost:40001")
	if err != nil {
		t.Fail()
	}
	defer session.Close()

	coll := session.DB("mydb").C("mycoll")

	bulk := coll.Bulk()
	for i := 0; i < 1010; i++ {
		bulk.Insert(bson.M{"n": i})
	}
	_, err = bulk.Run()
	if err != nil {
		t.Fail()
	}
	bulk = coll.Bulk()
	for i := 0; i < 1010; i++ {
		bulk.Update(bson.M{"n": i}, bson.M{"$set": bson.M{"m": i}})
	}
	r, err := bulk.Run()
	if err != nil {
		fmt.Println("err:", err)
		t.Fail()
	}
	fmt.Println("Result:", r)
}
