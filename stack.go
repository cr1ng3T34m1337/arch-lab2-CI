package lab2

import "errors"

type Stack []float64

func (s *Stack) Push(item float64) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (float64, error) {
	l := len(*s)
	if l == 0 {
		return 0.0, errors.New("Stack is empty")
	}
	item := (*s)[l-1]
	*s = (*s)[:l-1]
	return item, nil
}
