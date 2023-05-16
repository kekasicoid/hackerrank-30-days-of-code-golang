package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var incomingInt int64
	var incomingFloat float64
	var incomingString string
	inputVals := []string{}

	// Read and save an integer, double, and String to your variables.
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}
		inputVals = append(inputVals, input)
	}
	incomingInt, _ = strconv.ParseInt(inputVals[0], 10, 64)
	incomingFloat, _ = strconv.ParseFloat(inputVals[1], 64)
	incomingString = inputVals[2]

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + uint64(incomingInt))

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+incomingFloat)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + incomingString)
}
