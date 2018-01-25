package mgoT

import (
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
