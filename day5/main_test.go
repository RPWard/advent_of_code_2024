package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindValid(t *testing.T) {
	rules := readRules("testRules.txt")
	fmt.Println("rules")
	for k, v := range rules {
		fmt.Println(k, ": ", v)
	}
	printings := readPrinting("testPrinting.txt")
	fmt.Println("printings")
	for _, line := range printings {
		fmt.Println(line)
	}
	got, invalid := findValidPrints(rules, printings)
	want := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error - got: %v, want: %v", got, want)
	}

	wantInvalid := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	fmt.Println("ValidPrints")
	for _, line := range got {
		fmt.Println(line)
	}
	fmt.Println("InvalidPrints")
	for _, line := range invalid {
		fmt.Println(line)
	}
	if !reflect.DeepEqual(invalid, wantInvalid) {
		t.Errorf("Error - got %v want: %v", invalid, wantInvalid)
	}
}

func TestFindMidSum(t *testing.T) {
	rules := readRules("testRules.txt")
	printings := readPrinting("testPrinting.txt")
	validPrints, invalidPrints := findValidPrints(rules, printings)
	got := getMidpoint(validPrints)
	want := 143
	if got != want {
		t.Errorf("Error - got: %v, want: %v", got, want)
	}
	fixed := fixInvalid(rules, invalidPrints)
	got = getMidpoint(fixed)
	want = 123
	if got != want {
		t.Errorf("Error - got: %v, want: %v", got, want)
	}
}
