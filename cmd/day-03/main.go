package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

type Position struct {
	x int
	y int
}

type EnginePart struct {
	number   int
	position Position
	length   int
	isValid  bool
}

type Gear struct {
	position Position
	parts    []EnginePart
}

func isInt(char rune) bool {
	return char >= '0' && char <= '9'
}

func getInt(char rune) int {
	return int(char - '0')
}

func partsToInt(parts []int) int {
	var number int

	for z := len(parts) - 1; z >= 0; z-- {
		number += int(math.Pow10(len(parts)-1-z)) * parts[z]
	}

	return number
}

func extractEngineParts(engineMap [][]rune) []EnginePart {
	var parts []EnginePart

	for y, r := range engineMap {
		var numbersParts [][]int
		var positions []Position
		numberIndex := 0

		for x, c := range r {
			isInteger := isInt(c)
			isScanning := len(numbersParts) > 0 && len(numbersParts)-1 == numberIndex && len(numbersParts[numberIndex]) > 0
			relX := x

			if isInteger {
				if isScanning {
					numbersParts[numberIndex] = append(numbersParts[numberIndex], getInt(c))
				} else {
					numbersParts = append(numbersParts, []int{getInt(c)})
					isScanning = true
				}
			}

			if isScanning {
				if !isInteger || x == len(r)-1 {
					if isInteger {
						relX++
					}

					var position = new(Position)
					position.y = y
					position.x = relX - len(numbersParts[numberIndex])
					positions = append(positions, *position)
					numberIndex++
				}
			}
		}

		if len(numbersParts) > 0 {
			for i, n := range numbersParts {
				parts = append(parts, EnginePart{
					number:   partsToInt(n),
					position: positions[i],
					length:   len(n),
					isValid:  false,
				})
			}
		}
	}

	return parts
}

func inspectEngineParts(engineParts []EnginePart, engineMap [][]rune) []EnginePart {
	lenY := len(engineMap)

	for i, e := range engineParts {

	next_eng:
		for dy := -1; dy <= 1; dy++ {
			y := e.position.y + dy

			if y < 0 || y >= lenY {
				continue
			}

			//fmt.Println("line", y)

			lenX := len(engineMap[y])

			for dx := -1; dx <= e.length; dx++ {
				x := e.position.x + dx

				if x < 0 || x >= lenX {
					continue
				}

				c := engineMap[y][x]
				//fmt.Printf("n(%v) c(%v) i:%v .:%v | ", e.number, string(c), isInt(c), c == '.')

				if !isInt(c) && c != '.' {
					//fmt.Printf("E:%v Symbol found: %v\n", e.number, string(c))
					engineParts[i].isValid = true
					break next_eng
				}
			}
		}
	}

	return engineParts
}

func processPart1(scanner *bufio.Scanner) {
	fmt.Println("Day 3, part 1")

	var engineMap [][]rune

	for scanner.Scan() {
		engineMap = append(engineMap, []rune(scanner.Text()))
	}

	engineParts := extractEngineParts(engineMap)
	engineParts = inspectEngineParts(engineParts, engineMap)

	var total int

	for i, e := range engineParts {
		//fmt.Printf("#%v N:%v P:%vx%v L:%v V:%v\n", i, e.number, e.position.x, e.position.y, e.length, e.isValid)
		fmt.Printf("#%v N:%v V:%v\n", i, e.number, e.isValid)

		if e.isValid {
			total += e.number
		}
	}

	fmt.Printf("Total: %v\n", total)
}

func extractGears(engineMap [][]rune) []Gear {
	var gears []Gear

	for y, row := range engineMap {
		for x, r := range row {
			if r == '*' {
				gear := new(Gear)
				gear.position = Position{
					x: x,
					y: y,
				}

				gears = append(gears, *gear)
			}
		}
	}

	return gears
}

func extractGearParts(gears []Gear, engineMap [][]rune) []Gear {
	lenY := len(engineMap)

	for i, gear := range gears {

		for dy := -1; dy <= 1; dy++ {
			y := gear.position.y + dy

			if y < 0 || y >= lenY {
				continue
			}

			dx := -1
			lenX := len(engineMap[y])

			// retro scan to make sure we capture numbers that start before our scanner area
			for {
				x := gear.position.x + dx

				if x <= 0 || x >= lenX {
					break
				}

				if !isInt(engineMap[y][x]) {
					break
				} else {
					dx--
				}
			}

			var numberParts []int
			dx--

			for {
				dx++
				x := gear.position.x + dx

				if x < 0 || x >= lenX {
					break
				}

				r := engineMap[y][x]
				isInteger := isInt(r)
				isScanning := len(numberParts) > 0
				relX := x

				if isInteger {
					relX++

					if isScanning {
						numberParts = append(numberParts, getInt(r))
					} else {
						numberParts = []int{getInt(r)}
					}
				}

				if len(numberParts) > 0 && (!isInteger || x == lenX-1) {
					gears[i].parts = append(gears[i].parts, EnginePart{
						number: partsToInt(numberParts),
						position: Position{
							x: relX - len(numberParts),
							y: y,
						},
						length:  len(numberParts),
						isValid: true,
					})

					numberParts = []int{}
				}

				if dx >= 1 && !isInteger {
					break
				}
			}
		}
	}

	return gears
}

func processPart2(scanner *bufio.Scanner) {
	fmt.Println("Day 3, part 2")

	var engineMap [][]rune

	for scanner.Scan() {
		engineMap = append(engineMap, []rune(scanner.Text()))
	}

	gears := extractGears(engineMap)
	gears = extractGearParts(gears, engineMap)

	var total int

	for _, gear := range gears {
		fmt.Printf("Gear pL: %v p:%vx%v\n", len(gear.parts), gear.position.x, gear.position.y)

		if len(gear.parts) == 2 {
			total += gear.parts[0].number * gear.parts[1].number
		}
	}

	fmt.Printf("Total: %v\n", total)
}
