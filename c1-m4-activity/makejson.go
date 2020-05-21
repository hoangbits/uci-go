package main

import (
	"encoding/json"
	"fmt"
)

var name, address string

/**
* Zero value of a map: https://tour.golang.org/moretypes/19
* A nil map has no keys, nor can keys be added.
*  var addressBook map[string]string
**/

// AddressBook alias map type
type AddressBook map[string]string

func main() {
	// The make function returns a map of the given type, initialized and ready for use.
	addressBook := make(AddressBook)
	fmt.Printf("Please enter name: \n")
	fmt.Scan(&name)
	fmt.Printf("Please enter address: \n")
	fmt.Scan(&address)
	addressBook[name] = address
	barr, err := json.Marshal(addressBook)
	if err != nil {
		fmt.Printf("Can't create json" + err.Error())
	}
	fmt.Println("-----")
	fmt.Printf("Original Object: %v\n", addressBook)
	fmt.Printf("Json Object: %v\n", barr)

	// unmarshallAddressBook := make(AddressBook)
	// err = json.Unmarshal(barr, &unmarshallAddressBook)
	// fmt.Printf("Unmarshal Object: %v\n", unmarshallAddressBook)
	fmt.Println("-----")
}
