package main

import (
	"fmt"
	student "piscine"
)

func main() {
	fmt.Println(student.CountChar("Hello World", 'l'))
	fmt.Println(student.CountChar("5  balloons", 5))
	fmt.Println(student.CountChar("   ", ' '))
	fmt.Println(student.CountChar("The 7 deadly sins", '7'))
}
