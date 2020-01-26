package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseFile() []int64 {
	file, _ := os.Open("./input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int64
	for scanner.Scan() {
		numbersString := scanner.Text()
		for _, numberString := range strings.Split(numbersString, ",") {
			number, _ := strconv.ParseInt(numberString, 0, 64)
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func parseIntcode(noun, verb int64, instructionSet []int64) int64 {
	instructions := make([]int64, len(instructionSet))
	copy(instructions, instructionSet)
	instructions[1] = noun
	instructions[2] = verb
	head := 0
	for instructions[head] != 99 {
		val1 := instructions[instructions[head+1]]
		val2 := instructions[instructions[head+2]]
		saveLocation := instructions[head+3]

		switch currentInstruction := instructions[head]; currentInstruction {
		case 1:
			instructions[saveLocation] = val1 + val2
		case 2:
			instructions[saveLocation] = val1 * val2
		default:
			panic("not a known command")
		}

		head += 4
	}
	return instructions[0]
}

func main() {
	instructions := parseFile()
	result := parseIntcode(12, 2, instructions)
	println(result)

	var goal int64 = 19690720
	var noun int64
	var verb int64
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			if parseIntcode(noun, verb, instructions) == goal {
				println(100*noun + verb)
				return
			}
		}
	}
}
