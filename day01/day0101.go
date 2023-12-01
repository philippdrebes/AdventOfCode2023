package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := 0
	re := regexp.MustCompile("[0-9]")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		numbers := re.FindAllString(line, -1)

		val, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if err != nil {
			log.Fatal(err)
		}

		acc += val
	}

	fmt.Printf("%d\n", acc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
