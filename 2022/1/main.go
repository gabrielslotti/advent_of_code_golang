package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
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

	var greatest_amount int
	var past_elf int
	var current_elf int
	var amount_slice []int

    for _, line := range fileLines {
		switch line {
			case "":
				if current_elf > past_elf && current_elf > greatest_amount {
					greatest_amount = current_elf
				}

				amount_slice = append(amount_slice, current_elf)

				past_elf = current_elf
				current_elf = 0
				continue
			default:
				calories, err := strconv.Atoi(line)
				if err != nil {
					panic(err)
				}

				current_elf = current_elf + calories
		}
    }

	sort.Ints(amount_slice)
	top3 := amount_slice[len(amount_slice)-3:]

	var total int
	for _, value := range top3 {
		total = total + value	
  	}

	fmt.Println("#1 Amswer:", greatest_amount, "calories.")
	fmt.Println("#2 Answer:", total, "calories.")
}

