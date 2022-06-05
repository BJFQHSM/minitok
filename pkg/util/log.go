package util

import "log"

func LogInfo(v ...interface{}) {
	log.Println("[INFO]", v)
}

func LogError(v ...interface{}) {
	log.Println("[ERROR]", v)
}

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
}
