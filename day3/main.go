package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var fabric [1000][1000]int

type Claim struct {
	Id      int
	HOffset int
	VOffset int
	Width   int
	Height  int
}

// readInput reads newline-separated strings from an io.Reader into
// an array of strings, for which a pointer is then returned.
func readInput(reader io.Reader) ([]*Claim, error) {
	scanner := bufio.NewScanner(reader)
	var input []*Claim

	for scanner.Scan() {
		value := parseClaim(strings.TrimSpace(scanner.Text()))

		input = append(input, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return input, nil
}

// hamfistInt takes a string that presumably contains an int
// and returns an int. Don't do this.
func hamfistInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("you messed up")
	}
	return i
}

// parseClaim takes a line of input and parses it into a Claim.
func parseClaim(claim string) *Claim {
	parts := strings.Split(claim, " ")
	id := strings.Trim(parts[0], "#")
	offsets := strings.Split(strings.Trim(parts[2], ":"), ",")
	dimensions := strings.Split(parts[3], "x")
	c := Claim{
		Id:      hamfistInt(id),
		HOffset: hamfistInt(offsets[0]),
		VOffset: hamfistInt(offsets[1]),
		Width:   hamfistInt(dimensions[0]),
		Height:  hamfistInt(dimensions[1]),
	}
	return &c
}

// insertClaim increments `fabric` "square inches" for a claim. It returns
// the number of twos that result from insertion. Ugly and silly.
func insertClaim(c *Claim) int {
	row := (*c).VOffset
	column := (*c).HOffset
	height := (*c).Height
	width := (*c).Width
	var twos int
	for i := row; i < row+height; i++ {
		for j := column; j < column+width; j++ {
			fabric[j][i] += 1
			if fabric[j][i] == 2 {
				twos++
			}
		}
	}

	return twos
}

// isClaimClean checks to see if a claim is free from intersecting with
// any other claim. Also ugly and silly
func isClaimClean(c *Claim) bool {
	row := (*c).VOffset
	column := (*c).HOffset
	height := (*c).Height
	width := (*c).Width
	// fmt.Printf("adding %dx%d claim at %d,%d\n", width, height, column, row)
	clean := true
	for i := row; i < row+height; i++ {
		for j := column; j < column+width; j++ {
			if fabric[j][i] > 1 {
				clean = false
			}
		}
	}

	return clean
}

// mapClaims iterates through the claims and calls insertClaim() for each.
func mapClaims(claims []*Claim) int {
	var totalTwos int
	for _, c := range claims {
		totalTwos += insertClaim(c)
	}

	return totalTwos
}

// findCleanClaims iterates through the claims and calls isClaimClean()
// for each.
func findCleanClaims(claims []*Claim) {
	for _, c := range claims {
		if isClaimClean(c) {
			fmt.Printf("claim %d is clean\n", (*c).Id)
		}
	}
}

func main() {
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// Read into an array
	inputValues, err := readInput(inputFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// map the claims
	twos := mapClaims(inputValues)
	fmt.Println(twos)

	// find the clean one(s)
	findCleanClaims(inputValues)
}
