package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CalculatePriority(character int) int {
	a := 1
	A := 27
	if character >= 'A' && character <= 'Z' {
		return (character - 'A') + A
	} else {
		return (character - 'a') + a
	}
}

// SanitizeCompartment removes duplicates using a map as a set
func SanitizeCompartment(compartment string) map[int]int {
	chars := make(map[int]int)
	for _, char := range compartment {
		chars[int(char)] = 1
	}

	return chars
}

func CountItems(sanitizedCompartment map[int]int, result map[int]int) {
	for k := range sanitizedCompartment {
		if count, ok := result[k]; ok {
			result[k] = count + 1
		} else {
			result[k] = 1
		}
	}
}

func GetDuplicatedItem(rucksack string) int {
	mid := len(rucksack) / 2
	firstCompartment := rucksack[:mid]
	secondCompartment := rucksack[mid:]

	firstSanitizedCompartment := SanitizeCompartment(firstCompartment)
	secondSanitizedCompartment := SanitizeCompartment(secondCompartment)

	chars := make(map[int]int)
	CountItems(firstSanitizedCompartment, chars)
	CountItems(secondSanitizedCompartment, chars)

	for k, v := range chars {
		if v > 1 {
			return k
		}
	}

	return 0
}

func CalculateBadgePriority(rucksacks []string) int {
	sum := 0
	for i := range rucksacks {
		if (i+1)%3 == 0 {
			first := SanitizeCompartment(rucksacks[i])
			second := SanitizeCompartment(rucksacks[i-1])
			third := SanitizeCompartment(rucksacks[i-2])

			chars := make(map[int]int)
			CountItems(first, chars)
			CountItems(second, chars)
			CountItems(third, chars)

			for k, v := range chars {
				if v > 2 {
					sum = sum + CalculatePriority(k)
				}
			}
		}
	}
	return sum
}

func main() {
	rucksacks, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, rucksack := range rucksacks {
		sum = sum + CalculatePriority(GetDuplicatedItem(rucksack))
	}

	fmt.Println(sum)
	fmt.Println(CalculateBadgePriority(rucksacks))
}

func readInput() ([]string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rucksacks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rucksacks, nil
}
