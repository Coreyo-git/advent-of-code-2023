package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
)

var redMax = 12
var greenMax = 13
var blueMax = 14
var sumOfPossibleGames = 0

func main() {
	// Open file containing games data
	file, err := os.Open("games.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop each line in the file
	for scanner.Scan() {
		// Get current line from file
		line := scanner.Text()
		
		// Split the line by colon to get the game ID
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Error getting game ID from line:", line)
			continue
		}

		gameID := getGameIDAsInt(parts[0])

		fmt.Println("Game ID: ", gameID)
		break
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func gameIsPossible(game string) int {
	var possibleGame = 0

	

	return possibleGame
}

// returns an int of the gameID from the gameID String
func getGameIDAsInt(gameIDString string) int {
	// Regex to match the numeric in string
	numericRegex := regexp.MustCompile("[0-9]+")

	gameIDMatches := numericRegex.FindAllString(gameIDString, -1)
	
	if len(gameIDMatches) == 0 {
		fmt.Println("No numeric found in:", gameIDString)
		return 0
	}

	gameIDInt, err := strconv.Atoi(string(gameIDMatches[0]))
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		return gameIDInt
	}
}


// Sums two integers
func Add(x, y int) (res int) {
	return x + y
}
