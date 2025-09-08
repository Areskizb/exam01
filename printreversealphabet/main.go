package main

import "github.com/01-edu/z01"

func main() {
	for i := 25; i > -2; i-- {
		if i == 25 {
			z01.PrintRune('z') // pas de casting ? mrc z01
		}
		if i == 23 {
			z01.PrintRune('y')
		}
		if i == 22 {
			z01.PrintRune('x')
		}
		if i == 21 {
			z01.PrintRune('w')
		}
		if i == 20 {
			z01.PrintRune('v')
		}
		if i == 19 {
			z01.PrintRune('u')
		}
		if i == 18 {
			z01.PrintRune('t')
		}
		if i == 17 {
			z01.PrintRune('s')
		}
		if i == 16 {
			z01.PrintRune('r')
		}
		if i == 15 {
			z01.PrintRune('q')
		}
		if i == 14 {
			z01.PrintRune('p')
		}
		if i == 13 {
			z01.PrintRune('o')
		}
		if i == 12 {
			z01.PrintRune('n')
		}
		if i == 11 {
			z01.PrintRune('m')
		}
		if i == 10 {
			z01.PrintRune('l')
		}
		if i == 9 {
			z01.PrintRune('k')
		}
		if i == 8 {
			z01.PrintRune('j')
		}
		if i == 7 {
			z01.PrintRune('i')
		}
		if i == 6 {
			z01.PrintRune('h')
		}
		if i == 5 {
			z01.PrintRune('g')
		}
		if i == 4 {
			z01.PrintRune('f')
		}
		if i == 3 {
			z01.PrintRune('e')
		}
		if i == 2 {
			z01.PrintRune('d')
		}
		if i == 1 {
			z01.PrintRune('c')
		}
		if i == 0 {
			z01.PrintRune('b')
		}
		if i == -1 {
			z01.PrintRune('a') // mdrrrrrrrrrrrrr le code bien sale
		}
	}
	z01.PrintRune('\n')
}
