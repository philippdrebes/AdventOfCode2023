package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

		for index, element := range possibleMatches {
			leftBound := element[0] - 1
			rightBound := element[1]

			if leftBound < 0 {
				leftBound = 0
			}
			if rightBound >= len(row) {
				rightBound = len(row) - 1
			}

			fmt.Printf("Line: %d\n", line)

			// same line
			if checkForSymbol(row, rightBound, rightBound) >= 0 && len(possibleMatches) > index+1 {
				next := possibleMatches[index+1]
				number := parseNumber(matrix, line, element[0], element[1])
				second := parseNumber(matrix, line, next[0], next[1])
				fmt.Printf("number: %d, second: %d\n", number, second)
				acc += number * second
				continue
			}

			// below
			if line < len(matrix)-2 {
				symbolPos := checkForSymbol(matrix[line+1], leftBound, rightBound+1)
				if symbolPos >= 0 {
					symbolPos += leftBound

					val := extractNumber(matrix, possibleMatches[index+1:], line, symbolPos)
					if val == 0 {
						matches1 := pattern.FindAllStringIndex(string(matrix[line+1]), -1)
						val = extractNumber(matrix, matches1, line+2, symbolPos)
					}
					if val == 0 {
						matches2 := pattern.FindAllStringIndex(string(matrix[line+2]), -1)
						val = extractNumber(matrix, matches2, line+2, symbolPos)
					}

					number := parseNumber(matrix, line, element[0], element[1])
					fmt.Printf("number: %d, second: %d\n", number, val)
					acc += number * val
					continue
				}
			}

		}

	}

	fmt.Printf("%d\n", acc)
}

func extractNumber(matrix [][]rune, matches [][]int, line int, symbolPos int) int {
	val := 0
	for _, m := range matches {
		if m[0] <= symbolPos && symbolPos <= m[1] {
			val, _ = strconv.Atoi(string(matrix[line][m[0]:m[1]]))
			break
		}
		if m[0]-1 == symbolPos || symbolPos == m[1] {
			val, _ = strconv.Atoi(string(matrix[line][m[0]:m[1]]))
			break
		}
	}
	return val
}

func checkForSymbol(row []rune, leftBound int, rightBound int) int {
	if leftBound == rightBound {
		c := row[leftBound]
		if c == '*' {
			return leftBound
		}
		return -1
	}

	for i, c := range row[leftBound:rightBound] {
		if c == '*' {
			return i
		}
	}
	return -1
}

func parseNumber(matrix [][]rune, line int, start int, end int) int {
	number := matrix[line][start:end]
	val, err := strconv.Atoi(string(number))
	if err != nil {
		log.Fatal()
	}
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
