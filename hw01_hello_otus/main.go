package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	otus := "Hello, OTUS!"
	suto := stringutil.Reverse(otus)
	fmt.Println(suto)
}
