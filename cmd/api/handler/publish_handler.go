package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	// var req biz.PublishActionRequest
	resp := &biz.PublishActionResponse{StatusCode: 1}
	// err := c.ShouldBindQuery(&req)
	file, _, err := c.Request.FormFile("data")
	token := c.Request.FormValue("token")
	title := c.Request.FormValue("title")
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	if err != nil {
		log.Printf("获取文件错误%+v", err)
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Printf("错误%+v", err)
	}
	req := &biz.PublishActionRequest{
		Token: token,
		Title: title,
		Data:  buf.Bytes(),
	}
	// log.Printf("视频发布的参数信息：------%+v", req)
	if err != nil {
		log.Printf("%+v", err)
		msg := "invalid request"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("PublishAction request: %+v\n", &req)
		authResp, err := rpc.AuthClient.Authenticate(c, &auth.AuthenticateRequest{Token: req.Token})
		if err != nil || authResp == nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if !authResp.IsAuthed {
			msg := "token invalid"
			resp.StatusMsg = &msg
		} else {
			req.UserIdFromToken = authResp.UserId
			log.Printf("视频发布的参数信息：------%v", req.UserIdFromToken)
			resp, err = rpc.BizClient.PublishAction(c, req)
			if err != nil || resp == nil {
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
		c.JSON(http.StatusOK, resp)
	}
	util.LogInfof("PublishAction response: %+v\n", resp)
}

func QueryPublishList(c *gin.Context) {
	var req biz.QueryPublishListRequest
	resp := &biz.QueryPublishListResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request"
		resp.StatusMsg = &msg
		c.JSON(http.StatusInternalServerError, resp)
	} else {
		util.LogInfof("QueryPublishList request: %+v\n", &req)
		authResp, err := rpc.AuthClient.Authenticate(c, &auth.AuthenticateRequest{Token: req.Token})
		if err != nil || authResp == nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if !authResp.IsAuthed {
			msg := "token invalid"
			resp.StatusMsg = &msg
		} else {
			req.UserIdFromToken = authResp.UserId
			resp, err = rpc.BizClient.QueryPublishList(c, &req)
			if err != nil || resp == nil {
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
		util.LogInfof("QueryPublishList response: %+v\n", resp)
		c.JSON(http.StatusOK, resp)
	}
}
