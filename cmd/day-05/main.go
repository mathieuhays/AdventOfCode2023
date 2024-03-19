package main

import (
	"AdventOfCode2023/internals/day/day5"
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var start time.Time

type ConversionRange struct {
	start int
	end   int
}

func main() {
	start = time.Now()

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

func generateSeedsAndMaps(scanner *bufio.Scanner) (error, []int, [][]day5.ConversionMap) {
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
		return errors.New("failed to detect seed"), nil, nil
	}

	seeds = extractNumbers(firstLine[7:])

	if len(seeds) < 1 {
		return errors.New("empty seed"), nil, nil
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
				return errors.New("invalid conversion table"), nil, nil
			}

			maps[mapIndex] = append(maps[mapIndex], day5.ConversionMap{
				DestinationStart: numbers[0],
				SourceStart:      numbers[1],
				Size:             numbers[2],
				SourceEnd:        numbers[1] + numbers[2] - 1,
			})
		}
	}

	logWithTime("maps. extracting done")

	sortConversionTables(maps)

	logWithTime("maps. initial sort done")

	fillGaps(maps)

	logWithTime("maps. filling gaps done")

	sortConversionTables(maps)

	logWithTime("maps. final sort done")

	return nil, seeds, maps
}

func logWithTime(message string) {
	fmt.Printf("%s - t:%v\n", message, time.Now().Sub(start))
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

func sortConversionTables(conversionMaps [][]day5.ConversionMap) {
	for _, conversionMapCollection := range conversionMaps {
		sort.Slice(conversionMapCollection, func(i, j int) bool {
			return conversionMapCollection[i].SourceStart < conversionMapCollection[j].SourceStart
		})
	}
}

func fillGaps(conversionMaps [][]day5.ConversionMap) {
	for idx, collection := range conversionMaps {
		fillers := []day5.ConversionMap{}
		length := len(collection)

		if collection[0].SourceStart > 0 {
			fillers = append(fillers, day5.ConversionMap{
				DestinationStart: 0,
				SourceStart:      0,
				Size:             collection[0].SourceStart,
				SourceEnd:        collection[0].SourceStart - 1,
			})
		}

		for i := 1; i < length; i++ {
			diff := collection[i].SourceStart - (collection[i-1].SourceEnd + 1)

			if diff > 0 {
				fillers = append(fillers, day5.ConversionMap{
					DestinationStart: collection[i-1].SourceEnd,
					SourceStart:      collection[i-1].SourceEnd,
					Size:             diff,
					SourceEnd:        collection[i].SourceStart - 1,
				})
			}
		}

		conversionMaps[idx] = append(conversionMaps[idx], fillers...)
	}
}

func processPart1(scanner *bufio.Scanner) {
	err, seeds, maps := generateSeedsAndMaps(scanner)
	if err != nil {
		log.Fatalln(err.Error())
	}

	minLocation := -1

	for _, seed := range seeds {
		number := seed

		for _, conversionMapCollection := range maps {
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

		// number should now represent the location
		if minLocation == -1 {
			minLocation = number
		} else if minLocation > number {
			minLocation = number
		}
	}

	fmt.Println("Minimum Location is: ", minLocation)
}

func walk(conversionRanges []ConversionRange, converters [][]day5.ConversionMap, converterIdx int) []ConversionRange {
	if converterIdx >= len(converters) {
		return conversionRanges
	}

	nextRanges := []ConversionRange{}

	for _, numberRange := range conversionRanges {
		for i := 0; i < len(converters[converterIdx]); i++ {
			converter := converters[converterIdx][i]

			if i == 0 && converter.SourceStart > numberRange.start {
				nextRanges = append(nextRanges, ConversionRange{
					start: numberRange.start,
					end:   converter.SourceStart - 1,
				})
			}

			// check overlap
			if converter.SourceStart <= numberRange.end && converter.SourceEnd >= numberRange.start {
				sourceStart := int(math.Max(float64(converter.SourceStart), float64(numberRange.start)))
				sourceEnd := int(math.Min(float64(converter.SourceEnd), float64(numberRange.end)))
				destinationStart, err := converter.Convert(sourceStart)
				if err != nil {
					log.Fatalln(err.Error())
				}

				destinationEnd, err := converter.Convert(sourceEnd)
				if err != nil {
					log.Fatalln(err.Error())
				}

				nextRanges = append(nextRanges, ConversionRange{
					start: destinationStart,
					end:   destinationEnd,
				})
			}

			if i == len(converters[converterIdx])-1 && numberRange.end > converter.SourceEnd {
				nextRanges = append(nextRanges, ConversionRange{
					start: int(math.Max(float64(converter.SourceEnd+1), float64(numberRange.start))),
					end:   numberRange.end,
				})
			}
		}
	}

	return walk(nextRanges, converters, converterIdx+1)
}

func processPart2(scanner *bufio.Scanner) {
	err, seeds, maps := generateSeedsAndMaps(scanner)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if len(seeds)%2 != 0 {
		log.Fatalln("malformed seed. it must come in pairs")
	}

	minLocation := -1

	// generate seed ranges
	seedRanges := []ConversionRange{}

	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, ConversionRange{
			start: seeds[i],
			end:   seeds[i] + seeds[i+1] - 1,
		})
	}

	locationRanges := walk(seedRanges, maps, 0)

	for _, locationRange := range locationRanges {
		if minLocation == -1 {
			minLocation = locationRange.start
		} else if minLocation > locationRange.start {
			minLocation = locationRange.start
		}
	}

	fmt.Println("Minimum Location is: ", minLocation)

	// correct answer: 31161857
	// brute-force version takes 2m51 to process
	// new version takes 1ms
}
