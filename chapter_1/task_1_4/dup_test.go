package main

import (
	"os"
	"testing"
)

func TestDup(t *testing.T) {
	os.Args = append(os.Args, "no_dup_1.txt")
	os.Args = append(os.Args, "no_dup_2.txt")
	os.Args = append(os.Args, "with_dup_1.txt")
	os.Args = append(os.Args, "with_dup_2.txt")

	get := Dup()
	expected := "with_dup_1.txt\nwith_dup_2.txt"

	if get != expected {
		t.Errorf("expected %q but got %q", expected, get)
	}
}
