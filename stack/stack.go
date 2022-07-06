package stack

type stack struct {
	vec []int
}

// type iStack interface {
// 	push(int)
// 	pop() int
// 	size() int
// }

func (s *stack) Push(x int) {
	s.vec = append(s.vec, x)
}

func (s *stack) Pop() int {
	x := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return x
}

func (s *stack) Size() int {
	return len(s.vec)
}

func NewStack() stack {
	return stack{}
}
