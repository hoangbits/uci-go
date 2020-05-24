package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	maxLength = 20
)

// Name represents data in each line from text
type Name struct {
	fname string
	lname string
}

// InitMe init property for Name
func (n *Name) InitMe(fname, lname string) {
	n.fname = fname
	n.lname = lname
	var maxRunes []rune
	if len(fname) > maxLength {
		maxRunes = []rune(fname)
		n.fname = string(maxRunes[:maxLength])
	}
	if len(lname) > maxLength {
		maxRunes = []rune(lname)
		n.lname = string(maxRunes[:maxLength])
	}
}

/*
  successively read each line
  create a struct
  add each struct to slice
  iterate and print slice result
*/

func main() {
	var fileName string
	fmt.Println("Enter name of the file:")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	names := make([]Name, 0, 5)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		for pos, char := range []byte(text) {
			if char == ' ' {
				tempName := Name{string(text[:pos]), string(text[pos+1:])}
				names = append(names, tempName)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("names value: %v\n", names)
	fmt.Printf("length: %d,capacity: %d \n", len(names), cap(names))
}
