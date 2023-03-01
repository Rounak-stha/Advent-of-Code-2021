package main

import (
	"AOC/file_paths"
	"AOC/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(file_paths.Day_2)
	utils.HandleError(err)
	scanner := bufio.NewScanner(file)

	fmt.Println("Part 1: ", part1(*scanner))
	fmt.Println("Part 2: ", part2(*scanner))
}

func part1(scanner bufio.Scanner) int {
	var h, v int

	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		steps, err := strconv.Atoi(command[1])
		utils.HandleError(err)

		switch command[0] {
		case "forward":
			h += steps
		case "down":
			v += steps
		case "up":
			v -= steps
		}
	}

	return h * v
}

func part2(scanner bufio.Scanner) int {
	var h, v, a int

	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		steps, err := strconv.Atoi(command[1])
		utils.HandleError(err)

		switch command[0] {
		case "forward":
			h += steps
			v += a * steps
		case "down":
			a += steps
		case "up":
			a -= steps
		}
	}

	return h * v
}

