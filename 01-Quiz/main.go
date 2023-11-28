package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Text()
		println(line)

		// split the line in question and answer
		question, answer := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
	}
}
