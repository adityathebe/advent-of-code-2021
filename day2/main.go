package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Forward = "forward "
	Down    = "down "
	Up      = "up "
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}

	input := strings.Split(strings.TrimSpace(string(f)), "\n")
	result := getAnswer(input)
	log.Println(result)
	return nil
}

func getAnswer(input []string) int {
	var (
		depth int
		hor   int
	)

	for i := range input {
		if strings.Index(input[i], "f") == 0 {
			num, _ := strconv.Atoi(strings.ReplaceAll(input[i], Forward, ""))
			hor += num
		} else if strings.Index(input[i], "u") == 0 {
			num, _ := strconv.Atoi(strings.ReplaceAll(input[i], Up, ""))
			depth -= num
		} else if strings.Index(input[i], "d") == 0 {
			num, _ := strconv.Atoi(strings.ReplaceAll(input[i], Down, ""))
			depth += num
		}
	}

	return hor * depth
}
