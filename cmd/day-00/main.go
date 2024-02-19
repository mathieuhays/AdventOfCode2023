package main

import (
	"AdventOfCode2023/internals/io"
	"bufio"
	"fmt"
	"log"
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

	lineCount, err := io.LineCount(readFile)
	if err != nil {
		log.Fatalf("Error processing line count: %v", lineCount)
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

func processPart1(scanner *bufio.Scanner) {

}

func processPart2(scanner *bufio.Scanner) {

}
