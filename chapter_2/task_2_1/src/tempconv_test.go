package main

import (
	"task_2_1/src/tempconv"
	"testing"
)

func TestCelsiusConversations(t *testing.T) {
	t.Run("convert celsius to fahrenheit", func(t *testing.T) {
		got := tempconv.CToF(50)
		want := tempconv.Fahrenheit(122)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("convert celsius to kelvin", func(t *testing.T) {
		got := tempconv.CToK(50)
		want := tempconv.Kelvin(323.15)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestFahrenheitConversations(t *testing.T) {
	t.Run("convert fahrenheit to celsius", func(t *testing.T) {
		got := tempconv.FToC(50)
		want := tempconv.Celsius(10)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("convert fahrenheit to kelvin", func(t *testing.T) {
		got := tempconv.FToK(50)
		want := tempconv.Kelvin(283.15)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestKelvinConversations(t *testing.T) {

	t.Run("convert kelvin to celsius", func(t *testing.T) {
		got := tempconv.KToC(50)
		want := tempconv.Celsius(-223.15)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("convert kelvin to fahrenheit", func(t *testing.T) {
		got := tempconv.KToF(50)
		want := tempconv.Fahrenheit(-369.67)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
