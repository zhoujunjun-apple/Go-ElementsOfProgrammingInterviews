package main

import (
	"sync"
)

//Node represent a single list node
type Node struct {
	Value int
	Next  *Node
}

//SingleList represent a single list
type SingleList struct {
	Length     int   //node number
	Header     *Node //header node pointer
	GuardMutex sync.Mutex
}

//Add function add nodes at the head of single list
func (sl *SingleList) Add(values []int) (int, error) {
	var header **Node = &sl.Header

	for _, val := range values {
		node := Node{
			Value: val,
			Next:  *header,
		}

		sl.GuardMutex.Lock()
		*header = &node
		sl.Length++
		sl.GuardMutex.Unlock()
	}

	return sl.Length, nil
}

//GetValues function returns all values in single list
func (sl *SingleList) GetValues() []int {
	ret := []int{}
	header := sl.Header
	for header != nil {
		ret = append(ret, header.Value)
		header = header.Next
	}
	return ret
}

//NativeMerge function merge two sorted single list into one
func NativeMerge(left *SingleList, right *SingleList) *SingleList {
	ret := SingleList{}
	retTailer := &ret.Header

	leftHeader, rightHeader := left.Header, right.Header
	for leftHeader != nil && rightHeader != nil {
		leftVal := leftHeader.Value
		rightVal := rightHeader.Value

		node := Node{
			Next: nil,
		}

		if leftVal < rightVal {
			node.Value = leftVal
			leftHeader = leftHeader.Next
		} else {
			node.Value = rightVal
			rightHeader = rightHeader.Next
		}

		ret.GuardMutex.Lock()
		*retTailer = &node     //append new node at tail
		retTailer = &node.Next //update retTailer to tail
		ret.Length++
		ret.GuardMutex.Unlock()
	}

	for leftHeader != nil {
		node := Node{
			Value: leftHeader.Value,
			Next:  nil,
		}
		ret.GuardMutex.Lock()
		*retTailer = &node
		retTailer = &node.Next
		ret.Length++
		ret.GuardMutex.Unlock()
		leftHeader = leftHeader.Next
	}

	for rightHeader != nil {
		node := Node{
			Value: rightHeader.Value,
			Next:  nil,
		}
		ret.GuardMutex.Lock()
		*retTailer = &node
		retTailer = &node.Next
		ret.Length++
		ret.GuardMutex.Unlock()
		rightHeader = rightHeader.Next
	}

	return &ret
}

func main() {

}
