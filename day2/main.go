package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// readIDs reads newline-separated strings from an io.Reader into
// an array of strings, for which a pointer is then returned.
func readIDs(reader io.Reader) (*[]string, error) {
	scanner := bufio.NewScanner(reader)
	var ids []string

	for scanner.Scan() {
		value := strings.TrimSpace(scanner.Text())

		ids = append(ids, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &ids, nil
}

// parseID iterates through an ID string and returns a boolean
// indicating if the string has two and three occurrences of any
// one character, respectively.
func parseID(id string) (bool, bool) {
	counts := make(map[rune]int)
	var twos, threes bool

	for _, c := range id {
		counts[c]++
	}

	for _, v := range counts {
		if v == 2 {
			twos = true
		} else if v == 3 {
			threes = true
		}
	}

	return twos, threes
}

// countDupes iterates through an array of strings and returns the count of
// elements that have two and three duplicate characters, respectively.
func countDupes(input *[]string) (int, int) {
	var twos, threes int
	for _, s := range *input {
		hasTwos, hasThrees := parseID(s)
		if hasTwos {
			twos++
		}
		if hasThrees {
			threes++
		}
	}
	return twos, threes
}

// findSimilar iterates through a sorted input array and returns the first two
// elements with more than `threshold` similar characters. It's super inefficient
// and should not be used by anyone.
func findSimilar(input *[]string, threshold int) (string, string) {
	ids := *input
	l := len(ids)
	for i, _ := range ids[:l-1] {
		var similarCharacters int
		for _, c1 := range ids[i] {
			for _, c2 := range ids[i+1] {
				if c1 == c2 {
					similarCharacters++
				}
			}
		}
		if similarCharacters >= threshold {
			return ids[i], ids[i+1]
		}
	}
	return "", ""
}

func main() {
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// Read IDs into an array
	inputValues, err := readIDs(inputFile)

	// Sort IDs (for part two)
	sort.Strings(*inputValues)

	// Do part one
	twos, threes := countDupes(inputValues)
	fmt.Println("Part one solution:", twos*threes)

	// Get IDs for part two
	i1, i2 := findSimilar(inputValues, 30)
	fmt.Println("Part two IDs:", i1, i2)
}
