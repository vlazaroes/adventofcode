package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(values []string) bool {
	var increasing, decreasing = false, false

	for index := 0; index < len(values)-1; index++ {
		current, err := strconv.Atoi(values[index])
		if err != nil {
			log.Fatal(err)
		}

		next, err := strconv.Atoi(values[index+1])
		if err != nil {
			log.Fatal(err)
		}

		difference := current - next
		if difference == 0 {
			return false
		}

		if difference > 0 {
			increasing = true
		}
		if difference < 0 {
			decreasing = true
		}

		if increasing == true && decreasing == true {
			return false
		}

		if difference < -3 || difference > 3 {
			return false
		}
	}

	return true
}

func PartOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		values := strings.Split(string(line), " ")

		if isSafe(values) {
			result += 1
		}
	}

	return result
}

func checkDampener(values []string) bool {
	for index := range values {
		valuesCopy := make([]string, len(values))
		copy(valuesCopy, values)

		valuesCopy = append(valuesCopy[:index], valuesCopy[index+1:]...)

		if isSafe(valuesCopy) {
			return true
		}
	}

	return false
}

func PartTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		values := strings.Split(string(line), " ")

		if isSafe(values) || checkDampener(values) {
			result += 1
		}
	}

	return result
}

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}
