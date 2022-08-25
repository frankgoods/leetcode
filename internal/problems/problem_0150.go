package problems

import "strconv"

type Stack []int

func NewStack() *Stack {
	return (*Stack)(&[]int{})
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	v := (*s)[last]
	*s = (*s)[0:last]
	return v
}

func calc(v1, v2 int, op string) int {
	switch op {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	default:
		panic(op)
	}
}

func evalRPN(tokens []string) int {
	s := NewStack()
	for _, t := range tokens {
		switch t {
		case "+", "-", "*", "/":
			v2 := s.Pop()
			v1 := s.Pop()
			s.Push(calc(v1, v2, t))
		default:
			v, _ := strconv.Atoi(t)
			s.Push(v)
		}
	}
	return s.Pop()
}
