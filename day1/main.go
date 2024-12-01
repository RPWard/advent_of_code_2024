package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func toSlices(data []string) ([][]int, error) {
	var slice1 []int
	var slice2 []int
	for i, v := range data {
		num, err := strconv.Atoi(v)
		if err != nil {
			return [][]int{}, err
		}
		if i%2 == 0 {
			slice2 = append(slice2, num)
		} else {
			slice1 = append(slice1, num)
		}
	}

	return [][]int{slice1, slice2}, nil
}

func sortSlices(dataSlices [][]int) [][]int {
	sort.Ints(dataSlices[0])
	sort.Ints(dataSlices[1])
	return dataSlices
}

func getDistance(dataSlices [][]int) int {
	var total int
	for i := range dataSlices[0] {
		diff := dataSlices[0][i] - dataSlices[1][i]
		if diff < 0 {
			diff = diff * -1
		}
		total += diff
	}
	return total
}

func getSimilarity(dataSlices [][]int) int {
	var total int
	for _, v := range dataSlices[0] {
		timesFound := 0
		for _, z := range dataSlices[1] {
			if v == z {
				timesFound += 1
			}
		}
		total += v * timesFound
	}
	return total
}

func main() {
	data, err := getData("data.txt")
	if err != nil {
		log.Fatalf("Error getting data: %v", err)
	}

	splitData := strings.Fields(data)
	dataSlices, err := toSlices(splitData)
	if err != nil {
		log.Fatalf("Failed to split numbers to slices: %v", err)
	}

	sortedData := sortSlices(dataSlices)
	total := getDistance(sortedData)
	fmt.Printf("The total of differences between the two lists is %v\n", total)

	similarity := getSimilarity(dataSlices)
	fmt.Printf("The total similarity score is %v\n", similarity)
}
