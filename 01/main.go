package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func PartOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left, right []int

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		values := strings.Split(string(line), "   ")

		value, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, value)

		value, err = strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, value)
	}

	slices.Sort(left)
	slices.Sort(right)

	var result int
	for index := range left {
		result += int(math.Abs(float64(left[index]) - float64(right[index])))
	}

	return result
}

func PartTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	right := make(map[int]int)

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		values := strings.Split(string(line), "   ")

		value, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, value)

		value, err = strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}
		right[value] += 1
	}

	var result int
	for index := range left {
		result += left[index] * right[left[index]]
	}

	return result
}

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}
