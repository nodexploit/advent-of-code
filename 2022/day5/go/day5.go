package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1] // by default slices has [low : high] indexes, if you omit them it uses 0 for low and len(s) for high
}

type movement struct {
	Quantity int
	From     int
	To       int
}

func main() {
	stack1 := make(stack, 0)
	stack1 = stack1.Push('F')
	stack1 = stack1.Push('C')
	stack1 = stack1.Push('J')
	stack1 = stack1.Push('P')
	stack1 = stack1.Push('H')
	stack1 = stack1.Push('T')
	stack1 = stack1.Push('W')

	stack2 := make(stack, 0)
	stack2 = stack2.Push('G')
	stack2 = stack2.Push('R')
	stack2 = stack2.Push('V')
	stack2 = stack2.Push('F')
	stack2 = stack2.Push('Z')
	stack2 = stack2.Push('J')
	stack2 = stack2.Push('B')
	stack2 = stack2.Push('H')

	stack3 := make(stack, 0)
	stack3 = stack3.Push('H')
	stack3 = stack3.Push('P')
	stack3 = stack3.Push('T')
	stack3 = stack3.Push('R')

	stack4 := make(stack, 0)
	stack4 = stack4.Push('Z')
	stack4 = stack4.Push('S')
	stack4 = stack4.Push('N')
	stack4 = stack4.Push('P')
	stack4 = stack4.Push('H')
	stack4 = stack4.Push('T')

	stack5 := make(stack, 0)
	stack5 = stack5.Push('N')
	stack5 = stack5.Push('V')
	stack5 = stack5.Push('F')
	stack5 = stack5.Push('Z')
	stack5 = stack5.Push('H')
	stack5 = stack5.Push('J')
	stack5 = stack5.Push('C')
	stack5 = stack5.Push('D')

	stack6 := make(stack, 0)
	stack6 = stack6.Push('P')
	stack6 = stack6.Push('M')
	stack6 = stack6.Push('G')
	stack6 = stack6.Push('F')
	stack6 = stack6.Push('W')
	stack6 = stack6.Push('D')
	stack6 = stack6.Push('Z')

	stack7 := make(stack, 0)
	stack7 = stack7.Push('M')
	stack7 = stack7.Push('V')
	stack7 = stack7.Push('Z')
	stack7 = stack7.Push('W')
	stack7 = stack7.Push('S')
	stack7 = stack7.Push('J')
	stack7 = stack7.Push('D')
	stack7 = stack7.Push('P')

	stack8 := make(stack, 0)
	stack8 = stack8.Push('N')
	stack8 = stack8.Push('D')
	stack8 = stack8.Push('S')

	stack9 := make(stack, 0)
	stack9 = stack9.Push('D')
	stack9 = stack9.Push('Z')
	stack9 = stack9.Push('S')
	stack9 = stack9.Push('F')
	stack9 = stack9.Push('M')

	moves, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	stacks := [9]stack{stack1, stack2, stack3, stack4, stack5, stack6, stack7, stack8, stack9}

	for _, m := range moves {
		//makeMove(m, &stacks) // uncomment this for part1 solution
		makeMoveMultiple(m, &stacks)
	} 

	var result string
	for _, s := range stacks {
		_, elem := s.Pop()
		result = result + string(elem)
	}

	fmt.Println(result)

	
}

func makeMove(move movement, stacks *[9]stack) {
	fromIdx := move.From - 1
	toIdx := move.To - 1
	for i := 0; i < move.Quantity; i++ {
		s, elem := stacks[fromIdx].Pop()
		stacks[toIdx] = stacks[toIdx].Push(elem)
		stacks[fromIdx] = s
	}
}

func makeMoveMultiple(move movement, stacks *[9]stack) {
	fromIdx := move.From - 1
	toIdx := move.To - 1

	fromLen := len(stacks[fromIdx])
	fromS := stacks[fromIdx][fromLen-move.Quantity:]
	stacks[toIdx] = append(stacks[toIdx], fromS...)
	stacks[fromIdx] = stacks[fromIdx][:fromLen-move.Quantity]
}

func readInput() ([]movement, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var moves []movement
	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mo := re.FindAllString(scanner.Text(), -1)
		if len(mo) == 3 {
			q, _ := strconv.Atoi(mo[0])
			f, _ := strconv.Atoi(mo[1])
			t, _ := strconv.Atoi(mo[2])
			moves = append(moves, movement{
				Quantity: q,
				From:     f,
				To:       t,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return moves, nil
}
