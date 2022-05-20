package db

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestQueryUser(t *testing.T) {
	user, err := QueryUserByUID(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", user)
}
func TestQueryFollow(t *testing.T) {
	is_follow, err := QueryFollowUserByUID(context.Background(), 1, 4)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(is_follow)
}
