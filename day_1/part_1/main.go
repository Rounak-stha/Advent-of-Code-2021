package main

import (
	"AOC/file_paths"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var count int

	file, err := os.Open(file_paths.Day_1)
	check_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	prevVal, err := strconv.Atoi(scanner.Text()) // convert ascii to integer
	check_err(err)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		check_err(err)
		if (val > prevVal) { count++ }
		prevVal = val
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The Number of times the value increases is: %d", count)
}