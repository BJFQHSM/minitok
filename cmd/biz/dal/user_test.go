package dal

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestChangeFollowRelation(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	InitMongoDB()
	err = ChangeFollowRelation(context.Background(), 1, 4)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}
