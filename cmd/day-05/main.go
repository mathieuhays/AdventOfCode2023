package main

import (
	"AdventOfCode2023/internals/day/day5"
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

func extractNumbers(source string) []int {
	parts := strings.Split(strings.Trim(source, " "), " ")
	size := len(parts)
	out := make([]int, size)

	for i := 0; i < size; i++ {
		out[i], _ = strconv.Atoi(strings.Trim(parts[i], " "))
	}

	return out
}

func calculateLocation(seed int, conversionMaps *[][]day5.ConversionMap) int {
	number := seed

	for _, conversionMapCollection := range *conversionMaps {
		for _, conversionMap := range conversionMapCollection {
			if conversionMap.InRange(number) {
				n, err := conversionMap.Convert(number)
				if err != nil {
					log.Fatalln(err.Error())
				}

				number = n
				break
			}
		}
	}

	return number
}

func processPart1(scanner *bufio.Scanner) {
	var seeds []int
	maps := [][]day5.ConversionMap{
		[]day5.ConversionMap{}, // seed to soil
		[]day5.ConversionMap{}, // soil to fertilizer
		[]day5.ConversionMap{}, // fertilizer to water
		[]day5.ConversionMap{}, // water to light
		[]day5.ConversionMap{}, // light to temp
		[]day5.ConversionMap{}, // temp to humidity
		[]day5.ConversionMap{}, // humidity to location
	}
	mapIndex := -1

	scanner.Scan()
	firstLine := scanner.Text()

	if firstLine[:6] != "seeds:" {
		log.Fatalln("Invalid input. Missing seed")
	}

	seeds = extractNumbers(firstLine[7:])

	if len(seeds) < 1 {
		log.Fatalln("Seed detected. No data though")
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)

		if lineLength <= 0 {
			continue
		}

		if line[0] >= 'a' && line[0] <= 'z' {
			// headline row
			mapIndex++
		}

		if line[0] >= '0' && line[0] <= '9' {
			numbers := extractNumbers(line)

			if len(numbers) != 3 {
				log.Fatalln("Invalid map config")
			}

			maps[mapIndex] = append(maps[mapIndex], day5.ConversionMap{
				DestinationStart: numbers[0],
				SourceStart:      numbers[1],
				Size:             numbers[2],
			})
		}
	}

	minLocation := -1

	for _, seed := range seeds {
		number := calculateLocation(seed, &maps)

		// number should now represent the location
		if minLocation == -1 {
			minLocation = number
		} else if minLocation > number {
			minLocation = number
		}
	}

	fmt.Println("Minimum Location is: ", minLocation)
}

func processPart2(scanner *bufio.Scanner) {
	maps := [][]day5.ConversionMap{
		[]day5.ConversionMap{}, // seed to soil
		[]day5.ConversionMap{}, // soil to fertilizer
		[]day5.ConversionMap{}, // fertilizer to water
		[]day5.ConversionMap{}, // water to light
		[]day5.ConversionMap{}, // light to temp
		[]day5.ConversionMap{}, // temp to humidity
		[]day5.ConversionMap{}, // humidity to location
	}
	mapIndex := -1

	scanner.Scan()
	firstLine := scanner.Text()

	if firstLine[:6] != "seeds:" {
		log.Fatalln("Invalid input. Missing seed")
	}

	seeds := extractNumbers(firstLine[7:])

	if len(seeds) == 0 || len(seeds)%2 != 0 {
		log.Fatalln("Seed detected. No data though")
	}

	fmt.Println("Done extracting seeds")

	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)

		if lineLength <= 0 {
			continue
		}

		if line[0] >= 'a' && line[0] <= 'z' {
			// headline row
			mapIndex++
		}

		if line[0] >= '0' && line[0] <= '9' {
			numbers := extractNumbers(line)

			if len(numbers) != 3 {
				log.Fatalln("Invalid map config")
			}

			maps[mapIndex] = append(maps[mapIndex], day5.ConversionMap{
				DestinationStart: numbers[0],
				SourceStart:      numbers[1],
				Size:             numbers[2],
			})
		}
	}

	fmt.Println("Done extracting conversion maps")

	minLocation := -1

	for i := 0; i < len(seeds); i += 2 {
		s := seeds[i]
		l := seeds[i+1]

		for x := s; x < s+l; x++ {
			number := calculateLocation(x, &maps)

			// number should now represent the location
			if minLocation == -1 {
				minLocation = number
			} else if minLocation > number {
				minLocation = number
			}
		}
	}

	fmt.Println("Minimum Location is: ", minLocation)

	// correct answer: 31161857
	// currently takes 2m51 to process

	// @TODO improve efficiency of part 2
	// - Check for better way to process seeds vs conversion maps
	// - Try to order seed ranges and conversion maps to be able to bail early
	// - Try go routines with channels
}
