package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")
	name, _ := scanner.ReadString('\n')
	name = strings.TrimSuffix(name, "\r\n")
	fmt.Println("Please enter your address")
	addr, _ := scanner.ReadString('\n')
	addr = strings.TrimSuffix(addr, "\r\n")
	newMap := make(map[string]string)
	newMap["name"] = name
	newMap["address"] = addr
	fmt.Println("The contents of the map are -->", newMap)
	jsonmap, err := json.Marshal(newMap)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("The JSON represention is -->", string(jsonmap))

}
