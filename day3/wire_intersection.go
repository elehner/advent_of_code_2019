package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y float64
}

func (p *point) copy() point {
	return point{p.x, p.y}
}

func (p *point) costToPoint(p2 *point) float64 {
	return math.Abs(p.x-p2.x) + math.Abs(p.y-p2.y)
}

func isBetween(start, end, check float64) bool {
	return (start <= check && check <= end) || (end <= check && check <= start)
}

func intersection(p1, p2, p3, p4 *point) *point {
	// Assuming that:
	// 1. For each line, only 1 coordinate changes at a time
	// 2. The lines never overlap
	var res *point
	if p1.x == p2.x {
		if isBetween(p3.x, p4.x, p1.x) && isBetween(p1.y, p2.y, p3.y) {
			res = &point{p1.x, p3.y}
		}
	} else {
		if isBetween(p3.y, p4.y, p1.y) && isBetween(p1.x, p2.x, p3.x) {
			res = &point{p3.x, p1.y}
		}
	}

	return res
}

func parseFile() [][]string {
	file, _ := os.Open("./input")
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
		distance, _ := strconv.ParseFloat(direction[1:], 64)
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
	var minDistance float64 = -1
	var minCost float64 = -1
	var costToPoint1 float64
	for l1 := 1; l1 < len(line1); l1++ {
		p1 := line1[l1-1]
		p2 := line1[l1]

		var costToPoint2 float64
		for l2 := 1; l2 < len(line2); l2++ {
			p3 := line2[l2-1]
			p4 := line2[l2]
			if l2 > 1 {
				costToPoint2 += line2[l2-1].costToPoint(&line2[l2-2])
			}

			intersect := intersection(&p1, &p2, &p3, &p4)
			if intersect == nil {
				continue
			}
			if intersect.x == 0 && intersect.y == 0 {
				continue
			}

			distance := math.Abs(intersect.x) + math.Abs(intersect.y)
			if minDistance < 0 || distance < minDistance {
				minDistance = distance
			}

			costToStartingPoint := costToPoint1 + costToPoint2
			costToIntersection := costToStartingPoint + p1.costToPoint(intersect) + p3.costToPoint(intersect)
			if minCost < 0 || costToIntersection < minCost {
				minCost = costToIntersection
			}
		}
		costToPoint1 += p1.costToPoint(&p2)
	}
	fmt.Println(minDistance)
	fmt.Println(minCost)
}
