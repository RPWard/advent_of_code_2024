package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	mapData := readFile("testData.txt")
	targetGuardX := 4
	targetGuardY := 6
	numObstacles := 8
	targetRows := 10
	targetCols := 10
	targetSteps := 41
	targetLoops := 6
	/*
		targetNewCoords := []Coord{
			{3, 6},
			{6, 7},
			{7, 7},
			{1, 8},
			{3, 8},
			{7, 9},
		}
	*/
	if mapData.guard.pos.x != targetGuardX {
		t.Errorf("GuardX incorrect: got: %v want: %v", mapData.guard.pos.x, targetGuardX)
	}
	if mapData.guard.pos.y != targetGuardY {
		t.Errorf("GuardY incorrect: got: %v want: %v", mapData.guard.pos.y, targetGuardY)
	}
	if len(mapData.obstacles) != numObstacles {
		t.Errorf("Wrong number of obstacles: got: %v, want: %v, coords: %v", len(mapData.obstacles), numObstacles, mapData.obstacles)
	}
	if mapData.rowsMax != targetRows {
		t.Errorf("Wrong MaxRows: got %v, want %v", mapData.rowsMax, targetRows)
	}
	if mapData.colsMax != targetCols {
		t.Errorf("Wrong MaxCols: got %v, want %v", mapData.colsMax, targetCols)
	}

	steps, _ := countGuardSteps(&mapData)
	if steps != targetSteps {
		t.Errorf("Wrong number of steps! got: %v, want %v", steps, targetSteps)
	}
	mapData = readFile("testData.txt")
	mapData.guard.dir = North
	loops := getLoops(&mapData)
	if loops != targetLoops {
		t.Errorf("Wrong number of loops! got: %v, want: %v", loops, targetLoops)
	}

}
