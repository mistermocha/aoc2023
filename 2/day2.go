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

type Draw struct {
	red int
	blue int
	green int
}

type Game struct {
	index int
	impossible bool
	draws []Draw
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

	var games []Game
	var result int

	for scanner.Scan() {
	  gameDrawSplit := strings.Split(scanner.Text(), ":")
		mG := regexp.MustCompile(`Game (\d+)`)
		res := mG.FindStringSubmatch(gameDrawSplit[0])
		gameNum, _ := strconv.Atoi(res[1])
		mD := regexp.MustCompile(`(\d+) (blue|red|green)`)
		var draws []Draw
		var impossible bool
		for _, rawDraw := range strings.Split(gameDrawSplit[1], ";") {
			var draw Draw
			for _, match := range mD.FindAllStringSubmatch(rawDraw, -1) {
				count, _ := strconv.Atoi(match[1])
				switch match[2] {
					case "blue":
						draw.blue = count
						impossible = impossible || count > 14
					case "red":
						draw.red = count
						impossible = impossible || count > 12
					case "green":
						draw.green = count
						impossible = impossible || count > 13
					default:
						log.Fatal("Something happened")
				}
			}
			draws = append(draws, draw)
		}

		game := Game{gameNum, impossible, draws}
		if !impossible {
			result += gameNum
		}
		games = append(games, game)
		fmt.Println(game)
	}
	fmt.Println(result)
}
