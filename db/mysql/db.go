package mysql

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type database interface {
	get() int
}

var once sync.Once
var mysqlConfig mysql.Config

func InitConfig() {
	conf := config.Parse("config/mysql.yaml")
	fmt.Printf("%+v", conf)
}

func InitMySQL() *gorm.DB {
	var db *gorm.DB
	once.Do(func() {
		InitConfig()
		//db, err := gorm.Open(mysql.New(mysqlConfig))
		//if err != nil {
		//	log.Fatal(err)
		//}

	})
	return db
}