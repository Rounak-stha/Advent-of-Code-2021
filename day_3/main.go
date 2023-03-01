package main

import (
	"AOC/file_paths"
	"AOC/utils"
	"fmt"
	"strconv"
)

type Count struct {
	one int
	zero int
}

func main() {
	input := utils.GetInput(file_paths.Day_3)

	counts := getCountsOfBitsAtEachPosition(input)
	g, e := getGammaAndEpsilon(counts)
	o, c := getOxygenAndCo2(input)
	fmt.Printf("Power Consumption of the Submarine is: %d\n", g * e) 
	fmt.Printf("The Life Support rating of the Submarine is %d", o * c)
}

func getBitCountAtSpecificPosition(input []string, index int) Count {
	var count Count
	for _, numb := range input {
		if string(numb[index]) == "1" {
			count.one += 1
		} else {
			count.zero += 1
		}
	}

	return count
}

func getCountsOfBitsAtEachPosition(input []string) []Count {
	numb := input[0]
	
	counts := make([]Count, len(numb))

	// value if not initialized gets zero value of the value's type
	for i := 0; i < len(input); i++ {
		numb := input[i]
		for i, chr := range numb {
			if chr == '1' {
				counts[i].one += 1
			}
			if chr == '0' {
				counts[i].zero += 1
			}
		}
	}

	return counts
}

func getGammaAndEpsilon(counts []Count) (int64, int64) {
	var g,e string
	for _, value := range counts {
		if value.one > value.zero {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		} 
	}
	gamma, err := strconv.ParseInt(g, 2, 32)
	utils.HandleError(err)
	epsilon, err := strconv.ParseInt(e, 2, 32)
	utils.HandleError(err)
	return gamma, epsilon
}

/* 
	Oxygen: Most Common Bit(MSB) at each position
			If number of both bits (1 and 0) are same, keep number with 1 in that position
	CO2: Least Common Bit(MSB) at each position
			If number of both bits (1 and 0) are same, keep number with 0 in that position

*/
func getOxygenAndCo2(input []string) (int64, int64) {
	var oxygen, CO2, MSB, LSB string
	var index int
	
	numberLength := len(input[0])
	oFiltered, cFiltered := append(make([]string, 0, len(input)), input...), append(make([]string, 0, len(input)), input...)

	for {
		oBitCount := getBitCountAtSpecificPosition(oFiltered, index)
		if oBitCount.one >= oBitCount.zero {
			MSB = "1"
		} else {
			MSB = "0"
		}
		oFiltered = filterByBit(oFiltered, MSB, index)
		if len(oFiltered) == 1 { break }
		index++
		if index > numberLength {
			index = 0
		}
	}

	index = 0
	for {
		cBitCount := getBitCountAtSpecificPosition(cFiltered, index)

		if cBitCount.one < cBitCount.zero {
			LSB = "1"
		} else {
			LSB = "0"
		}
		cFiltered = filterByBit(cFiltered, LSB, index)
		if len(cFiltered) == 1 { break }
		index++
		if index > numberLength {
			index = 0
		}
	}
	oxygen = oFiltered[0]
	CO2 = cFiltered[0]

	O, err := strconv.ParseInt(oxygen, 2, 32)
	utils.HandleError(err) 
	C, err := strconv.ParseInt(CO2, 2, 32)
	utils.HandleError(err)

	return O, C
}

func filterByBit(strings []string, bit string, index int) []string {
	res := make([]string, 0, len(strings) / 2)

	for _, numb := range strings {
		if string(numb[index]) == bit {
			res = append(res, numb)
		}
	}
	return res
}





