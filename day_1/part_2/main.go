package main

import (
	"AOC/file_paths"
	"AOC/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile(file_paths.Day_2)
	utils.HandleError(err)
	content := string(contents)
	contentSlice := strings.Split(content, "\r\n")
	numbers := sliceStoI(contentSlice)
	fmt.Println(slidingWindowIncrementNumber(numbers))
}

func sliceStoI(sSlice []string) []int {
	var numbers []int
	for _, num := range sSlice[:len(sSlice) -1] {
		number, err := strconv.Atoi(num)
		utils.HandleError(err)
		numbers = append(numbers, number)
	}
	return numbers
}

func slidingWindowIncrementNumber(numbers []int) int {
	prev := numbers[0] + numbers[1] + numbers[2]
	var res int
	for i := 1; i <= len(numbers) - 3; i++ {
		curr := numbers[i] + numbers[i+1] + numbers[i+2]
		if curr > prev {
			res++
		}
		prev = curr
	}
	return res
}