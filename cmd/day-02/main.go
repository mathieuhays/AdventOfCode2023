package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) != 3 {
		log.Fatal("Usage: cmd <part:1,2> <input_file>")
	}

	part, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid part number")
	}

	if part < 1 || part > 2 {
		log.Fatal("Part number out of range")
	}

	inputPath := os.Args[2]
	readFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Input file fatal error: %v", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	if part == 1 {
		processPart1(fileScanner)
	} else if part == 2 {
		processPart2(fileScanner)
	} else {
		log.Fatalln("Unexpected use case")
	}

	fmt.Printf("Executed in %v\n", time.Now().Sub(start))
}

type Set struct {
	red   int
	blue  int
	green int
}

type Game struct {
	id       int
	sets     []Set
	maxRed   int
	maxGreen int
	maxBlue  int
}

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

func parseGameString(source string) (*Game, error) {
	parts := strings.Split(source, ":")
	gameId, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		return nil, err
	}

	game := new(Game)
	game.id = gameId

	var sets []Set

	for _, v := range strings.Split(strings.Trim(parts[1], " "), ";") {
		set := new(Set)

		for _, prop := range strings.Split(strings.Trim(v, " "), ",") {
			propParts := strings.Split(strings.Trim(prop, " "), " ")

			if propParts[1] == "blue" {
				set.blue, err = strconv.Atoi(propParts[0])

				if set.blue > game.maxBlue {
					game.maxBlue = set.blue
				}
			} else if propParts[1] == "red" {
				set.red, err = strconv.Atoi(propParts[0])

				if set.red > game.maxRed {
					game.maxRed = set.red
				}
			} else if propParts[1] == "green" {
				set.green, err = strconv.Atoi(propParts[0])

				if set.green > game.maxGreen {
					game.maxGreen = set.green
				}
			}

			if err != nil {
				return nil, err
			}
		}

		sets = append(sets, *set)
	}

	game.sets = sets

	return game, nil
}

func setIsPossible(set Set) bool {
	return set.red <= MaxRed && set.green <= MaxGreen && set.blue <= MaxBlue
}

func gameIsPossible(game Game) bool {
	for _, s := range game.sets {
		if !setIsPossible(s) {
			return false
		}
	}

	return true
}

func processPart1(scanner *bufio.Scanner) {
	fmt.Println("Day 2, Part 1")

	var total int

	for scanner.Scan() {
		game, err := parseGameString(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}

		if gameIsPossible(*game) {
			fmt.Printf("Possible Game: %v\n", game.id)
			total += game.id
		}
	}

	fmt.Printf("Total: %v\n", total)
}

func processPart2(scanner *bufio.Scanner) {
	fmt.Println("Day 2, Part 2")

	var total int

	for scanner.Scan() {
		game, err := parseGameString(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}

		power := game.maxRed * game.maxGreen * game.maxBlue
		fmt.Printf("G#%v: %v\n", game.id, power)

		total += power
	}

	fmt.Printf("Total: %v\n", total)
}
