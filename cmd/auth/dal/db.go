package dal

import (
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
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
	})
}

func parseConfigToDSN() string {
	conf := (util.Parse("config/auth.yaml")["mysql"]).(map[interface{}]interface{})
	mysqlConf := mysql.Config{
		User:              util.InterfaceToStr(conf["username"]),
		Passwd:            util.InterfaceToStr(conf["password"]),
		Net:               "tcp",
		Addr:              util.InterfaceToStr(conf["url"]),
		DBName:            util.InterfaceToStr(conf["database"]),
		InterpolateParams: true,
	}
	//dsn := mysqlConf.FormatDSN()
	//log.Printf("Info: init mysql dsn: %s", dsn)
	return mysqlConf.FormatDSN()
}
