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
	l2.Val = 4
	l3.Val = 10
	printList(l1)

	c1 := new(ListNode)
	c2 := new(ListNode)
	c3 := new(ListNode)
	c1.Next = c2
	c2.Next = c3
	c1.Val = 1
	c2.Val = 2
	c3.Val = 3
	printList(c1)

	mergedList := mergeTwoLists(l1, c1)
	printList(mergedList)
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Edge cases
	if(list1 == nil && list2 == nil){
		return nil
	}

	ansList := new(ListNode)
	temp := ansList
	// normal case
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			temp.Next = list1
			temp = temp.Next
			list1 = list1.Next
		} else {
			temp.Next = list2
			temp = temp.Next
			list2 = list2.Next
		}
	}

	for list1 != nil {
		temp.Next = list1
		temp = temp.Next
		list1 = list1.Next
	}

	for list2 != nil {
		temp.Next = list2
		temp = temp.Next
		list2 = list2.Next
	}
	return ansList.Next
}
