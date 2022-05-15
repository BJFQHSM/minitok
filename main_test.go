package main

import (
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"testing"
)

func TestParse(t *testing.T) {
	conf := (util.Parse("config/user.yaml")["mysql"]).(map[interface{}]interface{})
	fmt.Printf("%s\n", conf["url"])
}