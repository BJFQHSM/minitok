package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"grpc_test/cmd/api/rpc"
	"grpc_test/grpc_gen"
	"log"
	"net/http"
)

func main() {
	rpc.InitUser()
	r := gin.New()

	// todo routines
	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.GET("/login", login)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func login(c *gin.Context) {
	resp, err := rpc.UserClient.Login(context.Background(), &grpc_gen.LoginRequest{})
	if err != nil {
		fmt.Printf("error %v", err)
	}
	fmt.Printf("%+v\n", resp)
	c.JSON(http.StatusOK, resp)
}
