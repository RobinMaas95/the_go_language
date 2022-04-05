// Experiment to measure the difference in running time between our potentially inefficient versions
// and tht one that uses strings.Join.

package main

import (
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	//fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	//fmt.Println(s)
}

func echo3() {
	strings.Join(os.Args[1:], " ")
	//fmt.Println(s)
}

func main() {
	echo1()
	echo2()
	echo3()
}
