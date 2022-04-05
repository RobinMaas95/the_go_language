// This file benchmarks the three methods in the file task_1_3.go.

package main

import (
	"os"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	os.Args = append(os.Args, "ping")
	os.Args = append(os.Args, "www.google.de")
	os.Args = append(os.Args, "5")

	main()
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	os.Args = append(os.Args, "ping")
	os.Args = append(os.Args, "www.google.de")
	os.Args = append(os.Args, "5")

	main()
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	os.Args = append(os.Args, "ping")
	os.Args = append(os.Args, "www.google.de")
	os.Args = append(os.Args, "5")

	main()
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
