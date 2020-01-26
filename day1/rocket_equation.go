package main

import (
	"bufio"
	"os"
	"strconv"
)

func parseFile() []int64 {
	file, _ := os.Open("./input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int64
	for scanner.Scan() {
		number, _ := strconv.ParseInt(scanner.Text(), 0, 64)
		numbers = append(numbers, number)
	}

	return numbers
}

func main() {
	numbers := parseFile()

	var totalFuelRequired int64
	for _, number := range numbers {
		fuelRequired := (number / 3) - 2
		for fuelRequired > 0 {
			totalFuelRequired += fuelRequired
			fuelRequired = (fuelRequired / 3) - 2
		}
	}

	println(totalFuelRequired)
}
