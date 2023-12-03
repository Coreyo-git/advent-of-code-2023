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
		
		// sets game id string equal to parts [0]
		gameIDString := parts[0]

		// get game ID as int from game id string
		gameID := getStringAsInt(gameIDString)

		// debugging print of game id
		// fmt.Println("Game ID: ", gameID)

		// sets game as the game section of the split
		game := parts[1]

		if gameIsPossible(game) {
			sumOfPossibleGames = Add(sumOfPossibleGames, gameID)
		}
	}

	fmt.Println("Sum of possible games: ", sumOfPossibleGames)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func gameIsPossible(game string) bool {

	// split each draw in the game
	draws := strings.Split(game, ";")

	// loop through each draw and check if the draw is within bounds
	for _, draw := range draws {
		cubes := strings.Split(draw, ",")
		for _, cube := range cubes {
			if !isDrawWithinBounds(cube) {
				return false
			}
		}
	}

	// return true if all draws are within bounds
	return true
}

func isDrawWithinBounds(draw string) bool {
	// trim whitespace
	draw = strings.TrimSpace(draw)
	
	// seperate draw by space character
	parts := strings.Split(draw, " ")

	if len(parts) != 2 {
		fmt.Println("Error getting game ID from line:", draw)
		return false
	}

	// assign num and color sections
	num := getStringAsInt(parts[0])
	color := parts[1]

	switch color {
	case "red":
		if num > redMax{ return false }
	case "blue":
		if num > blueMax{ return false }
	case "green":
		if num > greenMax{ return false }
	default:
		fmt.Print("No Color found in draw evaluation")
		return false
	}

	// if it passes all cases return true
	return true
}

// returns an int of from a string
func getStringAsInt(String string) int {
	// Regex to match the numeric in string
	numericRegex := regexp.MustCompile("[0-9]+")

	numericMatches := numericRegex.FindAllString(String, -1)
	
	if len(numericMatches) == 0 {
		fmt.Println("No numeric found in:", String)
		return 0
	}

	gameIDInt, err := strconv.Atoi(string(numericMatches[0]))
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
