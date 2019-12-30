package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	dbmysql *gorm.DB
	err     error
)

func init() {

	dbmysql, err = gorm.Open("mysql", "root:123456@/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	//defer dbmysql.Close()

}

func DBclose() {
	defer dbmysql.Close()
}
