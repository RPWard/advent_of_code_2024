package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	X, Y int
}

type GridData struct {
	OriginalNodes map[string][]Coord
	ExtendedNodes map[string][]Coord
	MaxX, MaxY    int
}

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading file %v: %v", filename, err)
	}
	return string(data)
}

func isWithinBounds(coord Coord, data GridData) bool {
	return coord.X >= 0 && coord.X < data.MaxX && coord.Y >= 0 && coord.Y < data.MaxY
}

func parseGrid(input string) GridData {
	originalMap := make(map[string][]Coord)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	maxY := len(lines)
	maxX := len(lines[0])

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				key := string(char)
				originalMap[key] = append(originalMap[key], Coord{X: x, Y: y})
			}
		}
	}

	data := GridData{
		OriginalNodes: originalMap,
		ExtendedNodes: make(map[string][]Coord),
		MaxX:          maxX,
		MaxY:          maxY,
	}

	for char, coords := range originalMap {
		data.ExtendedNodes[char] = calculateDoubleDistanceNodes(coords, data)
	}

	return data
}

func calculateDoubleDistanceNodes(coords []Coord, data GridData) []Coord {
	seen := make(map[Coord]bool)
	var extendedNodes []Coord

	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			node1 := coords[i]
			node2 := coords[j]

			dx := node2.X - node1.X
			dy := node2.Y - node1.Y

			candidates := []Coord{
				{X: node1.X - dx, Y: node1.Y - dy},
				{X: node2.X + dx, Y: node2.Y + dy},
			}

			for _, node := range candidates {
				if isWithinBounds(node, data) && !seen[node] {
					seen[node] = true
					extendedNodes = append(extendedNodes, node)
				}
			}
		}
	}
	return extendedNodes
}

func calculateLineBasedNodes(data GridData) int {
	allPatternPoints := make(map[Coord]bool)

	for _, origNodes := range data.OriginalNodes {
		for i := 0; i < len(origNodes); i++ {
			for j := i + 1; j < len(origNodes); j++ {
				p1 := origNodes[i]
				p2 := origNodes[j]

				for y := 0; y < data.MaxY; y++ {
					for x := 0; x < data.MaxX; x++ {
						p3 := Coord{X: x, Y: y}
						if areCollinear(p1, p2, p3) {
							allPatternPoints[p3] = true
						}
					}
				}
			}
		}
	}

	return len(allPatternPoints)
}

func areCollinear(p1, p2, p3 Coord) bool {
	return (p2.Y-p1.Y)*(p3.X-p2.X)-(p2.X-p1.X)*(p3.Y-p2.Y) == 0
}

func countUniqueNodes(nodeMap map[string][]Coord) int {
	seen := make(map[Coord]bool)
	for _, coords := range nodeMap {
		for _, coord := range coords {
			seen[coord] = true
		}
	}
	return len(seen)
}

func main() {
	input := readFile("data.txt")
	data := parseGrid(input)

	fmt.Printf("Problem 1 - Double distance nodes: %d\n",
		countUniqueNodes(data.ExtendedNodes))

	fmt.Printf("Problem 2 - Line-based pattern nodes: %d\n",
		calculateLineBasedNodes(data))
}
