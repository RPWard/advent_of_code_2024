package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readRules(filename string) map[int][]int {
	data, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error reading rules: %v", err)
	}

	rulesMap := make(map[int][]int)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		key, err := strconv.Atoi(scanner.Text()[:2])
		if err != nil {
			log.Fatalf("could not convert key: %v to int: %v", key, err)
		}

		value, err := strconv.Atoi(scanner.Text()[len(scanner.Text())-2:])
		if err != nil {
			log.Fatalf("could not convert value: %v to int: %v", value, err)
		}

		if _, ok := rulesMap[key]; !ok {
			rulesMap[key] = []int{value}
		} else {
			rulesMap[key] = append(rulesMap[key], value)
		}
	}
	return rulesMap
}

func readPrinting(filename string) [][]int {
	data, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error reading printing: %v", err)
	}

	var printings [][]int
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		strSlice := strings.Split(scanner.Text(), ",")
		var intSlice []int
		for _, val := range strSlice {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalf("error converting value: %v to int: %v", val, err)
			}
			intSlice = append(intSlice, intVal)
		}
		printings = append(printings, intSlice)
	}
	return printings
}

func checkContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findValidPrints(rules map[int][]int, printings [][]int) ([][]int, [][]int) {
	var validPrints [][]int
	var invalidPrints [][]int

	for _, update := range printings {
		// looping through each value in the line
		for i, value := range update {
			if i == 0 {
				continue
			}
			if rule, ok := rules[value]; ok {
				// looping through each value before idx i
				for j := 0; j < i; j++ {
					if checkContains(rule, update[j]) {
						invalidPrints = append(invalidPrints, update)
						goto nextUpdate
					}
				}
			}
		}
		validPrints = append(validPrints, update)
	nextUpdate:
	}
	return validPrints, invalidPrints
}

func fixInvalid(rules map[int][]int, data [][]int) [][]int {
	var fixedData [][]int
	for _, slice := range data {
		sort.Slice(slice, func(i, j int) bool {
			// If slice[j] appears in rules[slice[i]], then slice[i] must come first
			if deps, exists := rules[slice[i]]; exists {
				for _, dep := range deps {
					if dep == slice[j] {
						return false
					}
				}
			}
			// If slice[i] appears in rules[slice[j]], then slice[j] must come first
			if deps, exists := rules[slice[j]]; exists {
				for _, dep := range deps {
					if dep == slice[i] {
						return true
					}
				}
			}
			// If neither is dependent on the other, sort by value
			return slice[i] < slice[j]
		})
		fixedData = append(fixedData, slice)
	}

	return fixedData
}

func getMidpoint(data [][]int) int {
	midSum := 0
	for _, line := range data {
		mid := math.Floor(float64(len(line)) / 2)
		midSum += line[int(mid)]
	}

	return midSum
}

func main() {
	rulesMap := readRules("rules.txt")
	printings := readPrinting("printing.txt")
	validPrints, invalidPrints := findValidPrints(rulesMap, printings)
	midSum := getMidpoint(validPrints)
	fmt.Println("The sum of the correct data midpoints is ", midSum)
	fixedInvalid := fixInvalid(rulesMap, invalidPrints)
	midSumInvalid := getMidpoint(fixedInvalid)
	fmt.Println("The sum of the fixed data midpoints is ", midSumInvalid)
}
