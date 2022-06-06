package mongo

import (
	"context"
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
	"time"
)

var Cli *mongo.Client
var once sync.Once

func InitDB() {
	once.Do(func() {
		util.LogInfo("MongoDB initiation starting...")

		serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().
			ApplyURI(parseMongoConf()).
			SetServerAPIOptions(serverAPIOptions)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var err error
		Cli, err = mongo.Connect(ctx, clientOptions)
		if err != nil || Cli == nil {
			util.LogFatalf("MongoDB initiate error : %+v\n", err)
		}

		util.LogInfo("MongoDB initiate success!")
	})
}

type mongoUriConfig struct {
	protocol string
	user     string
	password string
	url      string // host:port
}

func parseMongoConf() string {
	//var confFile string
	var protocol, suffix, user, password, url string
	env := os.Getenv("env")
	if env == "dev" {
		confFile := "config/biz-test.yaml"
		protocol = "mongodb+srv"
		conf := util.Parse(confFile)["mongodb"].(map[interface{}]interface{})
		user = util.InterfaceToStr(conf["user"])
		password = util.InterfaceToStr(conf["password"])
		url = util.InterfaceToStr(conf["url"])
	} else {
		protocol = "mongodb"
		suffix = "/?replicaSet=tiktok&connect=direct"
		user = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
		password = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
		url = os.Getenv("MONGO_ADDR")
	}

	//conf := util.Parse(confFile)["mongodb"].(map[interface{}]interface{})
	//log.Printf("%+v\n", conf)
	//uri := mongoUriConfig{
	//	protocol: protocol,
	//	user:     util.InterfaceToStr(conf["user"]),
	//	password: util.InterfaceToStr(conf["password"]),
	//	url:      util.InterfaceToStr(conf["url"]),
	//}
	uri := mongoUriConfig{
		protocol: protocol,
		user:     user,
		password: password,
		url:      url,
	}

	//URI := fmt.Sprintf("%s://%s:%s@%s/?connect=direct", uri.protocol, uri.user, uri.password, uri.url)
	URI := fmt.Sprintf("%s://%s:%s@%s%s", uri.protocol, uri.user, uri.password, uri.url, suffix)
	// URI := "mongodb://127.0.0.1:27017"
	log.Printf("%s\n", URI)
	return URI
}
