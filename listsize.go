package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}
