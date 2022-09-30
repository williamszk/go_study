package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "myuser:mypassword/table1?charset=utf8&parseTime=True&loc=Loc")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
