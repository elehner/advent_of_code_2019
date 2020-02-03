package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func passRange(input string) (int64, int64) {
	vals := strings.Split(input, "-")
	start, _ := strconv.ParseInt(vals[0], 0, 64)
	end, _ := strconv.ParseInt(vals[1], 0, 64)
	return start, end
}

func generateDigits(val int64) []int64 {
	var digits []int64
	for val > 0 {
		digits = append(digits, val%10)
		val /= 10
	}
	// reverse
	for i := 0; i < len(digits)/2; i++ {
		oppI := len(digits) - 1 - i
		digits[i], digits[oppI] = digits[oppI], digits[i]
	}
	return digits
}

func isValidWord(word int64) bool {
	counts := map[int64]int{}
	for word > 0 {
		num := word % 10
		counts[num]++
		// Part 1
		// if counts[num] > 1 {
		// 	return true
		// }
		word /= 10
	}
	for _, val := range counts {
		if val == 2 {
			return true
		}
	}

	return false
}

func recurGetValidWords(initDigits, acc []int64, max int64, currentWord int64) ([]int64, error) {
	lastVal := currentWord % 10
	currentVal := initDigits[0]
	var err error
	if currentVal < lastVal {
		currentVal = lastVal
	}
	for currentVal < 10 {
		newWord := currentWord*10 + currentVal
		if len(initDigits) > 1 {
			acc, err = recurGetValidWords(initDigits[1:], acc, max, newWord)
			if err != nil {
				return acc, err
			}
		} else {
			if newWord > max {
				return acc, errors.New("Hit end")
			} else if isValidWord(newWord) {
				acc = append(acc, newWord)
			}
		}
		currentVal++
	}
	// This should only be used for the first run
	initDigits[0] = 0
	return acc, nil
}

func getValidWords(digits []int64, max int64) []int64 {
	currentVal := digits[0]
	var acc []int64
	var err error
	for currentVal < 10 {
		acc, err = recurGetValidWords(digits[1:], acc, max, currentVal)
		if err != nil {
			break
		}
		currentVal++
	}
	return acc
}

func main() {
	start, end := passRange("145852-616942")
	startArray := generateDigits(start)
	validWords := getValidWords(startArray, end)
	fmt.Println(len(validWords))
}
