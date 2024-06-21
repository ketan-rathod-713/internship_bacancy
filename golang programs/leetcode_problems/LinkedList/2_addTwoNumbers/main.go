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
	l1.Val = 2
	l2.Val = 4
	l3.Val = 3

	c1 := new(ListNode)
	c2 := new(ListNode)
	c3 := new(ListNode)
	c1.Next = c2
	c2.Next = c3
	c1.Val = 5
	c2.Val = 6
	c3.Val = 4

	printList(l1)
	printList(c1)

	d := addTwoNumbers(l1, c1)
	printList(d)

}

func printList(ptr *ListNode) {
	if ptr == nil {
		fmt.Println("Ptr is Nil")
		return
	}
	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}
	fmt.Println("")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	curr_val := 0

	ans := new(ListNode)
	temp := ans

	for l1 != nil && l2 != nil {
		curr_val = l1.Val + l2.Val + carry
		carry = curr_val / 10

		newNode := new(ListNode)
		newNode.Val = curr_val % 10
		temp.Next = newNode
		temp = temp.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		curr_val = l1.Val + carry
		carry = curr_val / 10

		newNode := new(ListNode)
		newNode.Val = curr_val % 10
		temp.Next = newNode
		temp = temp.Next

		l1 = l1.Next
	}

	for l2 != nil {
		curr_val = l2.Val + carry
		carry = curr_val / 10

		newNode := new(ListNode)
		newNode.Val = curr_val % 10
		temp.Next = newNode
		temp = temp.Next

		l2 = l2.Next
	}

	for carry != 0 {
		newNode := new(ListNode)
		newNode.Val = carry
		temp.Next = newNode
		temp = temp.Next
		carry = 0
	}

	return ans.Next
}
