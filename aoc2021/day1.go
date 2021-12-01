package aoc2021

import (
	"bufio"
	"log"
)

func Part1(scanner *bufio.Scanner) {
	counter := 0
	for current, prev := 0, readNum(scanner); prev != -1 && current != -1; prev = current {
		current = readNum(scanner)
		if current > prev {
			counter++
		}
	}

	log.Println(counter)
}

func Part2(s *bufio.Scanner) {
	counter := 0

	for a, b, c, d := readNum(s), readNum(s), readNum(s), readNum(s); d != -1; a, b, c, d = b, c, d, readNum(s) {
		if b+c+d > a+b+c {
			counter++
		}
	}

	log.Println(counter)
}
