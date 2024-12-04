package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int

	re := regexp.MustCompile(`mul\(([0-9]{1,3},[0-9]{1,3})\)`)

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		values := re.FindAllSubmatch(line, -1)

		for _, value := range values {
			args := strings.Split(string(value[1]), ",")

			value1, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}

			value2, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
			}

			result += value1 * value2
		}
	}

	return result
}

func PartTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int
	isEnabled := true

	re := regexp.MustCompile(`mul\(([0-9]{1,3},[0-9]{1,3})\)|don't\(\)|do\(\)`)

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		values := re.FindAllSubmatch(line, -1)

		for _, value := range values {
			if string(value[0]) == "do()" {
				isEnabled = true
				continue
			}
			if string(value[0]) == "don't()" {
				isEnabled = false
				continue
			}

			if !isEnabled {
				continue
			}

			args := strings.Split(string(value[1]), ",")

			value1, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}

			value2, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
			}

			result += value1 * value2
		}
	}

	return result
}

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}
