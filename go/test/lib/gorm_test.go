package lib

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestGorm(t *testing.T) {

	db, err := gorm.Open("mysql", "root:justfortest@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	type User struct {
		Id   int64  `gorm:"id"`
		Name string `gorm:"name"`
	}
	var users []User

	db.Table("user").Find(&users, "1=1")

	t.Log(users)
}
