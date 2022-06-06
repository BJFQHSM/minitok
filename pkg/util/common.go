package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
	"unsafe"

	"gopkg.in/yaml.v2"
)

func Parse(path string) map[string]interface{} {
	dir := os.Getenv("WORK_DIR")
	file, err := ioutil.ReadFile(dir + path)

	if err != nil {
		log.Panicf("Failed to read yaml file.\nerr:%v", err)
		return nil
	}
	fmt.Println(string(file))
	result := make(map[string]interface{})

	err = yaml.Unmarshal(file, &result)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return result
}

func InterfaceToStr(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func GenerateRandomStr(targetLen int64) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, targetLen)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := targetLen-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func MD5Encrypt(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateRandomInt32() uint32 {
	src := rand.NewSource(time.Now().UnixNano())
	return uint32(src.Int63())
}
