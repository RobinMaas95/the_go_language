package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dup() (result string) {
	counts := make(map[string]int)
	//fileMap := make(map[string]string)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		countLines(f, counts, arg)
	}
	for line, n := range counts {
		if n > 1 {
			t := strings.Split(line, "-")[1] + "\n"
			result += t
			fmt.Printf(t)
		}
	}
	return strings.Trim(result, "\n")
}

func countLines(f *os.File, counts map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()+"-"+filename]++
	}
}

func main() {
	_ = Dup()
}
