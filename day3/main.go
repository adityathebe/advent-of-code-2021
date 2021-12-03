package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toBinary(num []int) (string, string) {
	var binary string
	var compliment string
	for i := range num {
		if num[i] > 0 {
			binary += "1"
			compliment += "0"
		} else if num[i] < 0 {
			binary += "0"
			compliment += "1"
		}
	}

	return binary, compliment
}

func GetAnswer(input []string) (int64, error) {
	output := make([]int, len(input[0]))
	for _, bin := range input {
		for i, c := range bin {
			if c == '0' {
				output[i]--
			} else if c == '1' {
				output[i]++
			}
		}
	}

	gamma, epsilon := toBinary(output)

	g, err := strconv.ParseInt(gamma, 2, 32)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseInt(%s); %w", gamma, err)
	}

	e, err := strconv.ParseInt(epsilon, 2, 32)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseInt(%s); %w", epsilon, err)
	}

	return g * e, nil
}

func run() error {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return fmt.Errorf("os.ReadFile(); %w", err)
	}

	input := strings.Split(strings.TrimSpace(string(content)), "\n")
	answer, err := GetAnswer(input)
	if err != nil {
		return fmt.Errorf("GetAnswer(); %w", err)
	}

	fmt.Printf("Answer = %d\n", answer)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
