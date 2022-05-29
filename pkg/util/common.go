package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func Parse(path string) map[string]interface{} {
	dir := os.Getenv("WORK_DIR")
	file, err := ioutil.ReadFile(dir + path)

	fmt.Println(string(file))
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
