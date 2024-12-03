package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readData(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	splitData := strings.Split(string(data), "\n")
	return splitData, nil
}

func convertToInts(data []string) ([][]int, error) {
	var formattedReports [][]int

	for _, report := range data {
		formattedReport := strings.Fields(report)
		var intReport []int

		for _, value := range formattedReport {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return [][]int{}, err
			}
			intReport = append(intReport, intValue)
		}
		formattedReports = append(formattedReports, intReport)
	}
	return formattedReports, nil
}

func checkLevel(diff, level int) (bool, int) {
	var currLevel int
	if diff < 0 {
		currLevel = -1
	} else if diff > 0 {
		currLevel = 1
	} else {
		fmt.Println("Failed same number")
		return false, currLevel
	}

	if level != 0 && level != currLevel {
		fmt.Println("Failed changed +/-")
		return false, currLevel
	}
	return true, currLevel
}

func checkAdjacent(diff int) bool {
	type Empty struct{}

	check := map[int]Empty{
		-3: Empty{},
		-2: Empty{},
		-1: Empty{},
		1:  Empty{},
		2:  Empty{},
		3:  Empty{},
	}

	_, ok := check[diff]
	if !ok {
		return false
	}
	return true
}

func testData(data []int) []bool {
	prev := -1
	level := 0
	var result []bool

	for _, v := range data {
		fmt.Println(v)
		if prev == -1 {
			prev = v
			continue
		}

		diff := v - prev
		ok1 := checkAdjacent(diff)
		ok2, nextLevel := checkLevel(diff, level)
		level = nextLevel
		prev = v
		result = append(result, ok1 && ok2)
	}
	fmt.Println(result)
	return result
}

func numberFalse(data []bool) int {
	numFalse := 0
	for _, v := range data {
		if !v {
			numFalse += 1
		}
	}
	return numFalse
}

func countSafe(data []int) int {
	result := testData(data)
	numFalse := numberFalse(result)
	if numFalse == 0 {
		return 1
	}
	return 0
}

func countSafeLoop(data [][]int) (int, int) {
	safe := 0
	tolerantSafe := 0
	for _, report := range data {
		if len(report) == 0 {
			continue
		}

		result := countSafe(report)
		if result == 1 {
			safe += result
			tolerantSafe += result
		} else {
			for i, _ := range report {
				newSlice := slices.Clone(report)
				newSlice = slices.Delete(newSlice, i, i+1)
				if x := countSafe(newSlice); x == 1 {
					tolerantSafe += x
					break
				}
			}
		}
	}
	return safe, tolerantSafe
}

func main() {
	data, err := readData("data.txt")
	if err != nil {
		log.Fatalf("Error occured reading file: %v", err)
	}

	formattedData, err := convertToInts(data)
	if err != nil {
		log.Fatalf("Error occured converting data: %v", err)
	}

	safe, tolerantSafe := countSafeLoop(formattedData)
	fmt.Printf("The number of safe reports is %v\n", safe)
	fmt.Printf("The number of tolerant safe reports is %v\n", tolerantSafe)
}
