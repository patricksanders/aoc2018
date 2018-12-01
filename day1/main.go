package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var frequencyMap map[int]bool

// readChanges reads newline-separated integers from an io.Reader into
// an array of integers, for which a pointer is then returned.
func readChanges(reader io.Reader) (*[]int, error) {
	scanner := bufio.NewScanner(reader)
	var changes []int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		changes = append(changes, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &changes, nil

}

// iterateChanges iterates over a provided integer array and adds them to
// the starting value.
func iterateChanges(input *[]int, start int) (int, bool) {
	result := start

	for _, v := range *input {
		result += v
		if frequencyMap[result] {
			// we've seen this value before, so we're done
			return result, true
		} else {
			frequencyMap[result] = true
		}
	}

	return result, false
}

func init() {
	frequencyMap = make(map[int]bool)
}

func main() {
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// Read "changes" into an array
	inputValues, err := readChanges(inputFile)

	var (
		result int
		done   bool
	)

	// Do the thing
	for !done {
		result, done = iterateChanges(inputValues, result)
	}

	fmt.Println(result)
}
