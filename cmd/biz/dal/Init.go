package dal

import (
	"os"
	"sync"
)

var once sync.Once

func Init() {
	once.Do(func() {
		initMongoDB()
		if os.Getenv("PUBLISH_TYPE") == "OSS" {
			initOSS()
		}
	})
}
