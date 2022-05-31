package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func Parse(path string) map[string]interface{} {
	dir := os.Getenv("WORK_DIR")
	file, err := ioutil.ReadFile(dir + path)
	log.Println(string(file))
	if err != nil {
		return nil
	}
	result := make(map[string]interface{})

	err = yaml.Unmarshal(file, &result)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return result
}

func InterfaceToStr(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
