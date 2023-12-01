package cube

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		// store game id
		var gameId string

		// Get current line from file
		line := scanner.Text()

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


// Sums two integers
func Add(x, y int) (res int) {
	return x + y
}
