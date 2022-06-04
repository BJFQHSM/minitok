package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	var req biz.FavoriteActionRequest
	err := c.ShouldBindQuery(&req)

	if err != nil {
		log.Printf("ERROR: 请求参数匹配错误 err = %v:", err)
		var resp *biz.FavoriteActionResponse
		resp.StatusCode = 1
		msg := "Params have errors!"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	}
	resp, err := rpc.BizClient.FavoriteAction(context.Background(), &req)
	if err != nil {
		// todo
		log.Printf("ERROR: 点赞操作报错 err = %v:", err)
	}
	log.Printf("Resp: %+v\n", resp)
	c.JSON(http.StatusOK, resp)
}

//
//type ListResponse struct {
//	VideoList  []*biz.Video `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
//	StatusCode int32        `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
//	StatusMsg  string
//}

func QueryFavoriteList(c *gin.Context) {
	var req biz.QueryFavoriteListRequest
	err := c.ShouldBindQuery(&req)

	if err != nil {
		// todo
		log.Printf("ERROR: 参数列表错误 err = %+v:", err)
	}
	resp, err := rpc.BizClient.QueryFavoriteList(context.Background(), &req)
	if err != nil {
		// todo
		log.Printf("ERROR: 点赞列表报错 err = %v:", err)
	}

	log.Printf("INFO: Resp: %+v\n", resp)
	c.JSON(http.StatusOK, resp)
}
