package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func convertStr(match []string) ([]int, error) {
	if len(match) < 2 {
		return nil, nil
	}

	num1, err := strconv.Atoi(match[1])
	if err != nil {
		return []int{}, err
	}
	num2, err := strconv.Atoi(match[2])
	if err != nil {
		return []int{}, err
	}

	return []int{num1, num2}, nil
}

func parseInput(data string) ([][]int, error) {
	matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(data, -1)

	var results [][]int
	for _, match := range matches {
		numMatch, err := convertStr(match)
		if err != nil {
			return [][]int{}, err
		}
		if numMatch == nil {
			continue
		}

		results = append(results, numMatch)

	}
	return results, nil
}

func parseDont(data string) ([][]int, error) {
	var results [][]int
	var donts []string
	doParts := strings.Split(data, "do()")
	for _, part := range doParts {
		dontPart := strings.Split(part, "don't()")
		donts = append(donts, dontPart[0])
	}

	matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(strings.Join(donts, ""), -1)

	for _, match := range matches {
		numParts, err := convertStr(match)
		if err != nil {
			return [][]int{}, err
		}
		if numParts == nil {
			continue
		}

		results = append(results, numParts)
	}
	return results, nil
}

func sumOfProducts(data [][]int) int {
	sumProduct := 0
	for _, v := range data {
		product := v[0] * v[1]
		sumProduct += product
	}
	return sumProduct
}

func main() {
	data, err := readInput("data.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	matches, err := parseInput(data)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	sum := sumOfProducts(matches)
	fmt.Printf("The sum of all mul(int,int) products is %v\n", sum)

	doMatches, err := parseDont(data)
	if err != nil {
		log.Fatalf("Error parsing input w/ don't: %v", err)
	}

	sumDo := sumOfProducts(doMatches)
	fmt.Printf("The sum of all mul(int,int) following a do() and not a don't() is %v\n", sumDo)
}
