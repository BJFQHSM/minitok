package util

<<<<<<< HEAD
import (
	"log"
)
=======
import "log"
>>>>>>> cloud_deploy

func LogInfo(v ...interface{}) {
	log.Println("[INFO]", v)
}

<<<<<<< HEAD
func LogInfof(format string, v ...interface{}) {
	log.Printf("[INFO]"+format, v...)
}

=======
>>>>>>> cloud_deploy
func LogError(v ...interface{}) {
	log.Println("[ERROR]", v)
}

<<<<<<< HEAD
func LogErrorf(format string, v ...interface{}) {
	log.Printf("[ERROR]"+format, v...)
}

func LogPanic(v ...interface{}) {
	log.Panic("[ERROR]", v)
}

func LogPanicf(format string, v ...interface{}) {
	log.Panicf("[PANIC]"+format, v...)
=======
func LogFatal(v ...interface{}) {
	log.Fatal("[ERROR]", v)
}

func LogInfof(format string, v ...interface{}) {
	log.Printf("[INFO]"+format, v)
}

func LogErrorf(format string, v ...interface{}) {
	log.Printf("[ERROR]"+format, v)
}

func LogFatalf(format string, v ...interface{}) {
	log.Fatalf("[FATAL]"+format, v)
>>>>>>> cloud_deploy
}
