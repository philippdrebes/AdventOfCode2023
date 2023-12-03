package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := 0

	maxCounts := make(map[string]int)
	maxCounts["red"] = 12
	maxCounts["green"] = 13
	maxCounts["blue"] = 14

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		game := strings.Split(line, ":")
		number, err := strconv.Atoi(strings.TrimSpace(game[0][4:]))
		if err != nil {
			log.Fatal(err)
		}

		draws := strings.Split(game[1], ";")

		if isImpossible(maxCounts, draws) {
			acc += number
		}
	}

	fmt.Printf("%d\n", acc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isImpossible(maxCounts map[string]int, draws []string) bool {
	for _, draw := range draws {
		for _, cubes := range strings.Split(draw, ",") {
			s := strings.Split(strings.TrimSpace(cubes), " ")

			amount, err := strconv.Atoi(strings.TrimSpace(s[0]))
			if err != nil {
				log.Fatal(err)
			}
			color := strings.TrimSpace(s[1])

			if maxCounts[color] < amount {
				return false
			}

		}
	}

	return true
}
