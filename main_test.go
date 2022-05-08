package main

import (
	"github.com/RaymondCode/simple-demo/db/mysql"
	"testing"
)

func TestInitDB(t *testing.T) {
	mysql.InitConfig()
}