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
	least Draw
	draws []Draw
}

func least(draws []Draw) Draw {
	theLeast := Draw{0, 0, 0}
	for _, draw := range draws {
		if draw.red > theLeast.red {
			theLeast.red = draw.red
		}
		if draw.blue > theLeast.blue {
			theLeast.blue = draw.blue
		}
		if draw.green > theLeast.green {
			theLeast.green = draw.green
		}
	}
	return theLeast
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
		for _, rawDraw := range strings.Split(gameDrawSplit[1], ";") {
			var draw Draw
			for _, match := range mD.FindAllStringSubmatch(rawDraw, -1) {
				count, _ := strconv.Atoi(match[1])
				switch match[2] {
					case "blue":
						draw.blue = count
					case "red":
						draw.red = count
					case "green":
						draw.green = count
					default:
						log.Fatal("Something happened")
				}
			}
			draws = append(draws, draw)
		}

		l := least(draws)
		result += l.red * l.blue * l.green
		game := Game{gameNum, least(draws), draws}

		games = append(games, game)
		fmt.Println(game)
	}
	fmt.Println(result)
}
