# Nil

- nil is (a) zero.

History of nil 
- C.A.R Hooare Communicating sequential processes. it became basis for the concurrent programming in go. hE also created a nil in 1965.
- After designing nil once he said that i call it my billion-dollar mistake. it was the invention of the ull reference in 1965.
- At that time he was designing the first comprehensive type system for the references in the object oriented language.
- nil leads to panic.
- pointers, slices, maps, channels, functions and interfaces can get nil
- zero value of struct is a empty struct with zero value of all fields.
- the type of nil ??
- nil has no type. like boolean, string, int etc..
- a := nil // use of untyped nil
- nil is a predeclared identifier.. nil is not a keyword among 25 one.
- var nil = errors.New("asga") // never do this. it is doable hence nil is not keyword
- Pointers
    - they point to position in memory
    - in go No pointer arithmatic -> memory safety
    - garbage collection
    - nil pointer
        - points to nothing.
- Slice
    - it has ptr, len and cap
    - nil slice means it doesnt have backing array. hence overall nil reference.
- Maps channels and functions are pretty much of the same styles. maps and channels used to be a pointers. but then we changed that. it is still pointers just not adding *.
- nil for Channel, map, function : pointer without the initalization

- Interfaces : it is not a pointers. 2 components (type, value)
- (nil, nil) equals nil in stringer interfacer.
for eg.
    var s fmt.Stringer
    fmt.Println(s == nil) // true
- we have pointer to person and it is nil and it satisfy the stringer interface

var p *person // nil of type *person
var s fmt.Stringer = p // Stringer(*person, nil)
fmt.Println(s == nil) // false

- (*Person, nil) is not nil. as value is nil still.

- Dont declare concrete error variables. Avoid it.

- nil is not a nil for some kinds of nil.

- IMPORTANT : nil slice is a different to the nil interface and so on.. for all.

- may be it would be less confusing if we would have a different names or it is good.

## Kinds of nil

- pointers points to nothing
    usecase:
    only can compare with nil nothing else.

- slices have to backing array
    - s[i] index out of..
    - append on nil slices and it is okay..
    - at the beggining cap 0 then it will increase and so on..

- maps are initialised
    - var m map[t]u
    - len(m) // 0
    - for range m // iterates zero items
    - v, ok := m[i] // zero(u) , false
    - m[i] = x // panic assignment to entry in nil map
    - if your function is only reading map values then pass the nil values to it if possible.

- channels not initialized
    - var c chan t
    - <-c // blocks forever
    - c <-x // blocks forever
    - close(c) // panic : close of nil channel

    - we can set channel nil to block it forever inside select statement. which is exactly opposite of initalized closed channel.
    - Hence use nil channels to disable select cases.


- functions are not initalized
    - to signal default value

- interfaces have no value assigned, not even a nil pointer
    - nil values can satisfy interfaces
    - nil values and default values ( just like functions, if accepted data is nil then use default values )
    - nil means go with the default values.
    - use nil interfaces to signal default interfaces. just like that of http.HandleFunc(":8080", nil)

'make the zero value useful' - rob pike

