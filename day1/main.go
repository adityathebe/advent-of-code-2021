package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetAnswer(scanner *bufio.Scanner) int {
	var out int
	var initial = true
	var increased int
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		if x > out && !initial {
			increased++
		}
		initial = false
		out = x
	}

	return increased
}

func run() error {
	f, err := os.Open("measurements.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	answer := GetAnswer(scanner)
	log.Println(answer)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
