package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type schemaNumber struct {
	value int
	posx int
	posy int
	machineCode string
}

func findNumCoords(schematic [][]string, locx int, locy int) (schemaNumber, error) {
	var res schemaNumber
	start := schematic[locx][locy]
	var length int
	if _, err := strconv.Atoi(start); err != nil {
		return res, err
	} else {
		res.posx = locx
		valueString := start
		tmpy := locy - 1
		for {
			if tmpy < 0 {
				break
			} else {
				left, err := strconv.Atoi(schematic[locx][tmpy])
				if err != nil {
					break
				} else {
					valueString = strconv.Itoa(left) + valueString
				}
			}
			tmpy = tmpy - 1
		}
		res.posy = tmpy + 1
		tmpy = locy + 1
		for {
			if tmpy > len(schematic[locx]) - 1 {
				break
			} else {
				right, err := strconv.Atoi(schematic[locx][tmpy])
				if err != nil {
					break
				} else {
					valueString += strconv.Itoa(right)
				}
			}
			tmpy = tmpy + 1
		}
		if final, err := strconv.Atoi(valueString); err == nil {
			res.value = final
		} else {
			return res, err
		}
		length = len(valueString)
	}

	for _, row := range getBlock(schematic, res.posx, res.posy, length) {
		for _, col := range row {
			m := regexp.MustCompile(`[^\d\.]`)
			if m.MatchString(col) {
				res.machineCode += col
			}
		}
	}

	return res, nil
}


func getBlock(schematic [][]string, x int, y int, l int) [][]string {
	var res [][]string

	minx := math.Max(float64(x-1), 0)
	maxx := math.Min(float64(x+2), float64(len(schematic)))
	for _, row := range schematic[int(minx):int(maxx)] {
		miny := math.Max(float64(y-1), 0)
		maxy := math.Min(float64(y+l+1), float64(len(row)))
		//fmt.Println(row)
		res = append(res, row[int(miny):int(maxy)])
	}
	return res
}

func printSchematic(schematic [][]string) {
	for row := range schematic {
		for col := range schematic[row] {
			fmt.Printf(" %v", schematic[row][col])
		}
		fmt.Print("\n")
	}
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

	var schematic [][]string

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		schematic = append(schematic, row)
	}

	schemaSet := make(map[schemaNumber]bool)
	for i, row := range schematic {
		for j, _:= range row {
			schema, err := findNumCoords(schematic, i, j)
			if err == nil {
				if schema.machineCode != "" {
					schemaSet[schema] = true
				}
			}
		}
	}

	var result int
	for k := range schemaSet {
		fmt.Println(k)
		result += k.value
	}
	fmt.Println(result)
}
