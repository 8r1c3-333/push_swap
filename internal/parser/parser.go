package parser

import (
	"errors"
	"strconv"
	"strings"
)

func Parse(args []string) ([]int, error) {
	nbVus := make(map[int]bool) // pour garder la trace des nombres déjà vus;
	nbValides := []int{}        // pour stocker les nombres valides
	for _, arg := range args {
		s := strings.Fields(arg) // pour séparer les nombres dans l'argument;
		for _, n := range s {
			nb, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			if nbVus[nb] {
				return nil, errors.New("duplicate number found")
			}
			nbValides = append(nbValides, nb)
			nbVus[nb] = true
		}
	}
	return nbValides, nil
}
