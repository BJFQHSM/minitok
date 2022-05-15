package main

import (
	"github.com/bytedance2022/minimal_tiktok/cmd/api/handler"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	_ "github.com/bytedance2022/minimal_tiktok/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
)

// @title           Swagger API
// @version         1.0

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	rpc.Init()
	r := gin.New()

	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.POST("/login", handler.Login)
	user1.POST("/register", handler.Register)
	user1.GET("", handler.QueryInfo)

	publish1 := v1.Group("/publish")
	publish1.POST("/action", handler.PublishAction)
	publish1.GET("/list", handler.QueryPublishList)

	favorite1 := v1.Group("/favorite")
	favorite1.POST("/action", handler.FavoriteAction)
	favorite1.GET("/list", handler.QueryFavoriteList)

	comment1 := v1.Group("/comment")
	comment1.POST("/action", handler.CommentAction)
	comment1.GET("/list", handler.QueryCommentList)

	feed1 := v1.Group("/feed")
	feed1.GET("", handler.Feed)

	relation1 := v1.Group("/relation")
	relation1.POST("/action", handler.RelationAction)
	relation1.GET("/follow/list", handler.QueryFollowList)
	relation1.GET("/follower/list", handler.QueryFollowerList)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

//func login(c *gin.Context) {
//	resp, err := rpc.UserClient.Login(context.Background(), &user.LoginRequest{})
//	if err != nil {
//		fmt.Printf("error %v", err)
//	}
//	fmt.Printf("%+v\n", resp)
//	c.JSON(http.StatusOK, resp)
//}
