package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var addr string
	fmt.Println("Enter name: ")
	fmt.Scan(&name)
	fmt.Println("Enter address: ")
	fmt.Scan(&addr)

	// var mp map[string]string
	mp := make(map[string]string)

	mp["Name"] = name
	mp["Address"] = addr

	jsn, _ := json.Marshal(mp)
	fmt.Println("JSON data in byte format:- ", jsn)

	fmt.Println("JSON data in string format:- ", string(jsn))

}
