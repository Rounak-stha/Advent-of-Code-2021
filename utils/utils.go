package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func GetIntParsedInput(path string) []int {
	allContent, err := os.ReadFile(path)
	HandleError(err)
	contentSlice := strings.Split(string(allContent), "\r\n")
	var numbers []int
	for _, chr := range contentSlice {
		numb, err := strconv.Atoi(chr)
		HandleError(err)
		numbers = append(numbers, numb)
	}
	return numbers
}

func GetInput(path string) []string {
	allContent, err := os.ReadFile(path)
	HandleError(err)
	contentSlice := strings.Split(string(allContent), "\r\n")
	return contentSlice
}


