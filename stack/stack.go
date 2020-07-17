package stack

import (
	"errors"
)

type Stack struct {
	Data []int
	Size int
}

func (s *Stack) Init(length int) {
	s.Data = make([]int, length)
}

func (s *Stack) Push(n int) error {
	if s.Size == len(s.Data) {
		return errors.New("out of range")
	}
	s.Data[s.Size] = n
	s.Size++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Size == 0 {
		return 0, errors.New("out of range")
	}
	tmp := s.Data[s.Size - 1]
	s.Data[s.Size - 1] = 0
	s.Size--
	return tmp, nil
}

func (s Stack) Top() int {
	return s.Data[s.Size - 1]
}

func (s Stack) IsEmpty() bool {
	if s.Size == 0 {
		return true
	}
	return false
}