package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]

	// Responsibly open/close the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Sum of all coordinates
	coordSum := 0

	// Read a file line by line and do a thing
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first := -1
		last := -1
		word := scanner.Text()
		for _, r := range strings.Split(word, "") {
			//fmt.Printf("%T %v\n", r, r)
			if s, err := strconv.Atoi(r); err == nil {
			  //fmt.Printf("%T %v :: %T %v\n", r, r, s, s)
				if first == -1 {
					first = s
					last = s
				} else {
					last = s
				}
			}

		}
		//fmt.Printf("Coordinates are %v\n", first * 10 + last)
		coordSum += first * 10 + last
		//log.Output(1, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}

	fmt.Println(coordSum)
}
