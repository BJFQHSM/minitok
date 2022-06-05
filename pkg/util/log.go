package util

import (
	"log"
)

func LogInfo(v ...interface{}) {
	log.Println("[INFO]", v)
}

func LogInfof(format string, v ...interface{}) {
	log.Printf("[INFO]"+format, v...)
}

func LogError(v ...interface{}) {
	log.Println("[ERROR]", v)
}

func LogErrorf(format string, v ...interface{}) {
	log.Printf("[ERROR]"+format, v...)
}

func LogPanic(v ...interface{}) {
	log.Panic("[ERROR]", v)
}

func LogPanicf(format string, v ...interface{}) {
	log.Panicf("[PANIC]"+format, v...)
}
