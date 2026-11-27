package parser

import (
	"slices"
	"testing"
)

func TestParseValide(t *testing.T) {
	got, err := Parse([]string{"3 1", "2"})
	if err != nil {
		t.Fatalf("erreur inattendue : %v", err)
	}
	want := []int{3, 1, 2}
	if !slices.Equal(got, want) {
		t.Errorf("Parse = %v, attendu %v", got, want)
	}
}

func TestParseErreurs(t *testing.T) {
	cas := [][]string{
		{"0 one 2"}, // pas un nombre
		{"5 5"},     // doublon dans un argument
		{"1", "1"},  // doublon entre 2 arguments
	}
	for _, args := range cas {
		if _, err := Parse(args); err == nil {
			t.Errorf("Parse(%v) : erreur attendue", args)
		}
	}
}

func TestParseVide(t *testing.T) {
	got, err := Parse([]string{})
	if err != nil || len(got) != 0 {
		t.Errorf("Parse vide = %v, %v ; attendu liste vide, nil", got, err)
	}
}
