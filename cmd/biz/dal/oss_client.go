package dal

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
)

var ossClient *oss.Client

func initOSS() {
	var err error
	endPoint := "oss-cn-hangzhou.aliyuncs.com"
	secretKeyId := "LTAI5t8a5gDiMWiL4dPXzv9N"
	accessKeySecret := "rlmbs6Q2zkYpndrZdaOLO1eSHSUtO1"
	ossClient, err = oss.New(endPoint, secretKeyId, accessKeySecret)
	util.LogInfo("Initiating OSSClient...")
	if err != nil {
		util.LogPanic(err.Error())
	} else if ossClient == nil {
		util.LogPanic("fail to init OSSClient")
	}
	util.LogInfo("Initiating OSSClient success!")
}

func PublishToOss(data []byte) (string, error) {
	if ossClient == nil {
		initOSS()
	}

	bucket, err := ossClient.Bucket("minimal-tiktok")
	if err != nil {
		return "", err
	}

	// 上传文件
	suffix := util.GenerateRandomStr(36) + ".mp4"
	err = bucket.PutObject(suffix, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return "minimal-tiktok.oss-cn-hangzhou.aliyuncs.com/" + suffix, nil
}
