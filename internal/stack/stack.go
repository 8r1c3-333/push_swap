package stack

type Stack struct {
	A []int
	B []int
}

func (s *Stack) Sa() {
	if len(s.A) < 2 {
		return
	}
	s.A[0], s.A[1] = s.A[1], s.A[0]
}

func (s *Stack) Sb() {
	if len(s.B) < 2 {
		return
	}
	s.B[0], s.B[1] = s.B[1], s.B[0]
}

func (s *Stack) Ss() {
	s.Sa()
	s.Sb()
}

func (s *Stack) Pb() {
	if len(s.A) == 0 {
		return
	}
	hautA := s.A[0]
	s.A = s.A[1:]
	s.B = append([]int{hautA}, s.B...)
}

func (s *Stack) Pa() {
	if len(s.B) == 0 {
		return
	}
	hautB := s.B[0]
	s.B = s.B[1:]
	s.A = append([]int{hautB}, s.A...)
}

func (s *Stack) Ra() {
	if len(s.A) == 0 {
		return
	}
	hautA := s.A[0]
	s.A = s.A[1:]
	s.A = append(s.A, hautA)
}

func (s *Stack) Rb() {
	if len(s.B) == 0 {
		return
	}
	hautB := s.B[0]
	s.B = s.B[1:]
	s.B = append(s.B, hautB)
}

func (s *Stack) Rr() {
	s.Ra()
	s.Rb()
}

func (s *Stack) Rra() {
	if len(s.A) == 0 {
		return
	}
	basA := s.A[len(s.A)-1]
	s.A = s.A[:len(s.A)-1]
	s.A = append([]int{basA}, s.A...)
}

func (s *Stack) Rrb() {
	if len(s.B) == 0 {
		return
	}
	basB := s.B[len(s.B)-1]
	s.B = s.B[:len(s.B)-1]
	s.B = append([]int{basB}, s.B...)
}

func (s *Stack) Rrr() {
	s.Rra()
	s.Rrb()
}

func (s *Stack) IsSorted() bool {
	if len(s.B) > 0 {
		return false
	}
	for i := 0; i < len(s.A)-1; i++ {
		if s.A[i] < s.A[i+1] {
			return false
		}
	}
	return true
}
