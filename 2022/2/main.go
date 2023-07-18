package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

    readFile, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()

	// Define shape score rules
	shape_scores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	// Game rules (key win over value)
	game_rules := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	// Lose game rules (key lose over value)
	lose_game_rules := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	// Map to decrypt input value
	encrypted := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	var total_score2 int
	var total_score int
	var score int
    for _, line := range fileLines {
		// Puzzle Part 1
		play := strings.Split(line, " ")
		p1 := play[0]
		p2 := encrypted[play[1]]

		if p2 == game_rules[p1] {
			// p1 win
			score = 0 + shape_scores[p2]
		} else if p1 == p2 {
			// draw
			score = 3 + shape_scores[p2]
		} else {
			// p2 win
			score = 6 + shape_scores[p2] 
		}

		// fmt.Println("Input:", p1, play[1], "|", "Decrypted:", p1, p2, "|", "Score:", score)
		total_score = total_score + score

		// Puzzle Part 2
		switch play[1] {
			case "X":
				p2 = game_rules[p1]
				score = 0 + shape_scores[p2] 
			case "Y":
				score = 3 + shape_scores[p1]
			case "Z":
				p2 = lose_game_rules[p1]
				score = 6 + shape_scores[p2]
		}

		total_score2 = score + total_score2
	}

	fmt.Println("#1 Total score:", total_score)
	fmt.Println("#2 Total score:", total_score2)
}
