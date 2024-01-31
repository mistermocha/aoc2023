package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
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

	scanner := bufio.NewScanner(file)
	var points int

	m1 := regexp.MustCompile(`Card(.*): (.*) \| (.*)`)
	for scanner.Scan() {
		s := scanner.Text()
		res := m1.FindStringSubmatch(s)

		hold := make(map[int][]int)

		for _, k := range []int{2, 3} {
			hold[k] = []int{}
			for _, i := range strings.Fields(res[k]) {
				num, _ := strconv.Atoi(i)
				hold[k] = append(hold[k], num)
			}
			slices.Sort(hold[k])
		}
		picked := hold[2]
		winners := hold[3]
		var score int
		for _, p := range picked {
			for _, w := range winners {
				switch {
				case p < w:
					continue
				case p == w:
					if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
					break
				default:
					break
				}
			}
		}
		fmt.Printf("Score for %v in %v is %v\n", picked, winners, score)
		points += score
	}
	fmt.Println(points)
}
