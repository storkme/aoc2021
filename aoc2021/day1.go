package aoc2021

import (
	"bufio"
	"log"
	"strconv"
)

func readNum(scanner *bufio.Scanner) int64 {
	if scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			log.Fatalf("Failed to parse input%v", err)
		}

		return num
	}
	return -1
}

func Part1(scanner *bufio.Scanner) {
	counter := 0
	current := int64(0)

	for prev := readNum(scanner); prev != -1 && current != -1; prev = current {
		current = readNum(scanner)
		if current > prev {
			counter++
		}
	}

	log.Println(counter)
}

func ReadInts(scanner bufio.Scanner) ([]int, error) {
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func Part2(scanner *bufio.Scanner) {
	counter := 0
	stuff, err := ReadInts(*scanner)
	if err != nil {
		log.Fatalf("Failed to read ints from input: %v", err)
	}
	for i := 0; i < len(stuff)-3; i++ {
		a, b := stuff[i]+stuff[i+1]+stuff[i+2], stuff[i+1]+stuff[i+2]+stuff[i+3]
		if b > a {
			counter++
		}
	}
	log.Println(counter)
}
