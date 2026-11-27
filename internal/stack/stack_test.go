package stack

import (
	"slices"
	"testing"
)

func TestGestes(t *testing.T) {
	s := Stack{A: []int{5, 8, 2}}

	s.Sa()
	if !slices.Equal(s.A, []int{8, 5, 2}) {
		t.Errorf("Sa : A = %v", s.A)
	}
	s.Ra()
	if !slices.Equal(s.A, []int{5, 2, 8}) {
		t.Errorf("Ra : A = %v", s.A)
	}
	s.Rra()
	if !slices.Equal(s.A, []int{8, 5, 2}) {
		t.Errorf("Rra : A = %v", s.A)
	}
	s.Pb()
	if !slices.Equal(s.A, []int{5, 2}) || !slices.Equal(s.B, []int{8}) {
		t.Errorf("Pb : A = %v, B = %v", s.A, s.B)
	}
	s.Pa()
	if !slices.Equal(s.A, []int{8, 5, 2}) || len(s.B) != 0 {
		t.Errorf("Pa : A = %v, B = %v", s.A, s.B)
	}
}

func TestPileTropPetite(t *testing.T) {
	s := Stack{A: []int{1}}
	s.Sa() // ne doit rien faire
	s.Pa() // B vide : ne doit rien faire
	if !slices.Equal(s.A, []int{1}) {
		t.Errorf("A = %v, attendu [1]", s.A)
	}
}

func TestIsSorted(t *testing.T) {
	if !(&Stack{A: []int{1, 2, 3}}).IsSorted() {
		t.Error("[1 2 3] devrait être trié")
	}
	if (&Stack{A: []int{2, 1, 3}}).IsSorted() {
		t.Error("[2 1 3] ne devrait pas être trié")
	}
	if (&Stack{A: []int{1, 2}, B: []int{3}}).IsSorted() {
		t.Error("B non vide : ne devrait pas être trié")
	}
}
