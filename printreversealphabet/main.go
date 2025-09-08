package main

import "github.com/01-edu/z01"

func main() {
	for i := 25; i > -1; i-- {
		rvalpha := i + 'a'
		z01.PrintRune(rune(rvalpha))
	}
	z01.PrintRune('\n')
}
