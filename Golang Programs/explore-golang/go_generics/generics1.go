package main

import ("fmt")

type Num2 int

// type Num interface {
// 	int | float64 
// }

// // Function name ke baju me hi declare karne padege
// func Add[T int](a T, b T) T{
// 	fmt.Println(a+ b)
// 	c := a + b 
// 	fmt.Printf("type %T", c)
// 	return c
// }

type Num interface {
	int
}

func Add2(a Num, b Num){
	fmt.Println("simple add", a, b)
}


func main(){
	
	var a,b Num
	a = 10
	b = 20
		
	Add2(a, b)
}