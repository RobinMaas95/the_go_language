package main

import (
	"math"
	"testing"
)

func TestConversion(t *testing.T) {

	t.Run("convert meter to feet", func(t *testing.T) {
		got := meterToFeet(Meter(50))
		want := Feet(164.042)

		if withinTolerance(want, got, 1e-5) {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("convert feet to meter", func(t *testing.T) {
		got := feetToMeter(Feet(50))
		want := Meter(15.240)

		if !withinTolerance(want, got, 1e-3) {
			t.Errorf("got %f want %f", got, want)
		}
	})
}

func withinTolerance[distance Feet | Meter](a distance, b distance, e float64) bool {
	if a == b {
		return true
	}
	d := math.Abs(float64(a - b))

	if b == 0 {
		return d < e
	}

	return (d / math.Abs(float64(b))) < e

}
