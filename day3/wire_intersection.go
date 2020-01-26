package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int64
}

func (p *point) copy() point {
	return point{p.x, p.y}
}

func (p *point) intersection(p2 *point) (point, error) {
	return point{p.x, p.y}, nil
}

func parseFile() [][]string {
	file, _ := os.Open("./test")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		lines = append(lines, line)
	}

	return lines
}

func drawLine(directions []string) []point {
	line := []point{point{0, 0}}
	point := point{0, 0}
	for _, direction := range directions {
		distance, _ := strconv.ParseInt(direction[1:], 0, 64)
		switch direction[0] {
		case 'R':
			point.x += distance
		case 'U':
			point.y += distance
		case 'L':
			point.x -= distance
		case 'D':
			point.y -= distance
		default:
			panic("not a known direction")
		}
		line = append(line, point.copy())
	}

	return line
}

func main() {
	lines := parseFile()
	line1 := drawLine(lines[0])
	line2 := drawLine(lines[1])
	fmt.Println(line1)
	fmt.Println(line2)
}
