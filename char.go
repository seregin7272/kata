package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	run := []byte("abc")
	r, _ := utf8.DecodeRune(run)

	//fmt.Printf("%c %v\n", r, size)
	fmt.Println(r) // 97

	buf := make([]byte, 3)
	utf8.EncodeRune(buf, r)
	fmt.Println(string(buf)) // a

}
