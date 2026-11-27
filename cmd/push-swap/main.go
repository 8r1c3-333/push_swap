package main

import (
	"fmt"
	"os"
	"push_swap/internal/parser"
	"push_swap/internal/solver"
	"push_swap/internal/stack"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	nb, err := parser.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	stacks := stack.Stack{A: nb}
	moves := solver.Solve(&stacks)
	for _, move := range moves {
		fmt.Println(move)
	}
}
