package dal

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"log"
)

var ossClient *oss.Client

func InitOSS() {
	var err error
	once.Do(func() {
		endPoint := "oss-cn-hangzhou.aliyuncs.com"
		secretKeyId := "LTAI5t8a5gDiMWiL4dPXzv9N"
		accessKeySecret := "rlmbs6Q2zkYpndrZdaOLO1eSHSUtO1"
		ossClient, err = oss.New(endPoint, secretKeyId, accessKeySecret)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func PublishToOss(data []byte) (string, error) {
	if ossClient == nil {
		InitOSS()
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
