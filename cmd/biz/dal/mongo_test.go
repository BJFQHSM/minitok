package dal

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestQuery(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	initMongoDB()
	//QueryVideoById(context.Background(), 1)
	user, err := QueryUserById(context.Background(), 2335433565)
	if err != nil {
		log.Printf("error to query %v\n", err)
		return
	}
	fmt.Printf("%+v\n", user)
}
