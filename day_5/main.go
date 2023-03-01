package main

import (
	"AOC/file_paths"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end Point
}

// each line is made of 2 point: Start and End; Each Point has x and y coordinate
// Line { start: Point {x1, y1}, End:Point { x2, y2} }
func parseData (data string) []Line {
	var lines []Line
	dataSlice := strings.Split(data, "\r\n")
	for _, line := range dataSlice {
		points := strings.Split(line, " -> ")
		point1 := strings.Split(points[0], ",")
		point2 := strings.Split(points[1], ",")
		x1, _ := strconv.Atoi(point1[0])
		y1, _ := strconv.Atoi(point1[1])
		x2, _ := strconv.Atoi(point2[0])
		y2, _ := strconv.Atoi(point2[1])
		lines = append(lines, Line{ Point{x1, y1}, Point{x2, y2} })
	}
	return lines
}

func main() {
	input, _ := os.ReadFile(file_paths.Day_5)
	parsedData := parseData(string(input)) // [... { start {x1, y1}, end {x2, y2} } ...]
	fmt.Println("part 1: ", main_game(parsedData, 1))
	fmt.Println("part 2: ", main_game(parsedData, 2))
}

// { 3, 2 } -> { 7, 2 }

func main_game(data []Line, part int8) int32 {
	var overlapCount int32
	pointOverlapCount := make(map[string]int)

	for _, line := range data {
		x1 := line.start.x
		y1 := line.start.y
		x2 := line.end.x
		y2 := line.end.y

		var start, end Point

		// Vertical Line
		if x1 == x2 {
			// movement style: top - bottom
			if y1 < y2 {
				start = line.start
				end = line.end
			} else {
				start = line.end
				end=  line.start
			}

			move(start, end, "down", pointOverlapCount)

		} else if y1 == y2 { // Horizontal Line
			// movement style: left - right
			if x1 < x2 {
				start = line.start
				end = line.end
			} else {
				start = line.end
				end = line.start
			}

			move(start, end, "right", pointOverlapCount)
		} else if part == 2 { // go diagonal
			// pick one style of Y movement
			// picking: top -> bottom
			if y1 < y2 {
				start = line.start
				end = line.end
			} else {
				start = line.end
				end = line.start
			}

			// check if movement should be towards right or left
			var moveTo string
			if start.x < end.x {
				moveTo = "diag_right"
			} else {
				moveTo = "diag_left"
			}
			// We always move from top - bottom
			// so check if we reached the correct bottom level
			move(start, end, moveTo, pointOverlapCount)
		}
	}
	for _, v := range pointOverlapCount {
		if (v >= 2) { overlapCount++ }
	}
	return overlapCount
}

func move(start, end Point, move string, pointOverlapCount map[string]int) {
	curr := start
	// for horizontal movement, we loop until we reached the last x coordinate towards right
	// for vertical movement, we loop until we reached the last y coordinate downwards
	// for diagonal movement, we fix the downward movement and move either right or left; 
	// that means, we loop just like the vertical movement
	var loop_start *int
	var loop_end int

	// we move down for both vertical and diagonal movement
	// else we move completely towards right
	// choose the startin and ending point for loop
	if move == "right" {
		loop_start = &curr.x
		loop_end = end.x
	} else {
		loop_start = &curr.y
		loop_end = end.y
	}
	
	for *loop_start <= loop_end {
		point := strconv.Itoa(curr.x) + "," + strconv.Itoa(curr.y)
		if _, ok := pointOverlapCount[point]; ok {
			pointOverlapCount[point] += 1
		} else {
			pointOverlapCount[point] = 1
		}
		step(&curr, move)
	} 
}

func step(curr *Point, move string) {
	switch move {
	case "down":
		curr.y += 1
	case "right":
		curr.x += 1
	case "diag_right":
		curr.x += 1
		curr.y += 1
	case "diag_left":
		curr.x -= 1
		curr.y += 1
	default:
		panic("Unsupported Move Step Requested")
	}
}



