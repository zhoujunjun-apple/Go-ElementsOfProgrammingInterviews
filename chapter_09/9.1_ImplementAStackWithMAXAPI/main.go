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

// StackMax struct support Max() operation in O(1) time and O(N) space
type StackMax struct {
	container []int //container used for saving element in stack
	//maxer used for recording the maximum value of current stack.
	//the length of maxer and container is the same
	//update this slice in Push operation
	maxer  []int
	length int
}

// Empty function check if stack is empty
func (sm *StackMax) Empty() bool {
	return sm.length > 0
}

// Push function add new value i onto stack and update the maxer array
func (sm *StackMax) Push(i int) {
	sm.container = append(sm.container, i)
	sm.length++
	if len(sm.maxer) > 0 {
		maxNow := sm.maxer[len(sm.maxer)-1]
		if maxNow < i {
			sm.maxer = append(sm.maxer, i)
		} else {
			sm.maxer = append(sm.maxer, maxNow)
		}
	} else {
		// sm.maxer is empty
		sm.maxer = append(sm.maxer, i)
	}
}

// Pushs function helps to build a stack fastly
func (sm *StackMax) Pushs(is []int) {
	for _, i := range is {
		sm.Push(i)
	}
}

// Top function peek the top element in stack
func (sm *StackMax) Top() (int, error) {
	if sm.Empty() {
		return -1, ErrstackEmpty
	}
	return sm.container[sm.length-1], nil
}

// Pop function return the top element in stack and remove it
func (sm *StackMax) Pop() (int, error) {
	if sm.Empty() {
		return -1, ErrstackEmpty
	}

	top, _ := sm.Top()

	sm.length--
	sm.container = sm.container[:sm.length]
	sm.maxer = sm.maxer[:sm.length]
	return top, nil
}

// Max function get the maximum value of current stack in O(1) time and O(1) space
func (sm *StackMax) Max() (int, error) {
	if sm.Empty() {
		return -1, ErrstackEmpty
	}
	return sm.maxer[sm.length-1], nil
}

//StackFastMax struct improve the space complexity of maxer variable compared with StackMax type
//Under the Best-Case, the space complexity of Max operation is O(1) other than O(N) in StackMax type
type StackFastMax struct {
	containter []int
	length     int
	maxer      []MaxItem
}

//MaxItem type record the current max value and it's occurrance time
type MaxItem struct {
	max   int //the max value in current stack
	count int //the count of max value in current stack
}

//Empty function check if stack is empty
func (sfm *StackFastMax) Empty() bool {
	return sfm.length > 0
}

//Push function add element at the top of stack and update maxer
func (sfm *StackFastMax) Push(i int) {
	//add element at the top of stack
	sfm.containter = append(sfm.containter, i)
	sfm.length++

	//update maxer
	if len(sfm.maxer) > 0 {
		last := sfm.maxer[len(sfm.maxer)-1]
		if last.max < i {
			newMax := MaxItem{max: i, count: 1}
			sfm.maxer = append(sfm.maxer, newMax)
		} else {
			last.count++
			sfm.maxer[len(sfm.maxer)-1] = last
		}
	} else {
		first := MaxItem{max: i, count: 1}
		sfm.maxer = append(sfm.maxer, first)
	}
}

//Pushs call Push function to add batch elements
func (sfm *StackFastMax) Pushs(is []int) {
	for _, i := range is {
		sfm.Push(i)
	}
}

//Top function peek the top element of stack
func (sfm *StackFastMax) Top() (int, error) {
	if sfm.Empty() {
		return -1, ErrstackEmpty
	}
	top := sfm.containter[sfm.length-1]
	return top, nil
}

//Pop function return the top element and remove it from stack,
//then update the maxer
func (sfm *StackFastMax) Pop() (int, error) {
	if sfm.Empty() {
		return -1, ErrstackEmpty
	}

	//update container
	sfm.length--
	top := sfm.containter[sfm.length]
	sfm.containter = sfm.containter[:sfm.length]

	//update maxer
	nowMax := sfm.maxer[len(sfm.maxer)-1]
	if nowMax.count > 1 {
		//the max value remain unchanged, only need to update count
		nowMax.count--
		sfm.maxer[len(sfm.maxer)-1] = nowMax
	} else {
		//top is the last max value of nowMax.max in stack
		sfm.maxer = sfm.maxer[:len(sfm.maxer)-1]
	}

	return top, nil
}

//Max function get the current max value of stack
func (sfm *StackFastMax) Max() (int, error) {
	if sfm.Empty() {
		return -1, ErrstackEmpty
	}
	return sfm.maxer[len(sfm.maxer)-1].max, nil
}

func main() {

}
