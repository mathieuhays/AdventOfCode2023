package main

import (
	"AdventOfCode2023/internals/algo"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

type Card struct {
	id        int
	winning   []int
	played    []int
	matches   []int
	copyCount int
}

func split(source string, separator string) []string {
	return strings.Split(strings.Trim(source, " "), separator)
}

func stringArrayToInt(parts []string) []int {
	numbers := []int{}

	for _, s := range parts {
		if s == "" {
			continue
		}

		number, _ := strconv.Atoi(strings.Trim(s, " "))
		numbers = append(numbers, number)
	}

	return numbers
}

func getCard(cardString string) Card {
	parts := split(cardString, ":")
	labelParts := split(parts[0], " ")
	id, _ := strconv.Atoi(labelParts[1])
	numberParts := split(parts[1], " | ")

	winningNumbers := stringArrayToInt(split(numberParts[0], " "))
	playedNumbers := stringArrayToInt(split(numberParts[1], " "))

	sort.Ints(winningNumbers)
	sort.Ints(playedNumbers)

	return Card{
		id:        id,
		winning:   winningNumbers,
		played:    playedNumbers,
		matches:   []int{},
		copyCount: 1,
	}
}

func checkMatches(cards []Card) []Card {
	for i, c := range cards {
		lastFoundIndex := 0

		for x, w := range c.winning {
			if algo.Search(c.played, w, lastFoundIndex) > -1 {
				cards[i].matches = append(cards[i].matches, w)
				lastFoundIndex = x
			}
		}
	}

	return cards
}

func processPart1(scanner *bufio.Scanner) {
	fmt.Println("Day 4, Part 1")

	var cards []Card

	for scanner.Scan() {
		cards = append(cards, getCard(scanner.Text()))
	}

	cards = checkMatches(cards)
	total := 0

	for _, c := range cards {
		localTotal := int(math.Pow(2, float64(len(c.matches)-1)))

		fmt.Printf("C#%v; Matches: %v; Points: %v\n", c.id, len(c.matches), localTotal)

		total += localTotal
	}

	fmt.Printf("Total: %v\n", total)
}

func processCardCopies(cards []Card) []Card {
	lenCards := len(cards)
	for i, c := range cards {

		for x := 1; x <= len(c.matches); x++ {
			y := i + x

			if y > lenCards {
				break
			}

			cards[y].copyCount += c.copyCount
		}

	}

	return cards
}

func processPart2(scanner *bufio.Scanner) {
	fmt.Println("Day 4, Part 1")

	var cards []Card

	for scanner.Scan() {
		cards = append(cards, getCard(scanner.Text()))
	}

	cards = checkMatches(cards)
	cards = processCardCopies(cards)
	total := 0

	for _, card := range cards {
		fmt.Printf("#%v, %v tickets\n", card.id, card.copyCount)
		total += card.copyCount
	}

	fmt.Printf("Total: %v\n", total)
}
