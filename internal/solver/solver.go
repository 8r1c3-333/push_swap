package solver

import (
	"push_swap/internal/stack"
	"slices"
)

func minIndex(list []int) int { //trouver le minimum de la liste
	if len(list) == 0 {
		return -1
	}
	valMin := slices.Min(list)
	posMin := slices.Index(list, valMin)
	return posMin
}

func Solve(s *stack.Stack) []string {
	moves := []string{} //liste vide de mouvements
	if s.IsSorted() {
		return moves
	}
	for len(s.A) > 0 {
		for minIndex(s.A) != 0 {
			s.Ra()
			moves = append(moves, "ra")
		}
		s.Pb()
		moves = append(moves, "pb")
	}
	for len(s.B) > 0 {
		s.Pa()
		moves = append(moves, "pa")
	}
	return moves
}
