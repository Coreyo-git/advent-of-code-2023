package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Init sum of all calibration values
	var sum int

	// Open file containing calibration data
	file, err := os.Open("calibrations.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop each line in the file
	for scanner.Scan() {
		// store first and last numbers as strings
		var first, last string

		// Get current line from file
		line := scanner.Text()

		// loop through each char in line
		for _, char := range line {
			// check if is a number and updates first and last
			if isNumber(string(char)) {
				// if first is empty, assign it and last to char
				if first == "" {
					first = string(char)
					last = string(char)
				} else {
					// else update last
					last = string(char)
				}
			}
		}

		// Convert first and last numbers from strings to an int
		i, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		} else {
			// Print the current num and update the sum with the current calibration value
			fmt.Print("Adding: " + strconv.Itoa(i) + "\n")
			sum += i
		}

		// fmt.Print("Current Word: " + first + last + "\n")
	}

	// Print the total sum of all calibration values
	fmt.Print("Sum: " + strconv.Itoa(sum) + "\n")

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// isNumber checks if a given string represents a number
func isNumber(c string) bool {
	if _, err := strconv.Atoi(string(c)); err == nil {
		return true
	}
	return false
}
