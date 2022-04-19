// Mf converts its numeric argument to meter and feet
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// meterToFeet takes a distance in meter an returns the same distance in feet.
func meterToFeet(input Meter) Feet {
	return Feet(input * 3.2808)
}

// feetToMeter takes a distance in feet and returns the same distance in meter.
func feetToMeter(input Feet) Meter {
	return Meter(input / 3.2808)
}

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%.2f Meter", m) }
func (f Feet) String() string  { return fmt.Sprintf("%.2f Feet", f) }

// getUserInput prompts the user for input. The input is returned as a list
// of strings (input splitted at whitespaces.)
func getUserInput() []string {
	var distances []string
	fmt.Println("Enter your value(s) (seperated by space):")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		distances = strings.Fields(line)
	}

	return distances
}

// getDistanceLists returns a list of distances (strings). If arguments are passed,
// they are returned as a list, otherwise the user is prompted for values.
func getDistanceList() (distances []string) {
	distances = os.Args[1:]
	if len(distances) == 0 {
		distances = getUserInput()
	}
	return
}

// convertStrToFloat converts the passed string to a float64 and handles a possible error.
func convertStrToFloat(s string) float64 {
	d, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	return d
}

// printOutput converts distance in feet and meter and prints the vice versa conversation.
func printOutput(distance string) {
	d := convertStrToFloat(distance)
	m := Meter(d)
	f := Feet(d)
	fmt.Printf("%s = %s, %s = %s\n", m, meterToFeet(m), f, feetToMeter(f))
}

func main() {
	distances := getDistanceList()
	for _, arg := range distances {
		printOutput(arg)
	}
}
