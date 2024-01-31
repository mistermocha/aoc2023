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
	cardCounts := make(map[int]int)
	var cards int

	m1 := regexp.MustCompile(`Card(.*): (.*) \| (.*)`)
	for scanner.Scan() {
		s := scanner.Text()
		res := m1.FindStringSubmatch(s)
		card, _ := strconv.Atoi(strings.TrimSpace(res[1]))
		cardCounts[card]++

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
		//fmt.Printf("Scoring %v cards\n", cardCounts[card])
		for i := 1; i <= cardCounts[card]; i++ {
			addToCard := card
			for _, p := range picked {
				for _, w := range winners {
					switch {
					case p < w:
						continue
					case p == w:
						score++
						addToCard++
						//fmt.Printf("From score on %vth card %v, Adding card %v\n", i, card, addToCard)
						cardCounts[addToCard]++
						break
					default:
						break
					}
				}
			}
			cards++
			//fmt.Printf("Added one more of card %v\n", cardCounts[addToCard])
		}
		//fmt.Printf("Card %v: Scored %v cards for %v points each\n", card, cardCounts[card], score)
		// The cardCounts are coming out wrong here. I expect them all to be light by one, because the
		// first-issued card is added in at the top of the loop, but they're way off.
		//fmt.Printf("I have %v of card %v: %v in %v is %v\n", cardCounts[card], card, picked, winners, score)
	}
	fmt.Println(cardCounts)
	fmt.Println(cards)
}
