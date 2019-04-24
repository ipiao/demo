package mgoT

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	Id_         string `bson:"_id"`
	Title       string
	Description string
	By          string
	Url         string
	Tags        []string
	Likes       float64
}

func initMgo() *mgo.Database {
	session, err := mgo.Dial("mongodb://47.244.48.245:27017")
	if err != nil {
		log.Fatal(err)
	}
	return session.DB("test")
}

type Message struct {
	Id      int64 `bson:"_id"`
	Content string
}

func TestIsertMgo(t *testing.T) {
	mgb := initMgo()
	m := &Message{Content: "hello"}
	mgb.C("test").Insert(m)
	t.Log(m)

}

func TestMgo(t *testing.T) {
	session, err := mgo.Dial("")
	assert.Nil(t, err)
	assert.Nil(t, session.Ping())
	db := session.DB("runoob")
	col := db.C("col")
	var res []Course
	err = col.Find(bson.M{}).All(&res)
	assert.Nil(t, err)
	t.Log(res)

	id := bson.ObjectId(res[0].Id_).Hex()
	t.Log(id)
}

type User struct {
	Id   int `json:"id"`
	Name string
}

func TestBson(t *testing.T) {
	u := User{
		Id:   1,
		Name: "aaa",
	}
	b, _ := json.Marshal(&u)
	u2 := &User{}
	err := bson.UnmarshalJSON(b, u2)
	if err != nil {
		t.Fatal(err)
	}

	b2, _ := bson.MarshalJSON(u2)
	t.Log(string(b2))
	t.Log(u2)
}
