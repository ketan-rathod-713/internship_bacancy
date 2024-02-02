package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l3 := new(ListNode)
	l1.Next = l2
	l2.Next = l3
	l1.Val = 1
	l2.Val = 2
	l3.Val = 3

	c1 := new(ListNode)
	c2 := new(ListNode)
	c3 := new(ListNode)
	c1.Next = c2
	c2.Next = c3
	c1.Val = 2
	c2.Val = 3
	c3.Val = 9

	printList(l1)
	printList(c1)
	d := addTwoNumbers(l1, c1)
	printList(d)
}

func printList(ptr *ListNode) {
	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Println("")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	curr_val := 0 

	for l1!= nil && l2 != nil {
		curr_val = l1.Val + l2.Val + carry 
		carry = curr_val / 10
		
	}

	return nil
}
