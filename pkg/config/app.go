package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func Connect() {
	//dsn := "rutvik:rUTVIK@13@/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("sqlite3", "rutvik.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established")

	db = d
}

func GetDB() *gorm.DB {
	return db
}
