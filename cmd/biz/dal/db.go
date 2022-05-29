package dal

import (
	"context"
	_ "github.com/bytedance2022/minimal_tiktok/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var MongoCli *mongo.Client
var once sync.Once

func InitMongoDB() {
	once.Do(func() {
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
	})
}

type mongoUriConfig struct {
	protocol string
	user     string
	password string
	url      string // host:port
}

func parseMongoConf() string {
	//conf := util.Parse("config/biz.yaml")["mongodb"].(map[interface{}]interface{})
	//log.Printf("%+v\n", conf)
	//uri := mongoUriConfig{
	//	protocol: "mongodb+srv",
	//	user:     util.InterfaceToStr(conf["auth"]),
	//	password: util.InterfaceToStr(conf["password"]),
	//	url:      util.InterfaceToStr(conf["url"]),
	//}
	//URI := fmt.Sprintf("%s://%s:%s@%s", uri.protocol, uri.user, uri.password, uri.url)
	URI:="mongodb://127.0.0.1:27017"
	log.Printf("%s\n", URI)
	return URI
}