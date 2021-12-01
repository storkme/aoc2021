package aoc2021

import (
	"bufio"
	"log"
	"strconv"
)

func readNum(scanner *bufio.Scanner) int {
	if scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			log.Fatalf("Failed to parse input%v", err)
		}

		return int(num)
	}
	return -1
}