package main

import (
	"reflect"
	"strings"
	"testing"
)

var testData string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestReadFile(t *testing.T) {
	got, _ := readFile("dataTest.txt")
	want := strings.Split(testData, "\n")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error! got: %v, want: %v", got, want)
	}
}

func TestCountHorizontal(t *testing.T) {
	data, _ := readFile("dataTest.txt")
	got := countHorizontal(data)
	want := 4
	if got != want {
		t.Errorf("Error count horizontal. got: %v, want: %v", got, want)
	}
}

func TestCountVertical(t *testing.T) {
	data, _ := readFile("dataTest.txt")
	rotatedData := rotateData(data)
	got := countHorizontal(rotatedData)
	want := 2
	if got != want {
		t.Errorf("Error count vertical. got %v, want %v", got, want)
	}
}
