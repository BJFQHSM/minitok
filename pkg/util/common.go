package util

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func Parse(path string) map[string]interface{} {
	dir := "D:\\GO\\GOWORK\\src\\minimal_tiktok\\"
	file, err := ioutil.ReadFile(dir + path)
	if err != nil {
		log.Panicf("Failed to read yaml file.\nerr:%v", err)
		return nil
	}
	fmt.Println(string(file))
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
