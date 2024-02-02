package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln := new(ListNode) // it returns the address of the Node
	fmt.Println(ln, reflect.TypeOf(ln))

	ln1 := new(ListNode) // it returns the address of the Node
	ln.Next = ln1

	for ln != nil {
		fmt.Println(ln)
		ln = ln.Next
	}

	ll := new(LinkedList)
	ll.AddNode(10)
	ll.AddNode(20)
	ll.AddNode(30)
	ll.AddNodeAtBegining(40)
	ll.AddNodeAtBegining(50)
	ll.AddNodeAtBegining(60)
	ll.PrintLinkedList()
	fmt.Println(ll.Length())
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

// we will get pointer to Node block
func (ll *LinkedList) AddNode(val int) {
	if ll.Head == nil { // create new LL
		node := new(ListNode)
		node.Val = val
		ll.Head = node // SET variables
		ll.Tail = node
	} else { // already linked list he
		ll.AddNodeAtEnd(val)
	}
}

func (ll *LinkedList) AddNodeAtEnd(val int){
	node := new(ListNode)
    node.Val = val
    ll.Tail.Next = node
    ll.Tail = node
}

func (ll *LinkedList) AddNodeAtBegining(val int){
	node := new(ListNode)
    node.Val = val
	node.Next = ll.Head 
	ll.Head = node
}

func (ll *LinkedList) Length() int {
	node := ll.Head
	ln := 0 // length 
	for node != nil {
		ln++
		node = node.Next
	}

	return ln
}

func (ll *LinkedList) PrintLinkedList() {
	node := ll.Head
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}

	fmt.Println("")
}
