package main

import (
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/handler"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/gin-gonic/gin"
)

func main() {

	rpc.Init()
	r := gin.New()

	r.Use(logger())

	douyin := r.Group("/douyin")
	user1 := douyin.Group("/user")
	user1.POST("/login/", handler.Login)
	user1.POST("/register/", handler.Register)
	user1.GET("/", handler.QueryInfo)

	publish1 := douyin.Group("/publish")
	publish1.POST("/action/", handler.PublishAction)
	publish1.GET("/list/", handler.QueryPublishList)

	favorite1 := douyin.Group("/favorite")
	favorite1.POST("/action/", handler.FavoriteAction)
	favorite1.GET("/list/", handler.QueryFavoriteList)

	comment1 := douyin.Group("/comment")
	comment1.POST("/action/", handler.CommentAction)
	comment1.GET("/list/", handler.QueryCommentList)

	feed1 := douyin.Group("/feed")
	feed1.GET("", handler.Feed)
	//feed1.GET("/", handler.Feed)

	relation1 := douyin.Group("/relation")
	relation1.POST("/action/", handler.RelationAction)
	relation1.GET("/follow/list/", handler.QueryFollowList)
	relation1.GET("/follower/list/", handler.QueryFollowerList)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("FATAL: api bind error, err = %+v\n", err)
	}
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("INFO: request url: %s\n", c.Request.RequestURI)
		//请求处理
		c.Next()

	}
}
