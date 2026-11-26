package solver

import (
	"push_swap/internal/stack"
	"slices"
)

func minIndex(list []int) int {
	if len(list) == 0 {
		return -1
	}
	valMin := slices.Min(list)
	posMin := slices.Index(list, valMin)
	return posMin
}

func Solve(s *stack.Stack) []string {
	moves := []string{} //liste vide de mouvements
	for len(s.A) > 0 {

	}
	return moves
}
