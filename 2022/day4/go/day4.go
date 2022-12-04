package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func (aRange Range) isEqual() bool {
	return aRange.Min == aRange.Max
}

type Pair struct {
	Left  Range
	Right Range
}

func (pair Pair) isEqual() bool {
	if pair.Left.isEqual() || pair.Right.isEqual() {
		return pair.Left.Min >= pair.Right.Min && pair.Left.Min <= pair.Right.Max
	}

	return false
}

func (pair Pair) isSubset() bool {
	return pair.isEqual() || (pair.Left.Min <= pair.Right.Min && pair.Left.Max >= pair.Right.Max)
}

func (pair Pair) overlap() bool {
	return pair.isEqual() || (pair.Left.Min <= pair.Right.Min || pair.Left.Min <= pair.Right.Max) && (pair.Left.Max <= pair.Right.Max && pair.Left.Max >= pair.Right.Min)
}

func (pair Pair) overlaps() bool {
	return pair.isSubset() || pair.flip().isSubset() || pair.overlap() || pair.flip().overlap()
}

func (pair Pair) flip() Pair {
	return Pair{
		Left:  pair.Right,
		Right: pair.Left,
	}
}

func NewPair(lmin, lmax, rmin, rmax int) Pair {
	return Pair{
		Left: Range{
			Min: lmin,
			Max: lmax,
		},
		Right: Range{
			Min: rmin,
			Max: rmax,
		},
	}
}

func main() {
	pairs, err := readInput()
	if err != nil {
		log.Fatal(err.Error())
	}

	sum := 0
	for _, pair := range pairs {
		if pair.isSubset() || pair.flip().isSubset() {
			sum = sum + 1
		}
	}

	fmt.Println(sum)

	sum2 := 0
	for _, pair := range pairs {
		if pair.overlaps() {
			sum2 = sum2 + 1
		}
	}

	fmt.Println(sum2)
}

func readInput() ([]Pair, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pairs []Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lrPair := strings.Split(scanner.Text(), ",")
		l := strings.Split(lrPair[0], "-")
		r := strings.Split(lrPair[1], "-")

		lmin, _ := strconv.Atoi(l[0])
		lmax, _ := strconv.Atoi(l[1])
		rmin, _ := strconv.Atoi(r[0])
		rmax, _ := strconv.Atoi(r[1])
		pairs = append(pairs, NewPair(lmin, lmax, rmin, rmax))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pairs, nil
}
