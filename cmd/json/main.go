package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "/var/log/dpkg.log"
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Read file failed")
	}

	var data map[string]interface{}

	json.Unmarshal([]byte(f), &data)

	fmt.Println(f)
	fmt.Println(data)
}
