// Day 2 of Advent of Code
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
	shapeScores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	// Game rules (key win over value)
	gameRules := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	// Lose game rules (key lose over value)
	loseGameRules := map[string]string{
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

	var totalScore2 int
	var totalScore int
	var score int
	for _, line := range fileLines {
		// Puzzle Part 1
		play := strings.Split(line, " ")
		p1 := play[0]
		p2 := encrypted[play[1]]

		if p2 == gameRules[p1] {
			// p1 win
			score = 0 + shapeScores[p2]
		} else if p1 == p2 {
			// draw
			score = 3 + shapeScores[p2]
		} else {
			// p2 win
			score = 6 + shapeScores[p2]
		}

		// fmt.Println("Input:", p1, play[1], "|", "Decrypted:", p1, p2, "|", "Score:", score)
		totalScore = totalScore + score

		// Puzzle Part 2
		switch play[1] {
		case "X":
			p2 = gameRules[p1]
			score = 0 + shapeScores[p2]
		case "Y":
			score = 3 + shapeScores[p1]
		case "Z":
			p2 = loseGameRules[p1]
			score = 6 + shapeScores[p2]
		}

		totalScore2 = score + totalScore2
	}

	fmt.Println("#1 Total score:", totalScore)
	fmt.Println("#2 Total score:", totalScore2)
}
