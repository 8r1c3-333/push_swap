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

func sortSmall(s *stack.Stack, moves []string) []string { //trie A quand il reste 2 ou 3 nombres, en 2 gestes max.
	if len(s.A) == 2 {
		if s.A[0] > s.A[1] {
			s.Sa()
			moves = append(moves, "sa")
		}
		return moves
	}
	if len(s.A) != 3 {
		return moves
	}
	a, b, c := s.A[0], s.A[1], s.A[2]
	switch {
	case a > b && b < c && a < c: // 2 1 3
		s.Sa()
		moves = append(moves, "sa")
	case a > b && b < c && a > c: // 3 1 2
		s.Ra()
		moves = append(moves, "ra")
	case a < b && b > c && a > c: // 2 3 1
		s.Rra()
		moves = append(moves, "rra")
	case a < b && b > c && a < c: // 1 3 2
		s.Sa()
		s.Ra()
		moves = append(moves, "sa", "ra")
	case a > b && b > c: // 3 2 1
		s.Sa()
		s.Rra()
		moves = append(moves, "sa", "rra")
	}
	return moves
}

func Solve(s *stack.Stack) []string {
	moves := []string{} // liste vide de mouvements
	if s.IsSorted() {
		return moves
	}
	for len(s.A) > 3 {
		if minIndex(s.A) <= len(s.A)/2 {
			for minIndex(s.A) != 0 {
				s.Ra()
				moves = append(moves, "ra")
			}
		} else {
			for minIndex(s.A) != 0 {
				s.Rra()
				moves = append(moves, "rra")
			}
		}
		s.Pb()
		moves = append(moves, "pb")
	}
	moves = sortSmall(s, moves)
	for len(s.B) > 0 {
		s.Pa()
		moves = append(moves, "pa")
	}
	return moves
}
