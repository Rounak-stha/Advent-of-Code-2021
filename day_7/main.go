package main

import (
	"AOC/file_paths"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() []int {
	var parsedData []int
	rawData, _ := os.ReadFile(file_paths.Day_7)
	data := strings.Split(string(rawData), ",")
	for _, strNum := range data {
		num, _ := strconv.Atoi(strNum)
		parsedData = append(parsedData, num)
	}
	return parsedData
}

func calculateMedian(data []int) (int, []int, map[int]int) {
	sort.Ints(data)
	var median int
	frequency := make(map[int]int)
	cFrequency := make(map[int]int)
	uniquePositions := []int{data[0]}
	prev := data[0]

	frequency[prev] = 1

	var medianAt int
	if len(data) % 2 == 0 {
		medianAt = (len(data) / 2) + 1
	} else {
		medianAt = (len(data) + 1) / 2
	}

	for _, num := range data[1:] {
		if num == prev {
			frequency[num] += 1
		} else {
			frequency[num] = 1
			uniquePositions = append(uniquePositions, num)
			prev = num
		}
	}
	cFrequency[uniquePositions[0]] = frequency[uniquePositions[0]]
	if cFrequency[uniquePositions[0]] >= medianAt {
		median = uniquePositions[0]
	} else {
		for i, key:= range uniquePositions[1:] {
			cFrequency[key] = cFrequency[uniquePositions[i]] + frequency[key]
			if cFrequency[key] >= medianAt {
				median = key
				break
			}
		}
	}
	
	return median, uniquePositions, frequency
}

func sumOfNaturalNumbers(n int) int {
	return (n * (n + 1 )) / 2
}

// since fuel cost increases by 1 unit each step
// so if the number of steps is 5, total fuel cost: 1 + 2 + 3 + 4 + 5 = 15
// which is the sum of first n natural numbers
// in this case n is 5
func part2(uniquePositions []int, frequency map[int]int) int {
	var lowestFuelCost int
	lastPosition := uniquePositions[len(uniquePositions) - 1]

	// for each position calculate the number of steps crabs at each position have to step and find
	// the sum of first ( n = steps ) natural numbers for all the position and crabs at each position
	
	// initial cost 
	for _, num := range uniquePositions[1:] {
		lowestFuelCost += frequency[num] * sumOfNaturalNumbers(int(math.Abs(float64(uniquePositions[0]) - float64(num))))
	}

	for position := 1; position <= lastPosition ; position++ {
		fuelCost := 0
		for _, num := range uniquePositions {
			if position == num {
				continue
			} else {
				f := frequency[num]
				nSum := sumOfNaturalNumbers(int(math.Abs(float64(position) - float64(num))))
				fuelCost += f * nSum
			}
		}
		if fuelCost < lowestFuelCost {
			lowestFuelCost = fuelCost
		}
	}
	return lowestFuelCost

}

// Find out where the maxinum number of crabs lie (median of the positions (median of frequency distribution))
// so that fewer crabs have to move
// hence low fuel cost
func part1(median int, data []int) int {
	fuelCost := 0.
	for _, pos := range data {
		fuelCost += math.Abs(float64(pos - median))
	}
	return int(fuelCost)
}

func main() {
	data := getInput()
	median, uniquePositions, frequency := calculateMedian(data)
	fuelCost := part1(median, data)
	fmt.Println("Part 1 Lowest Fuel Cost: ", int(fuelCost))
	part2FuelCost := part2(uniquePositions, frequency)
	fmt.Println("Part 1 Lowest Fuel Cost: ",part2FuelCost)
}