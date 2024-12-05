package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	splitData := strings.Fields(string(data))
	return splitData, nil
}

func solve(data []string, slices [][][2]int, variants map[string]bool) int {
	count := 0
	rows := len(data)
	cols := len(data[0])

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			for _, slice := range slices {
				word := ""
				valid := true

				for _, offset := range slice {
					dx, dy := offset[0], offset[1]
					newX, newY := x+dx, y+dy

					if newX >= cols || newY >= rows {
						valid = false
						break
					}

					word += string(data[newY][newX])
				}

				if valid && variants[word] {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	slicesOne := [][][2]int{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, // horizontal
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, // vertical
		{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, // diagonal down-right
		{{0, 3}, {1, 2}, {2, 1}, {3, 0}}, // diagonal down-left
	}

	slicesTwo := [][][2]int{
		{{0, 0}, {1, 1}, {2, 2}, {0, 2}, {2, 0}},
	}

	variantsOne := map[string]bool{
		"XMAS": true,
		"SAMX": true,
	}

	variantsTwo := map[string]bool{
		"MASMS": true,
		"SAMSM": true,
		"MASSM": true,
		"SAMMS": true,
	}

	data, _ := readFile("data.txt")
	partOne := solve(data, slicesOne, variantsOne)
	partTwo := solve(data, slicesTwo, variantsTwo)
	fmt.Printf("Part One: Found %d matches\n", partOne)
	fmt.Printf("Part Two: Found %d matches\n", partTwo)
}
