package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	var lines []string

	for scan.Scan() {
		line := scan.Text()
		lines = append(lines, line)
	}

	for _, line := range lines {
		// split the line in question and answer
		question, answer := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		var input string
		fmt.Printf("What is %s?\n", question)
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if answer == input {
			fmt.Println("Yes!")
		} else {
			fmt.Println("No!")
		}

	}
}
