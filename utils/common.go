package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func Parse(path string) map[string]interface{} {
	file, err := ioutil.ReadFile(path)
	pwd, err := os.Getwd()
	fmt.Println(pwd)
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
