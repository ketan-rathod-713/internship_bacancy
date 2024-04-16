package main

import (
	"fmt"
	"io"
	"os"
)

// func WriteString ¶
// func WriteString(w Writer, s string) (n int, err error)
// WriteString writes the contents of the string s to w, which accepts a slice of bytes. If w implements StringWriter, its WriteString method is invoked directly. Otherwise, w.Write is called exactly once.

func main(){
	n, err := io.WriteString(os.Stdout, "Str ");
	if(err != nil){
		fmt.Println("error occured")
	} else {
		fmt.Println("success done with returned n ", n)
		// here n is the number of characters written
	}
}