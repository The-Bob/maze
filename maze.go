package main

import (
	"fmt"
	//"github.com/gdamore/tcell"
	"github.com/golang-collections/collections/stack"
	//"os"
)

type maze struct {
	mazeWidth, mazeHeight, nVisitedCells int
	cells                                []int
	stack                                stack.Stack

	cellN       int
	cellE       int
	cellS       int
	cellW       int
	cellVisited int
}

func newMaze(mazeWidth, mazeHeight, nVisitedCells int, cells []int, stack *stack.Stack) maze {
	var m maze
	return m
}

func main() {
	/*	screen, err := tcell.NewScreen()
		if err != nil {
			fmt.Errorf("%v", err)
			os.Exit(1)
		}
		err = screen.Init()
		if err != nil {
			fmt.Errorf("%v", err)
			os.Exit(1)
		}

		screen.SetStyle(tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorBlack))
		screen.S
		fmt.Println(screen.Size())
		screen.SetContent(3, 3, 'q', nil, tcell.StyleDefault)
		screen.Sync()
	*/
	var s = stack.New()
	s.Push("potato")
	fmt.Println(s)

	var maze maze
	maze = newMaze(1, 1, 1, []int{1, 2, 3}, s)
	fmt.Println(maze)
}
