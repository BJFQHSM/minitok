package dal

import (
	"context"
	"log"
	"testing"
)

func TestFavoriteAction(t *testing.T) {
	InitMongoDB()
	err := FavoriteAction(context.TODO(), 1001, 1, 1)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}
