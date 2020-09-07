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

//Merge function merge newList into sl
//use double scan method. O(N+M) time and O(1) space
func (sl *SingleList) Merge(newList *SingleList) {
	mainTailer := sl.Header
	subeTailer := newList.Header

	var insertTailer *Node = nil
	for subeTailer != nil {
		//find the position in sl to insert node
		for mainTailer != nil && mainTailer.Value <= subeTailer.Value {
			insertTailer = mainTailer
			mainTailer = mainTailer.Next
		}

		//find the insert position or sl has reach the tail
		if mainTailer != nil {
			//mainTailer != nil, find the insert position

			//check if insert at the head of sl
			if insertTailer == nil {
				//insertTailer == nil, new node from newList need to insert at the head of sl
				sl.Header = subeTailer
				subeTailer = subeTailer.Next
				sl.Header.Next = mainTailer
				insertTailer = sl.Header
			} else {
				//new node from newList need to insert after the insertTailer node
				insertTailer.Next = subeTailer
				subeTailer = subeTailer.Next
				insertTailer.Next.Next = mainTailer
				insertTailer = insertTailer.Next
			}

			//update Length field of sl and newList
			newList.GuardMutex.Lock()
			newList.Length--
			newList.GuardMutex.Unlock()

			sl.GuardMutex.Lock()
			sl.Length++
			sl.GuardMutex.Unlock()
		} else {
			//mainTailer == nil, sl has reach the tail, append the rest of newList onto sl
			if insertTailer == nil {
				sl.Header = subeTailer
			} else {
				insertTailer.Next = subeTailer
			}

			sl.GuardMutex.Lock()
			sl.Length += newList.Length
			sl.GuardMutex.Unlock()

			newList.GuardMutex.Lock()
			newList.Length = 0
			newList.GuardMutex.Unlock()
			break
		}
	}
}

//FastMerge function optimize the Merge() method
//use double pointer to eliminate the IF/ELSE statements of insertTailer
//O(N+M) time and O(1) space
func (sl *SingleList) FastMerge(newList *SingleList) {
	mainTailer := sl.Header
	subeTailer := newList.Header

	var insertTailer **Node = &sl.Header
	for subeTailer != nil {
		//find the position in sl to insert node
		for mainTailer != nil && mainTailer.Value <= subeTailer.Value {
			//the following two statements cannot exchange
			insertTailer = &(mainTailer.Next)
			mainTailer = mainTailer.Next
		}

		//find the insert position or sl has reach the tail
		if mainTailer != nil {
			//mainTailer != nil, find the insert position

			//new node from newList need to insert after the insertTailer node
			*insertTailer = subeTailer           //connect sl to new added node
			subeTailer = subeTailer.Next         //update subeTailer
			(*insertTailer).Next = mainTailer    //connect the new added node back to sl
			insertTailer = &(*insertTailer).Next //update insertTailer to next node

			//update Length field of sl and newList
			newList.GuardMutex.Lock()
			newList.Length--
			newList.GuardMutex.Unlock()

			sl.GuardMutex.Lock()
			sl.Length++
			sl.GuardMutex.Unlock()
		} else {
			//mainTailer == nil, sl has reach the tail, append the rest of newList onto sl
			*insertTailer = subeTailer

			sl.GuardMutex.Lock()
			sl.Length += newList.Length
			sl.GuardMutex.Unlock()

			newList.GuardMutex.Lock()
			newList.Length = 0
			newList.GuardMutex.Unlock()
			break
		}
	}
}

//NativeMerge function merge two sorted single list into one. O(N+M) time and O(N+M) space
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
