package solver

import (
	"push_swap/internal/stack"
	"slices"
	"testing"
)

// rejoue les gestes sur une copie des nombres, comme le ferait le checker
func rejouer(t *testing.T, nums []int, moves []string) stack.Stack {
	s := stack.Stack{A: slices.Clone(nums)}
	for _, m := range moves {
		switch m {
		case "sa":
			s.Sa()
		case "ra":
			s.Ra()
		case "rra":
			s.Rra()
		case "pb":
			s.Pb()
		case "pa":
			s.Pa()
		default:
			t.Fatalf("geste inattendu : %q", m)
		}
	}
	return s
}

func TestSolveTrie(t *testing.T) {
	cas := [][]int{
		{1},
		{2, 1},
		{3, 2, 1},
		{2, 1, 3, 6, 5, 8},
		{4, 67, 3, 87, 23},
		{48213, 48209, 48215, 48210, 48212, 48211},
		{-5, 0, 10, -20, 7},
	}
	for _, nums := range cas {
		s := stack.Stack{A: slices.Clone(nums)}
		moves := Solve(&s)
		r := rejouer(t, nums, moves)
		if !r.IsSorted() {
			t.Errorf("%v : pas trié après %v", nums, moves)
		}
	}
}

func TestSolveDejaTrie(t *testing.T) {
	s := stack.Stack{A: []int{1, 2, 3}}
	if moves := Solve(&s); len(moves) != 0 {
		t.Errorf("déjà trié : attendu 0 geste, obtenu %v", moves)
	}
}
