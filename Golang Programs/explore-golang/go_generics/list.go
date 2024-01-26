package main

type List[T any] struct {
	head, tail *element[T]
}

// head and tail will point to one element type
// and Here List means -> DOUBT

// head and tail both are type of a element jo ki ek element ko define karte he
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(value T) { // methods for a List type // it has head and tail pointers // ohkk just like linkedlist
	if lst.tail == nil {
		lst.head = &element[T]{val: value}
		lst.head = lst.tail
	} else {
		lst.tail.next = &element[T]{val: value} // Ohk here we are passing an address rather then copying the value ha ha
		lst.tail = lst.tail.next
	}
}

// func (lst *List[T]) Pop(value T) {
// 	if lst.head == nil {
// 		fmt.Println("Can't pop ha ha")
// 	} else {
// 		if lst.head == lst.tail {
// 			lst.head = nil
// 			lst.tail = nil
// 		} else {

// 		}
// 	}
// }

func usingList() {

}
