package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pattern := regexp.MustCompile("[0-9]+")
	matrix := stringTo2DArray(string(file))

	acc := 0

	for line, row := range matrix {
		possibleMatches := pattern.FindAllStringIndex(string(row), -1)

		if len(possibleMatches) == 0 {
			continue
		}

		for _, element := range possibleMatches {
			leftBound := element[0] - 1
			rightBound := element[1]

			if leftBound < 0 {
				leftBound = 0
			}
			if rightBound >= len(row) {
				rightBound = len(row) - 1
			}

			//fmt.Printf("line: %d, start: %d, end: %d\n", line, element[0], element[1])

			// above
			if line > 0 {
				if checkForSymbol(matrix[line-1], leftBound, rightBound+1) {
					acc += parseNumber(matrix, line, element[0], element[1])
					continue
				}
			}

			// same line
			if checkForSymbol(row, leftBound, leftBound) {
				acc += parseNumber(matrix, line, element[0], element[1])
				continue
			}

			if checkForSymbol(row, rightBound, rightBound) {
				acc += parseNumber(matrix, line, element[0], element[1])
				continue
			}

			// below
			if line < len(matrix)-1 {
				if checkForSymbol(matrix[line+1], leftBound, rightBound+1) {
					acc += parseNumber(matrix, line, element[0], element[1])
					continue
				}
			}

		}

	}

	fmt.Printf("%d\n", acc)
}

func checkForSymbol(row []rune, leftBound int, rightBound int) bool {
	if leftBound == rightBound {
		c := row[leftBound]
		if c != '.' && !unicode.IsNumber(c) {
			return true
		}
		return false
	}

	for _, c := range row[leftBound:rightBound] {
		if c != '.' && !unicode.IsNumber(c) {
			return true
		}
	}
	return false
}

func parseNumber(matrix [][]rune, line int, start int, end int) int {
	number := matrix[line][start:end]
	val, err := strconv.Atoi(string(number))
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("Number: %d\n", val)
	return val
}

func stringTo2DArray(str string) [][]rune {
	rows := strings.Split(strings.TrimSpace(str), "\n")

	arr := make([][]rune, len(rows))
	for i := range arr {
		arr[i] = make([]rune, len(rows[0]))
	}

	for i, row := range rows {
		for j, char := range row {
			arr[i][j] = char
		}
	}

	return arr
}
