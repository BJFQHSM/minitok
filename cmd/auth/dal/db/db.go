package db

import (
	"log"
	"os"
	"sync"

	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"github.com/go-sql-driver/mysql"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB
var once sync.Once

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../../")
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
	// dsn := mysqlConf.FormatDSN()
	// log.Printf("Info: init mysql dsn: %s", dsn)
	return mysqlConf.FormatDSN()
}
