package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	root := NewNode("\000")
	return &Trie{RootNode: root}
}

type Node struct {
	Char     string
	Value    int
	Children [26]*Node
}

func NewNode(char string) *Node {
	node := &Node{Char: char, Value: -1}
	for i := 0; i < 26; i++ {
		node.Children[i] = nil
	}
	return node
}

func (t *Trie) Insert(word string, value int) error {
	current := t.RootNode
	strippedWord := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	for i := 0; i < len(strippedWord); i++ {
		index := strippedWord[i] - 'a'
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(strippedWord[i]))
		}
		current = current.Children[index]
	}
	current.Value = value
	return nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	trie := NewTrie()
	trie.Insert("one", 1)
	trie.Insert("two", 2)
	trie.Insert("three", 3)
	trie.Insert("four", 4)
	trie.Insert("five", 5)
	trie.Insert("six", 6)
	trie.Insert("seven", 7)
	trie.Insert("eight", 8)
	trie.Insert("nine", 9)

	acc := 0
	re := regexp.MustCompile("[0-9]")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = replaceWords(line, trie)

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

func replaceWords(input string, trie *Trie) string {
	output := ""
	for i := 0; i < len(input); i++ {
		current := trie.RootNode
		index := input[i] - 'a'
		if index > 25 || current == nil || current.Children[index] == nil {
			output += string(input[i])
			continue
		}

		current = current.Children[index]

		for j := i + 1; j < len(input); j++ {
			index := input[j] - 'a'
			if index > 25 || current == nil || current.Children[index] == nil {
				output += string(input[i])
				break
			}

			current = current.Children[index]
			if current.Value == -1 {
				continue
			}

			output += fmt.Sprint(current.Value)
			i = j - 1
			break
		}
	}

	return output
}
