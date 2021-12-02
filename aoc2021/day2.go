package aoc2021

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func parseLint(str string) (string, int64) {
	fields := strings.Fields(str)
	command := fields[0]
	delta, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	return command, delta
}

func Day2Part1(s *bufio.Scanner) {
	x, y := int64(0), int64(0)
	for s.Scan() {
		command, delta := parseLint(s.Text())
		switch command {
		case "forward":
			x += delta
		case "up":
			y -= delta
		case "down":
			y += delta
		}
	}
	println(x * y)
}

func Day2Part2(s *bufio.Scanner) {
	x, y, aim := int64(0), int64(0), int64(0)
	for s.Scan() {
		command, delta := parseLint(s.Text())
		switch command {
		case "forward":
			x += delta
			y += delta * aim
		case "up":
			aim -= delta
		case "down":
			aim += delta
		}
	}
	println(x * y)
}
