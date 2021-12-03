package main

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkGetAnswer(b *testing.B) {
	f, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	GetAnswer(scanner)
}
