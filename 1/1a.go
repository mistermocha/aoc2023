package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func words2nums(word string) string {
		wtoi := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}

	var keys []string
	for key := range wtoi {
		keys = append(keys, key)
	}

	//fmt.Println(strings.Join(keys, "|"))

	replacer := func(match string) string {
		return wtoi[match]
	}

	m1 := regexp.MustCompile(strings.Join(keys, "|"))
	return m1.ReplaceAllStringFunc(word, replacer)
}


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
		fword := words2nums(word)
		//fmt.Println(word, fword)
		for _, r := range strings.Split(fword, "") {
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
		coord := first * 10 + last
		if coord < 10 {
			log.Printf("%s %s %d", scanner.Text(), word, coord)
		}
		coordSum += coord
		//log.Output(1, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}

	fmt.Println(coordSum)
}
