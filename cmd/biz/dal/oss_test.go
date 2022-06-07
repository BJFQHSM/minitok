package dal

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestPublish(t *testing.T) {
	open, err := os.Open("/Users/codingdog/Downloads/douyin.mp4")
	if err != nil {
		return
	}

	all, err := ioutil.ReadAll(open)
	if err != nil {
		return
	}

	str, _ := PublishToOss(all)
	fmt.Println(str)
}
