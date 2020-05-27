package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	maxLength         = 20
	defaultScanRange  = 2
	defaultReadOffset = 2
)

// Name represents data in each line from text
type Name struct {
	fname string
	lname string
}

// InitMe init property for Name
func (n *Name) InitMe(fname, lname string) {
	n.fname = strings.Trim(fname, "\n")
	n.lname = strings.Trim(lname, "\n")
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
	// var fileName string
	// fmt.Println("Enter name of the file:")
	// fmt.Scan(&fileName)

	names := readUsingOs("names.txt")
	fmt.Printf("--------")
	fmt.Printf("names value: %v\n", names)
	// fmt.Printf("names value 12: %v\n", names[12])
	// fmt.Printf("length: %d,capacity: %d \n", len(names), cap(names))
}

func readUsingBuf(fileName string) []Name {
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
				var tempName = &Name{}
				tempName.InitMe(string(text[:pos]), string(text[pos+1:]))
				names = append(names, *tempName)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return names
}

func readUsingOs(fileName string) []Name {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	names := make([]Name, 0, 20)

	var scanRange int = 4
	var readOffset int64 = 0
	var foundNewLine bool = false
	var meetEOF bool = false

	text := make([]byte, scanRange)
	count, err := file.Read(text)
	var cycle int = 0
	for count != 0 {
		cycle++
		// fmt.Printf("~~~~~~~~readOffset: %v~~~~scanRange: %v \n", readOffset, scanRange)
		// fmt.Printf("text before check newline and whitespace:%v-with len:%v\n", string(text), len(text))
		foundNewLine = false
		for pos, char := range text {
			// fmt.Printf("------POS: %v char: %c", pos, char)
			if (char == '\n' || meetEOF) && pos != 0 {
				foundNewLine = true
				scanRange = defaultScanRange
				// fmt.Printf("found '\\n' readOffset %v add pos %v in text have '\\n'\n", readOffset, pos)
				line := text[:pos]
				if meetEOF {
					fmt.Printf("Count %v\n", count)
					line = text[:count]
					readOffset += int64(count)
				} else {
					readOffset += int64(pos)
				}
				fmt.Printf("found '\\n' line: %v and pos: %v\n", string(line), pos)
				for pos, char := range line {
					if char == ' ' {
						tempName := &Name{}
						fmt.Printf("ADDING fname %v \n", string(line[:pos]))
						fmt.Printf("ADDING lname %v \n", string(line[pos+1:]))
						tempName.InitMe(string(line[:pos]), string(line[pos+1:]))
						names = append(names, *tempName)
						// only check first new line
						break
					}
				}
				break
			}
		}
		fmt.Println("ReadAtReadAtReadAtReadAtReadAtReadAt")
		if !foundNewLine {
			scanRange *= 2
			// fmt.Printf("NF scanRange %v count: %v readOffset %v\n", scanRange, count, readOffset)
		}
		text = make([]byte, scanRange)
		count, err = file.ReadAt(text, readOffset)
		// fmt.Printf("Error:%v-Count:%v\n", err, count)
		if err == io.EOF && meetEOF {
			break
		}
		if err == io.EOF {
			meetEOF = true
			text = make([]byte, count)
			count, err = file.ReadAt(text, readOffset)
		}
	}
	return names
}
