package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Shape uint8

const (
	ElfRock     Shape = 'A'
	ElfPaper    Shape = 'B'
	ElfScissors Shape = 'C'
	Rock        Shape = 'X'
	Paper       Shape = 'Y'
	Scissors    Shape = 'Z'
)

type SelectedShape struct {
	Shape  Shape
	Points int
}

func NewSelectedShape(shape Shape) SelectedShape {
	var points int

	switch shape {
	case Rock:
		points = 1
	case Paper:
		points = 2
	case Scissors:
		points = 3
	}

	return SelectedShape{
		Shape:  shape,
		Points: points,
	}
}

func SameShape(shape Shape) Shape {
	switch shape {
	case ElfRock, Rock:
		return Rock
	case ElfPaper, Paper:
		return Paper
	case ElfScissors, Scissors:
		return Scissors
	}
	return 0
}

const (
	Loss = 0
	Draw = 3
	Win  = 6
)

type Game struct {
	ElfShape      Shape
	SelectedShape SelectedShape
}

func main() {
	var winMoves = make(map[Shape]Shape)
	winMoves[Rock] = ElfScissors
	winMoves[Scissors] = ElfPaper
	winMoves[Paper] = ElfRock

	games, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	var gameResult int
	for _, game := range games {
		gameResult = gameResult + Play(winMoves, game.ElfShape, game.SelectedShape)
	}

	fmt.Println(gameResult)

	var gameResult2 int
	for _, game2 := range games {
		gameResult2 = gameResult2 + Play2(winMoves, game2.ElfShape, game2.SelectedShape)
	}

	fmt.Println(gameResult2)
}

func Play(winMoves map[Shape]Shape, elfShape Shape, selectedShape SelectedShape) int {
	selectedMove := winMoves[selectedShape.Shape]
	if elfShape == selectedMove {
		return selectedShape.Points + Win
	} else if SameShape(elfShape) == SameShape(selectedShape.Shape) {
		return selectedShape.Points + Draw
	} else {
		return selectedShape.Points + Loss
	}
}

func Play2(winMoves map[Shape]Shape, elfShape Shape, selectedShape SelectedShape) int {
	if selectedShape.Shape == Paper {
		return Play(winMoves, elfShape, NewSelectedShape(SameShape(elfShape)))
	} else if selectedShape.Shape == Rock {
		var loseElfMoves = make(map[Shape]Shape)
		loseElfMoves[ElfRock] = Scissors
		loseElfMoves[ElfScissors] = Paper
		loseElfMoves[ElfPaper] = Rock

		return Play(winMoves, elfShape, NewSelectedShape(loseElfMoves[elfShape]))
	} else {
		var winElfMoves = make(map[Shape]Shape)
		winElfMoves[ElfRock] = Paper
		winElfMoves[ElfScissors] = Rock
		winElfMoves[ElfPaper] = Scissors

		return Play(winMoves, elfShape, NewSelectedShape(winElfMoves[elfShape]))
	}
}

func readInput() ([]Game, error) {
	file, err := os.Open("../input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var games []Game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var elfShape Shape = Shape(scanner.Text()[0])
		var selectedShape Shape = Shape(scanner.Text()[2])

		games = append(games, Game{
			ElfShape:      elfShape,
			SelectedShape: NewSelectedShape(selectedShape),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return games, nil
}
