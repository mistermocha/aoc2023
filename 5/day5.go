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

type MapAssociation struct {
	from string
	to   string
}

type MappingData struct {
	destination int
	source      int
	rangeLen    int
}

var Almanac = make(map[MapAssociation][]MappingData)

func walkMapping(seed int) int {
	res := seed
	category := MapAssociation{"seed", "soil"}
	data := Almanac[category]

outer:
	for {
		fmt.Printf("-- %v-to-%v-map --\n", category.from, category.to)
		for _, d := range data {
			if res < d.source+d.rangeLen && seed >= d.source {
				fmt.Printf("%v %v between %v and %v landing at %v\n",
					category.from,
					res,
					d.source+d.rangeLen,
					d.source,
					d.destination+res-d.source)
				res = d.destination + seed - d.source
			} else {
				fmt.Printf("%v %v remained unchanged\n", category.from, res)
			}
		}
		for key, value := range Almanac {
			if key.to == "location" {
				break outer
			} else {
				if key.from == category.to {
					category = key
					data = value
					//fmt.Printf(" >> New category: %v-to-%v-map --\n", category.from, category.to)
					break
				}
			}
		}
	}
	return res
}

func main() {
	fileName := os.Args[1]

	// Responsibly open/close the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ref MapAssociation
	var seeds []int

	seedsMatcher := regexp.MustCompile(`^seeds: (.*)`)
	mapAssocMatcher := regexp.MustCompile(`(\w+)-to-(\w+) map:`)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case seedsMatcher.MatchString(line):
			match := seedsMatcher.FindStringSubmatch(line)
			for _, seed := range strings.Fields(match[1]) {
				s, _ := strconv.Atoi(seed)
				seeds = append(seeds, s)
			}
			//fmt.Printf("Indexing seeds: %v\n", seeds)
		case mapAssocMatcher.MatchString(line):
			match := mapAssocMatcher.FindStringSubmatch(line)
			ref.from = match[1]
			ref.to = match[2]
			//fmt.Printf("Indexing associations: %v\n", ref)
		default:
			data := strings.Fields(line)
			if len(data) == 0 {
				continue
			} else if len(data) != 3 {
				log.Fatalf("Found %v fields (should be 3) in %v", len(data), line)
			} else {
				if ref == (MapAssociation{}) {
					log.Fatalf("No MapAssociation for %v", line)
				}
				var m MappingData
				m.destination, _ = strconv.Atoi(data[0])
				m.source, _ = strconv.Atoi(data[1])
				m.rangeLen, _ = strconv.Atoi(data[2])
				Almanac[ref] = append(Almanac[ref], m)
				//fmt.Printf("Indexing mappings to %v: %v\n", ref, m)
			}
		}

	}

	fmt.Println(seeds)
	for _, seed := range seeds {
		fmt.Printf("Seed %v lands at %v\n", seed, walkMapping(seed))
	}
	fmt.Println(Almanac)
}
