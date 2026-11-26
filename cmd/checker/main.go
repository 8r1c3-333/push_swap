package main

import (
	"bufio"
	"fmt"
	"os"
	"push_swap/internal/parser"
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ligne := scanner.Text()
		switch ligne {
		case "sa":
			stacks.Sa()
		case "sb":
			stacks.Sb()
		case "ss":
			stacks.Ss()
		case "pa":
			stacks.Pa()
		case "pb":
			stacks.Pb()
		case "ra":
			stacks.Ra()
		case "rb":
			stacks.Rb()
		case "rr":
			stacks.Rr()
		case "rra":
			stacks.Rra()
		case "rrb":
			stacks.Rrb()
		case "rrr":
			stacks.Rrr()
		default:
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	if stacks.IsSorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
