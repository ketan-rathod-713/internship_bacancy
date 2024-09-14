# Explore Golang

Here i will be summarising my golang code. This was the first phase in learning go.

What i learned in this first go ??!!!

1. Basic datatypes

    - int, float etc.
    - strings and runes

check datatype using reflect.Typeof

2. Conditionals

if else switch statement

3. Loops in golang

for using iterator, range for loop, for while loop

4. Functions
    - closures
    - variadic functions
    - defer keyword

5. Errors and custom errors
    - errors package and errors.Is method
    - error builtin interface implements `Error() string` method

6. Json marshallig and unmarshalling
    - marshal and unmarshal
    - indent
    - encoder and decoder
    - other usefull methods

7. Map Internals
    - https://www.youtube.com/watch?v=Tl7mi9QmLns
    - https://phati-sawant.medium.com/internals-of-map-in-golang-33db6e25b3f8
    - Pass by reference always does not change like slice

8. Slice Internals

    - len
    - cap
    - zero value is nil
    - slicing of exisiting slice or array returns slice
    - https://go.dev/blog/slices
    - copy function for copying slices
    - slice is a value not a pointer but inside it is pointing to a array hence for slice recivers should be pointers by default if we want to change slice else there will be no effect.


9. Pointers
    - same as other programming language
    - nil

10. Take input from user
    - fmt scan
    - bufio reader and writter implementations

11. Interface Internals
    - underhood tuple of (type, value)
    - pointer reciever

12. Others
    - struct embeddings
    - generics

## Packages To Look

- fmt : for formatted io and similar like printf and scanf of c language.
    - 
- bufio : for buffered io. provides Reader and Writter over io.Reader and io.Writter interfaces.

- encoding/json : 
    - Unmarshaller and Marshaller interfaces
    - TextMarshaler and TextUnmarshaler interfaces for text data of type
    - Marshal and Unmarshal functions
    - Valid reports validity of json string
    - Encoder and Decoder and their methods and Tokens etc.

- encoding/csv :
    - reader and writter implementation
    - FieldPos function understood
    - What is InputOffset function

- time :
    - Sleep
    - Formate

- TODO:
    - reflect
    - strings
    - strconv