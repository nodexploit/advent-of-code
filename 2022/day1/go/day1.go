package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	calories, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	var elfSumCalories []int
	var elfSum int
	for _, calorie := range calories {
		if calorie == -1 {
			elfSumCalories = append(elfSumCalories, elfSum)
			elfSum = 0
		} else {
			elfSum = elfSum + calorie
		}
	}

	var max int
	for _, sum := range elfSumCalories {
		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)

	sort.Sort(sort.Reverse(sort.IntSlice(elfSumCalories)))

	fmt.Println(elfSumCalories[0] + elfSumCalories[1] + elfSumCalories[2])
}

// readInput loads the input file with -1 for whitespaces in a slice
func readInput() ([]int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var calories []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calorie, err := strconv.Atoi(scanner.Text())
		if err != nil {
			calories = append(calories, -1)
		} else {
			calories = append(calories, calorie)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return calories, nil
}
