package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readData(filename string) map[int][]int {
	data, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error reading file %s: %v", filename, err)
	}

	var numMap = make(map[int][]int)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var values []int
		line := strings.Split(scanner.Text(), ":")
		key, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatalf("Error converting key to int: %s - %v", line[0], err)
		}
		for _, v := range strings.Fields(line[1]) {
			intVal, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Error converting value to int: %s = %v", v, err)
			}
			values = append(values, intVal)
		}
		numMap[key] = values
	}
	return numMap
}

func generateOperatorCombos(operators []string, slots int) chan []string {
	ch := make(chan []string)

	go func() {
		defer close(ch)

		current := make([]string, slots)

		var generate func(position int)
		generate = func(position int) {
			if position == slots {
				combo := make([]string, slots)
				copy(combo, current)
				ch <- combo
				return
			}

			for _, op := range operators {
				current[position] = op
				generate(position + 1)
			}
		}

		generate(0)
	}()
	return ch
}

func evaluate(numbers []int, operators []string) float64 {
	result := float64(numbers[0])
	lastIntResult := numbers[0] // keep track of integer result for concatenation

	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += float64(numbers[i+1])
			lastIntResult = int(result)
		case "-": // didn't end up needing this
			result -= float64(numbers[i+1])
			lastIntResult = int(result)
		case "*":
			result *= float64(numbers[i+1])
			lastIntResult = int(result)
		case "/": // didn't end up needing this
			if numbers[i+1] != 0 {
				result /= float64(numbers[i+1])
				lastIntResult = int(result)
			}
		case "||":
			lastIntResult = concatenate(lastIntResult, numbers[i+1])
			result = float64(lastIntResult)
		}
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(abs(n)))) + 1
}

func concatenate(a, b int) int {
	isNegative := a < 0
	if isNegative {
		a = abs(a)
	}

	b = abs(b)
	digits := countDigits(b)
	result := a*int(math.Pow10(digits)) + b

	if isNegative {
		return -result
	}
	return result
}

func solve(data map[int][]int, operators []string) int {
	var answer int
	for k, v := range data {
		for ops := range generateOperatorCombos(operators, len(v)-1) {
			result := int(evaluate(v, ops))
			if result == k {
				answer += result
				goto nextCombo
			}
		}
	nextCombo:
	}
	return answer
}

func main() {
	data := readData("data.txt")
	operators := []string{"+", "*"}
	concatOperators := []string{"+", "*", "||"}
	answer := solve(data, operators)
	concatAnswer := solve(data, concatOperators)
	fmt.Println("The answer is ", answer)
	fmt.Println("The answer is ", concatAnswer)
}
