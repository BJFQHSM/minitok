package mysql

import (
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/go-sql-driver/mysql"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var MysqlDB *gorm.DB
var once sync.Once

//type User struct {
//	ID        uint           `gorm:"primaryKey"`
//	Name	string
//	Age		int
//}
//func (User) TableName() string {
//	return "user_info"
//}

func InitMysql() {
	once.Do(func() {
		var err error
		MysqlDB, err = gorm.Open(mysql2.Open(parseConfigToDSN()), &gorm.Config{})
		if err != nil {
			log.Printf("%s", err)
		}
		//var user User
		//db.First(&user)
		//fmt.Printf("%+v", user)
		//db, err := gorm.Open(mysql.New(mysqlConfig))
		//if err != nil {
		//	log.Fatal(err)
		//}
	})
}

func parseConfigToDSN() string {
	conf := utils.Parse("config/mysql.yaml")
	mysqlConf := mysql.Config{
		User:   utils.InterfaceToStr(conf["username"]),
		Passwd: utils.InterfaceToStr(conf["password"]),
		Net:    "tcp",
		Addr:   utils.InterfaceToStr(conf["url"]),
		DBName: utils.InterfaceToStr(conf["database"]),
	}
	//dsn := mysqlConf.FormatDSN()
	//log.Printf("Info: init mysql dsn: %s", dsn)
	return mysqlConf.FormatDSN()
}
