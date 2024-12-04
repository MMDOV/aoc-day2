package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func increasing(numbers []int) bool {
	for i, num := range numbers {
		if i != 0 {
			prev_num := numbers[i-1]
			num_distance := num - prev_num
			if !(num > prev_num && num_distance >= 1 && num_distance <= 3) {
				return false
			}
		}
	}
	return true
}

func deceresing(numbers []int) bool {
	for i, num := range numbers {
		if i != 0 {
			prev_num := numbers[i-1]
			num_distance := prev_num - num
			if !(num < prev_num && num_distance >= 1 && num_distance <= 3) {
				return false
			}
		}
	}
	return true
}

func checkBoth(arr []int) bool {
	if increasing(arr) {
		return true
	} else if deceresing(arr) {
		return true
	} else {
		return false
	}
}

func removeFromArray(arr []int, indexToRemove int) []int {
	var newNumbers []int
	for index, num := range arr {
		if index != indexToRemove {
			newNumbers = append(newNumbers, num)
		}
	}
	return newNumbers
}

func readNumbers(fileName string) [][]int {
	var allNumbers [][]int
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return allNumbers
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strNumbers := strings.Fields(scanner.Text())
		var intNumbers []int
		for _, num := range strNumbers {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
				return allNumbers
			}
			intNumbers = append(intNumbers, numInt)
		}
		allNumbers = append(allNumbers, intNumbers)
	}
	return allNumbers
}

func main() {
	allNumbers := readNumbers("day2.input")
	var safeCount int
	for _, arr := range allNumbers {
		if checkBoth(arr) {
			safeCount += 1
			continue
		} else {
			for i := 0; i < len(arr); i++ {
				newArr := removeFromArray(arr, i)
				if checkBoth(newArr) {
					safeCount += 1
					break
				}
			}
		}
	}
	fmt.Println(safeCount)
}
