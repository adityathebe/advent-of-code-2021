package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func BenchmarkGetAnswer(b *testing.B) {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		b.Fatal()
	}

	input := strings.Split(strings.TrimSpace(string(f)), "\n")
	getAnswer(input)
}
