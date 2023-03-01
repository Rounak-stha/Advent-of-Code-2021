package main

import (
	"AOC/file_paths"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() []int {
	var parsedData []int
	rawData, _ := os.ReadFile(file_paths.Day_6)
	data := strings.Split(string(rawData), ",")
	for _, strNum := range data {
		num, _ := strconv.Atoi(strNum)
		parsedData = append(parsedData, num)
	}
	return parsedData
}

func initialPopulationTimers(data []int) map[int]int {
	fishWithSameTimer := make(map[int]int)
	for _, timer := range data {
		if _, ok := fishWithSameTimer[timer]; ok {
			fishWithSameTimer[timer] += 1
		} else {
			fishWithSameTimer[timer] = 1
		}
	}
	return fishWithSameTimer
}

func finalPopulationTimers(initialTimers map[int]int, days int) map[int]int {
	currDay := 1
	fishWithSameTimer := initialTimers
	for currDay <= days {
		tempFishTimerData := make(map[int]int)
		for timer, numOfFishes := range fishWithSameTimer {
			if timer == 0 {
				tempFishTimerData[8] = numOfFishes
				tempFishTimerData[6] += numOfFishes
			}else{ 
				tempFishTimerData[timer - 1] += numOfFishes
			}
		} 
		fishWithSameTimer = tempFishTimerData
		currDay++
	}
	return fishWithSameTimer
}

func main() {
	var part1_total int
	var part2_total int

	data := getInput()

	// Timer Structiure map { ...timer: num_of_fishes... }
	fishWithSameTimer := initialPopulationTimers(data)
	part1_finalTimers := finalPopulationTimers(fishWithSameTimer, 80)
	part2_finalTimers := finalPopulationTimers(fishWithSameTimer, 256)

	for _, num := range part1_finalTimers {
		part1_total += num
	}

	for _, num := range part2_finalTimers {
		part2_total += num
	}
	fmt.Println("part_1: ", part1_total)
	fmt.Println("part_2: ", part2_total)
}