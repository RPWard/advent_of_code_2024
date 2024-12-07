package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Guard struct {
	pos Coord
	dir Direction
}

type Coord struct {
	x, y int
}

type MapData struct {
	obstacles []Coord
	guard     Guard
	rowsMax   int
	colsMax   int
	visited   map[Coord][]Direction
}

func readFile(filename string) MapData {
	data, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var mapData MapData
	y := 0
	rowLen := 0

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		if rowLen == 0 {
			rowLen = len(string(scanner.Text()))
		}
		for x, v := range scanner.Text() {
			switch string(v) {
			case "^":
				mapData.guard = Guard{pos: Coord{x: x, y: y}, dir: North}
			case "#":
				mapData.obstacles = append(mapData.obstacles, Coord{x: x, y: y})
			}
		}
		y += 1
	}

	mapData.rowsMax = y
	mapData.colsMax = rowLen
	return mapData
}

func checkBounds(md *MapData) bool {
	if md.guard.pos.x < 0 || md.guard.pos.x > md.colsMax-1 || md.guard.pos.y < 0 || md.guard.pos.y > md.rowsMax-1 {
		return false
	}
	return true
}

func hasObstacle(coord Coord, obstacles []Coord) bool {
	return slices.Contains(obstacles, coord)
}

func turnGuard(dir Direction) Direction {
	switch dir {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	return North
}

func getNextPosition(g Guard) Coord {
	switch g.dir {
	case North:
		return Coord{g.pos.x, g.pos.y - 1}
	case East:
		return Coord{g.pos.x + 1, g.pos.y}
	case South:
		return Coord{g.pos.x, g.pos.y + 1}
	case West:
		return Coord{g.pos.x - 1, g.pos.y}
	}
	return g.pos
}

func countGuardSteps(md *MapData) (int, bool) {
	if md.visited == nil {
		md.visited = make(map[Coord][]Direction)
	}
	md.visited[md.guard.pos] = append(md.visited[md.guard.pos], md.guard.dir)

	for {
		nextPos := md.guard.pos
		switch md.guard.dir {
		case North:
			nextPos.y--
		case East:
			nextPos.x++
		case South:
			nextPos.y++
		case West:
			nextPos.x--
		}

		if nextPos.x < 0 || nextPos.x >= md.colsMax ||
			nextPos.y < 0 || nextPos.y >= md.rowsMax {
			return len(md.visited), false
		}

		if hasObstacle(nextPos, md.obstacles) {
			md.guard.dir = turnGuard(md.guard.dir)
			continue
		}
		md.guard.pos = nextPos

		if dirs, exists := md.visited[md.guard.pos]; exists {
			if slices.Contains(dirs, md.guard.dir) {
				return len(md.visited), true
			}
		}

		md.visited[md.guard.pos] = append(md.visited[md.guard.pos], md.guard.dir)
	}
}

func getLoops(md *MapData) int {
	initialGuard := md.guard
	countGuardSteps(md)

	originalPath := make([]Coord, 0)
	for pos := range md.visited {
		originalPath = append(originalPath, pos)
	}

	count := 0
	for _, pos := range originalPath {
		if hasObstacle(pos, md.obstacles) {
			continue
		}
		if pos == initialGuard.pos {
			continue
		}

		tempMd := MapData{
			obstacles: append(slices.Clone(md.obstacles), pos),
			guard: Guard{
				pos: initialGuard.pos,
				dir: North,
			},
			rowsMax: md.rowsMax,
			colsMax: md.colsMax,
			visited: make(map[Coord][]Direction),
		}

		_, foundCycle := countGuardSteps(&tempMd)
		if foundCycle {
			// fmt.Printf("Found cycle when blocking position: %v\n", pos) // Debug print
			count++
		}
	}

	return count
}

func main() {
	mapData := readFile("data.txt")
	mapData.guard.dir = North

	steps, _ := countGuardSteps(&mapData)
	fmt.Printf("Part 1: Number of steps is %d\n", steps)

	mapData = readFile("data.txt")
	mapData.guard.dir = North
	loops := getLoops(&mapData)
	fmt.Printf("Part 2: Number of possible loops is %d\n", loops)
}
