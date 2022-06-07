package dal

import (
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/mongo"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/mysql"
)

func Init() {
	mongo.InitDB()
	mysql.InitDB()
}