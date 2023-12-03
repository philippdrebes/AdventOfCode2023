package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		game := strings.Split(line, ":")
		if err != nil {
			log.Fatal(err)
		}

		draws := strings.Split(game[1], ";")

		acc += calculatePower(draws)
	}

	fmt.Printf("%d\n", acc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculatePower(draws []string) int {

	minCounts := make(map[string]int)
	minCounts["red"] = math.MinInt
	minCounts["green"] = math.MinInt
	minCounts["blue"] = math.MinInt

	for _, draw := range draws {
		for _, cubes := range strings.Split(draw, ",") {
			s := strings.Split(strings.TrimSpace(cubes), " ")

			amount, err := strconv.Atoi(strings.TrimSpace(s[0]))
			if err != nil {
				log.Fatal(err)
			}
			color := strings.TrimSpace(s[1])

			if minCounts[color] < amount {
				minCounts[color] = amount
			}
		}
	}

	return minCounts["red"] * minCounts["green"] * minCounts["blue"]
}
