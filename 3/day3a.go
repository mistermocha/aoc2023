package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)



func getRatio(schematic [][]string, row int, col int) int {
	block = getBlock(schematic, row, col)
	for i, row := range block {
		for j, col := range row {
			if num, err := strconv.Atoi(col); err == nil {

			}
		}
	}
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

}
