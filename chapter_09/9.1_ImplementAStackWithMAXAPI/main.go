package main

import "fmt"

var (
	//ErrstackEmpty represents error for empty stack
	ErrstackEmpty = fmt.Errorf("stack is empty")
)

//Stack type
type Stack struct {
	container []int
	Length    int
}

//Push function add a new element at the top of stack
func (s *Stack) Push(e int) {
	s.container = append(s.container, e)
	s.Length++
}

//Pushs function add multi elements into the stack
func (s *Stack) Pushs(es []int) {
	for _, e := range es {
		s.Push(e)
	}
}

//Pop function return and remove the top element
func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return -1, ErrstackEmpty
	} else {
		top := s.container[s.Length-1]
		s.container = s.container[:s.Length-1]
		s.Length--
		return top, nil
	}
}

//Top function retrieve, but does not remove the top element
func (s *Stack) Top() (int, error) {
	if s.Empty() {
		return -1, ErrstackEmpty
	} else {
		return s.container[s.Length-1], nil
	}
}

//Empty function check if stack is empty
func (s *Stack) Empty() bool {
	return s.Length < 1
}

//Max function return the max value in stack without change the stack structure. O(N) time and O(N) space
func (s *Stack) Max() (int, error) {
	if s.Empty() {
		return -1, ErrstackEmpty
	}

	helpStack := Stack{}
	max, _ := s.Top()

	for !s.Empty() {
		top, _ := s.Pop()
		if top > max {
			max = top
		}
		helpStack.Push(top)
	}

	for !helpStack.Empty() {
		top, _ := helpStack.Pop()
		s.Push(top)
	}

	return max, nil
}

func main() {

}
