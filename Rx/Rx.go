package main

import (
	"log"
	"fmt"
	"strings"
	// "encoding/json"
	jsonsocket "example.com/gojsonsocket"
	// jsonsocket "github.com/Txiaozhe/go-json-socket"
)

func divide_element(result string) (keyword, target string, key, data []string) {
	splite_string := strings.Split(result, ",")
	for i, value := range splite_string {
		length := len(value)
		if i == 0 {
			keyword = value[11:]
		} else if i == 1 {
			target = value[9:]
		} else if i == 2 {
			key = append(key, value[17:])
		} else if i == (len(splite_string)-1) {
			data =  append(data, value[7:strings.LastIndex(value, "}")-2])
		} else {
			if (i % 2 == 0) {
				key = append(key, value[7:])
			} else {
				data =  append(data, value[7:length-1])
			}
		}
	}
	return
}
func main() {
	// connect to remote service.
	conn, err := jsonsocket.Connect("127.0.0.1:3000")
	if err != nil {
		log.Println("connect error: ", err)
	}
	for {
		ch1, err := jsonsocket.HandleMessage(conn)
		if err != nil {
			log.Println("handle msg error: ", err)
		}
		res := <-ch1
		result := res.Data.(string)

		var keyword, target string
		key := []string{} 
		data := []string{} 

		keyword, target, key, data = divide_element(result)
		fmt.Println("\n", keyword, target, key, data)
	}
}