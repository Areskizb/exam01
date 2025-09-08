package student

import "github.com/01-edu/z01"

func PrintIf(str string) string {
	lenCount := len(str)
	if lenCount >= 3 || lenCount == 0 {
		z01.PrintRune('G')
		z01.PrintRune('\n')
	} else {
		print("Invalid Input")
		z01.PrintRune('\n')
	}
	return "\x00"
}
