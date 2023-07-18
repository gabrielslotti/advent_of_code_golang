// Day 1 of Advent of Code
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	var greatestAmount int
	var pastElf int
	var currentElf int
	var amountSlice []int

	for _, line := range fileLines {
		switch line {
		case "":
			if currentElf > pastElf && currentElf > greatestAmount {
				greatestAmount = currentElf
			}

			amountSlice = append(amountSlice, currentElf)

			pastElf = currentElf
			currentElf = 0
			continue
		default:
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			currentElf = currentElf + calories
		}
	}

	sort.Ints(amountSlice)
	top3 := amountSlice[len(amountSlice)-3:]

	var total int
	for _, value := range top3 {
		total = total + value
	}

	fmt.Println("#1 Amswer:", greatestAmount, "calories.")
	fmt.Println("#2 Answer:", total, "calories.")
}
