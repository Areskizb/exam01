package main

import (
	"fmt"
	student "piscine"
)

func main() {
	fmt.Println(student.CountAlpha("Hello world"))
	fmt.Println(student.CountAlpha("H e l l o"))
	fmt.Println(student.CountAlpha("H1e2l3l4o"))
}
