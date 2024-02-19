package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	firstIntRune = '0'
	lastIntRune  = '9'
	numbers      = [][]rune{
		[]rune("one"),
		[]rune("two"),
		[]rune("three"),
		[]rune("four"),
		[]rune("five"),
		[]rune("six"),
		[]rune("seven"),
		[]rune("eight"),
		[]rune("nine"),
	}
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

func getFirstInt(runes []rune) int {
	for i := 0; i < len(runes); i++ {
		if runes[i] >= firstIntRune && runes[i] <= lastIntRune {
			return int(runes[i] - firstIntRune)
		}
	}

	return 0
}

func getLastInt(runes []rune) int {
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= firstIntRune && runes[i] <= lastIntRune {
			return int(runes[i] - firstIntRune)
		}
	}

	return 0
}

func processPart1(scanner *bufio.Scanner) {
	fmt.Println("Processing Day 1, Part 1")
	total := 0

	for scanner.Scan() {
		parts := []rune(scanner.Text())
		firstDigit := getFirstInt(parts)
		lastDigit := getLastInt(parts)
		total += (firstDigit * 10) + lastDigit
	}

	fmt.Printf("Total: %v\n", total)
}

func getFirstNumber(runes []rune) int {
	lastRuneIndex := len(runes) - 1
	for i := 0; i <= lastRuneIndex; i++ {
		if runes[i] >= firstIntRune && runes[i] <= lastIntRune {
			return int(runes[i] - firstIntRune)
		}

		for x, number := range numbers {
			lastIndex := len(number) - 1

			if (i + lastIndex) > lastRuneIndex {
				continue
			}

			for r := 0; r <= lastIndex; r++ {
				offset := i + r

				if runes[offset] != number[r] {
					break
				}

				if r == lastIndex {
					return x + 1
				}
			}
		}
	}

	return 0
}

func getLastNumber(runes []rune) int {
	lastRuneIndex := len(runes) - 1

	for i := lastRuneIndex; i >= 0; i-- {
		if runes[i] >= firstIntRune && runes[i] <= lastIntRune {
			return int(runes[i] - firstIntRune)
		}

		for x, number := range numbers {
			lastIndex := len(number) - 1

			if (i - lastIndex) < 0 {
				continue
			}

			for r := lastIndex; r >= 0; r-- {
				offset := i - (lastIndex - r)

				if runes[offset] != number[r] {
					break
				}

				if r == 0 {
					return x + 1
				}
			}
		}
	}

	return 0
}

func processPart2(scanner *bufio.Scanner) {
	fmt.Println("Processing Day 1, Part 2")
	total := 0

	for scanner.Scan() {
		parts := []rune(scanner.Text())
		firstDigit := getFirstNumber(parts)
		lastDigit := getLastNumber(parts)
		total += (firstDigit * 10) + lastDigit
	}

	fmt.Printf("Total: %v\n", total)
}
