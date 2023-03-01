package main

import (
	"AOC/file_paths"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PatternAndOutput struct {
	pattern []string
	output []string
}

type Input []PatternAndOutput

var NUMBERS_AND_ITS_UNIQUE_SEGMENTS = map[int8]int8 {
	1: 2,
	4: 4,
	7: 3,
	8: 7,
}

var UNIQUE_LENGTHS = [...]int{2, 3, 4, 7}

type Set map[string]bool

func setFromString(str string) Set {
	set := make(Set)
	for _, chr := range str {
        if _, ok := set[string(chr)]; !ok {
			set[string(chr)] = true
		}
	}
	return set
}

func (s Set) isSubsetOf(parentSet Set) bool {
	subSetCheckBool := make([]bool, 0, len(s)) 
	for sKey:= range s {
		for key := range parentSet {
			if (sKey == key) {
				subSetCheckBool = append(subSetCheckBool, true)
				break
			}
		}
	}
	if (len(subSetCheckBool) == len(s)) { return true }
	return false
}

func (s Set) isEqual(anotherSet Set) bool {
	if len(s) != len(anotherSet) { return false }
	isEqual := true
	for sKey := range s {
		hasKey := false
		for key := range anotherSet {
			if (sKey == key) {
				hasKey = true
				break
			}
		}
		if !hasKey {
			isEqual = false
			break
		}
	}
	return isEqual
}

func getInput() Input {
	var input Input
	rawData, _ := os.ReadFile(file_paths.Day_8)
	linesOfRawData := strings.Split(string(rawData), "\r\n")
	for _, line := range linesOfRawData {
		// fmt.Println(line)
		patternAndOutputSlice := strings.Split(line, " | ")
		var pAo PatternAndOutput
		pAo.pattern = strings.Split(patternAndOutputSlice[0], " ")
		pAo.output = strings.Split(patternAndOutputSlice[1], " ")
		input = append(input, pAo)
	}
	return input
}

func part1(input Input) {
	count := 0
	for _, patternAndOutput := range input {
		output := patternAndOutput.output
		for _, pattern := range output {
			for _, uniqLen := range UNIQUE_LENGTHS {
				if len(pattern) == uniqLen { count++ }
			}
		}
	}
	fmt.Println("Part 1: ", count)
}

func part2(input Input) {
	numbersAndSegments := make([]map[int]Set, len(input))
	for i := 0; i < len(input); i++ {
		numbersAndSegments[i] = make(map[int]Set)
	}
	for i, patternAndOutput := range input {
		segmentPattern := patternAndOutput.pattern

		// create a set of numbers of known segments
		for _, pattern := range segmentPattern {
			switch (len(pattern)) {
			case 2:
				numbersAndSegments[i][1] = setFromString(pattern)
			case 3:
				numbersAndSegments[i][7] = setFromString(pattern)
			case 4:
				numbersAndSegments[i][4] = setFromString(pattern)
			case 7:
				numbersAndSegments[i][8] = setFromString(pattern)
			}
		}

		// check for 0, 6, 9
		// if 4 -> subset of number then number = 9
		// if 1 -> subset of number then number = 0
		// else number = 6
		for _, pattern := range segmentPattern {
			patternLength := len(pattern)
			if (patternLength != 6) { continue }
			patternSet := setFromString(pattern)
			if numbersAndSegments[i][4].isSubsetOf(patternSet) {
				numbersAndSegments[i][9] = patternSet
			} else if numbersAndSegments[i][1].isSubsetOf(patternSet) {
				numbersAndSegments[i][0] = patternSet
			} else {
				numbersAndSegments[i][6] = patternSet
			}
		}

		for _, pattern := range segmentPattern {
			patternLength := len(pattern)
			if (patternLength != 5) { continue }
			patternSet := setFromString(pattern)
			if patternSet.isSubsetOf(numbersAndSegments[i][6]) {
				numbersAndSegments[i][5] = patternSet
			} else if patternSet.isSubsetOf(numbersAndSegments[i][9]) {
				numbersAndSegments[i][3] = patternSet
			} else {
				numbersAndSegments[i][2] = patternSet
			}
		}
	}

	total := 0

	for i, patternAndOutput := range input {
		output := patternAndOutput.output
		number := ""
		for _, pattern := range output {
			patternSet := setFromString(pattern)
			for j := 0; j < 10; j++ {
				if numbersAndSegments[i][j].isEqual(patternSet) {
					number += strconv.Itoa(j)
				}
			}
		}
		outputNumber, _ := strconv.Atoi(number)
		total += outputNumber
	}
	fmt.Println("Part 2: ", total)
}

func main() {
	input := getInput()
	part1(input)
	part2(input)
}