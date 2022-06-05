package dal

import (
	"context"
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var MongoCli *mongo.Client

func initMongoDB() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(parseMongoConf()).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoCli, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

type mongoUriConfig struct {
	protocol string
	user     string
	password string
	url      string // host:port
}

func parseMongoConf() string {
	var confFile string
	var protocol string
	var suffix string
	env := os.Getenv("env")
	if env == "dev" {
		confFile = "config/biz-test.yaml"
		protocol = "mongodb+srv"
		suffix = "/"
	} else {
		confFile = "config/biz.yaml"
		protocol = "mongodb"
		suffix = "/?replicaSet=tiktok&connect=direct"
	}
	conf := util.Parse(confFile)["mongodb"].(map[interface{}]interface{})
	log.Printf("%+v\n", conf)
	uri := mongoUriConfig{
		protocol: protocol,
		user:     util.InterfaceToStr(conf["user"]),
		password: util.InterfaceToStr(conf["password"]),
		url:      util.InterfaceToStr(conf["url"]),
	}

	//URI := fmt.Sprintf("%s://%s:%s@%s/?connect=direct", uri.protocol, uri.user, uri.password, uri.url)
	URI := fmt.Sprintf("%s://%s:%s@%s%s", uri.protocol, uri.user, uri.password, uri.url, suffix)
	// URI := "mongodb://127.0.0.1:27017"
	log.Printf("%s\n", URI)
	return URI
}
