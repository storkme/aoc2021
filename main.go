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

	switch os.Args[1] {
	case "1":
		switch os.Args[2] {
		case "1":
			aoc2021.Part1(bufio.NewScanner(file))
		case "2":
			aoc2021.Part2(bufio.NewScanner(file))
		default:
			log.Fatalln("Invalid part for day1")
		}
	default:
		log.Fatal("specify dayN")
	}
}
