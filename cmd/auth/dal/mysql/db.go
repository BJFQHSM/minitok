package mysql

import (
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"github.com/go-sql-driver/mysql"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

var DB *gorm.DB
var once sync.Once

func InitDB() {
	once.Do(func() {
		util.LogInfo("MySQL initiation starting...")
		var err error
		DB, err = gorm.Open(mysql2.Open(parseConfigToDSN()), &gorm.Config{})
		if err != nil {
			util.LogFatalf("MySQL initiate error : %+v\n", err)
		}

		util.LogInfo("MySQL initiate success!")
	})
}

func parseConfigToDSN() string {
	var user, password, url, db string

	env := os.Getenv("env")
	if env == "dev" {
		confFile := "config/auth-test.yaml"
		conf := (util.Parse(confFile)["mysql"]).(map[interface{}]interface{})
		user = util.InterfaceToStr(conf["username"])
		password = util.InterfaceToStr(conf["password"])
		url = util.InterfaceToStr(conf["url"])
		db = util.InterfaceToStr(conf["database"])
	} else {
		user = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		url = os.Getenv("MYSQL_ADDR")
		db = os.Getenv("MYSQL_DATABASE")
	}

	mysqlConf := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 url,
		DBName:               db,
		InterpolateParams:    true,
		AllowNativePasswords: true,
	}
	//sqlConf := mysql.Config{
	//	User:              util.InterfaceToStr(conf["username"]),
	//	Passwd:            util.InterfaceToStr(conf["password"]),
	//	Net:               "tcp",
	//	Addr:              util.InterfaceToStr(conf["url"]),
	//	DBName:            util.InterfaceToStr(conf["database"]),
	//	InterpolateParams: true,
	//}
	//dsn := mysqlConf.FormatDSN()
	//log.Printf("Info: init mysql dsn: %s", dsn)
	return mysqlConf.FormatDSN()
}
