package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	Name     string
	Size     int
	Prev     *Directory
	Children []*Directory
}

func (d Directory) add(directory *Directory) *Directory {
	return &Directory{
		Name:     d.Name,
		Size:     d.Size + directory.Size,
		Prev:     d.Prev,
		Children: append(d.Children, directory),
	}
}

func main() {
	commands, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	tree := CreateDirectoryTree(commands, &Directory{
		Name:     "/",
		Size:     0,
		Children: nil,
	})
	fmt.Println(CalculateSum(&tree, Part1, false))
	fmt.Println(CalculateSum(&tree, Part2(tree.Size), true))
}

func Part1(dir *Directory) bool {
	return dir.Size <= 100000 && dir.Children != nil
}

func Part2(rootSize int) func(dir *Directory) bool {
	return func(dir *Directory) bool {
		total := 70000000
		atLeast := 30000000
		minimum := total - rootSize
		result := atLeast - minimum
		return dir.Size >= result && dir.Children != nil
	}
}

func CalculateSum(root *Directory, predicate func(dir *Directory) bool, isPart2 bool) int {
	result := make([]*Directory, 0)
	queue := make([]*Directory, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if predicate(nextUp) {
			result = append(result, nextUp)
		}
		if len(nextUp.Children) > 0 {
			for _, child := range nextUp.Children {
				queue = append(queue, child)
			}
		}
	}

	if isPart2 {
		min := root.Size
		for _, d := range result {
			if d.Size < min {
				min = d.Size
			}
		}
		return min
	} else {
		sum := 0
		for _, d := range result {
			sum += d.Size
		}
		return sum
	}
}

func CreateDirectoryTree(commands []string, directory *Directory) Directory {
	if len(commands) == 0 {
		for i := 0; i < len(directory.Prev.Children); i++ {
			child := directory.Prev.Children[i]
			if child.Name == directory.Name {
				directory.Prev.Children[i] = directory
				directory.Prev.Size = directory.Prev.Size + directory.Size
			}
		}
		return *directory.Prev
	}
	command := commands[0]
	if strings.Contains(command, "$ cd ..") {
		for i := 0; i < len(directory.Prev.Children); i++ {
			child := directory.Prev.Children[i]
			if child.Name == directory.Name {
				directory.Prev.Children[i] = directory
				directory.Prev.Size = directory.Prev.Size + directory.Size
			}
		}
		return CreateDirectoryTree(commands[1:], directory.Prev)
	} else if strings.Contains(command, "$ ls") {
		return CreateDirectoryTree(commands[1:], directory)
	} else if strings.Contains(command, "$ cd") {
		if nameFromCd(command) == "/" {
			return CreateDirectoryTree(commands[1:], directory)
		} else {
			for i := 0; i < len(directory.Children); i++ {
				child := directory.Children[i]
				if child.Name == nameFromCd(command) {
					child.Prev = directory
					return CreateDirectoryTree(commands[1:], child)
				}
			}
			return *directory
		}
	} else if strings.Contains(command, "dir") {
		newDir := Directory{
			Name:     strings.Replace(command, "dir ", "", 1),
			Size:     0,
			Prev:     nil,
			Children: nil,
		}
		return CreateDirectoryTree(commands[1:], directory.add(&newDir))
	} else {
		file := strings.Split(command, " ")
		size, _ := strconv.Atoi(file[0])
		newFile := Directory{
			Name:     file[1],
			Size:     size,
			Prev:     nil,
			Children: nil,
		}

		return CreateDirectoryTree(commands[1:], directory.add(&newFile))
	}
}

func nameFromCd(command string) string {
	return strings.Replace(command, "$ cd ", "", 1)
}

func readInput() ([]string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands, nil
}
