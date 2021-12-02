package main

import (
	"aoc/aoc2021"
	"bufio"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage:    ./aoc2021 <day> <part> <input>")
	}

	file, err := os.Open(os.Args[3])

	if err != nil {
		log.Fatalf("Reading input file failed\n%v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if os.Args[2] != "1" && os.Args[2] != "2" {
		log.Fatalf("Invalid part: '%v'", os.Args[2])
	}
	switch os.Args[1] {
	case "1":
		switch os.Args[2] {
		case "1":
			aoc2021.Part1(scanner)
		case "2":
			aoc2021.Part2(scanner)
		}
	case "2":
		switch os.Args[2] {
		case "1": aoc2021.Day2Part1(scanner)
		case "2": aoc2021.Day2Part2(scanner)
		}
	default:
		log.Fatal("specify dayN")
	}
}
