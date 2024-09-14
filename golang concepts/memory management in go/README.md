# Memory Management In Go

The role of memory management tool is to give illusion of infinite memory available for developer. it does things underhood like

- memory allocation
- marking of live objects
- freeing space used by dead objects

## Garbage Collection

- garbage generation
- garbage collection
- measuring performance
- configuring the garbage collector

Stack allocation is cheap and doesn't require garbage collection. It is essential to know that what part goes to the stack and what part goes to the heap.

For eg.

```
func main(){
NewDuck()
}

type Duck struct{}

func NewDuck() Duck {
	return Duck{}
}

```

Above example doesn't create any pointers or referencies and hence it only creates objects inside stack and returns.

Let's take another example.

```
func main(){
NewDuck()
}

type Duck struct{}

func NewDuck() *Duck {
	return &Duck{}
}
```

Above examples stores the object inside the heap memory and reference inside the stack memory. Hence in this case we require a garbage collection to be done inside the heap memory. as for stack memory it will get automatically removed once the function completes it's execution.

**What do mean by escape analysis in compiler ?!**

It checks weather the object will go to heap or stack.

Reference :- https://en.wikipedia.org/wiki/Escape_analysis

In [compiler optimization](https://en.wikipedia.org/wiki/Compiler_optimization "Compiler optimization"), **escape analysis** is a method for determining the dynamic scope of [pointers](https://en.wikipedia.org/wiki/Pointer_(computer_programming)) "Pointer (computer programming)") –  where in the program a pointer can be accessed. It is related to [pointer analysis](https://en.wikipedia.org/wiki/Pointer_analysis "Pointer analysis") and [shape analysis](https://en.wikipedia.org/wiki/Shape_analysis_(program_analysis)) "Shape analysis program analysis".

**How to check the escape analysis of compiler** 

`go run -gcflags -m main.go`

It will output the escaped pointers and related information.

**Trigger the garbage collector**
