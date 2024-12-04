package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findWord(matrix [][]byte, letters []byte, currentX, currentY, deltaX, deltaY int) bool {
	for index := 1; index < len(letters); index++ {
		newX, newY := currentX+index*deltaX, currentY+index*deltaY

		if newX < 0 || newX >= len(matrix[currentY]) || newY < 0 || newY >= len(matrix) || matrix[newY][newX] != letters[index] {
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

	var lineCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Seek(0, 0)

	var result int

	matrix := make([][]byte, lineCount)

	reader := bufio.NewReader(file)
	for index := range matrix {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)

		matrix[index] = lineCopy
	}

	letters := []byte{'X', 'M', 'A', 'S'}
	deltas := []struct{ x, y int }{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] != letters[0] {
				continue
			}

			for _, delta := range deltas {
				if findWord(matrix, letters, x, y, delta.x, delta.y) {
					result++
				}
			}
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

	var lineCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Seek(0, 0)

	var result int

	matrix := make([][]byte, lineCount)

	reader := bufio.NewReader(file)
	for index := range matrix {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)

		matrix[index] = lineCopy
	}

	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] != 'A' {
				continue
			}

			if x <= 0 || x >= len(matrix[y])-1 || y <= 0 || y >= len(matrix[y])-1 {
				continue
			}

			chars := string(matrix[y-1][x-1]) + string(matrix[y-1][x+1]) + string(matrix[y+1][x+1]) + string(matrix[y+1][x-1])

			if chars == "MSSM" || chars == "MMSS" || chars == "SMMS" || chars == "SSMM" {
				result++
			}
		}
	}

	return result
}

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}
