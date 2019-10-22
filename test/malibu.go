package main

import "github.com/01-edu/z01"

func Raid1c(x, y int) {
	if x < 0 || y < 0 {
	} else {
		if x == 0 || y == 0 {
		} else {
			if x == 1 && y == 1 {
				z01.PrintRune('A')
				z01.PrintRune('\n')
			} else {
				if y == 1 {
					for i := 1; i <= x; i++ {
						if i == 1 || i == x {
							z01.PrintRune('A')
						} else {
							z01.PrintRune('B')
						}
					}
					z01.PrintRune('\n')
				} else {
					for i := 1; i <= x; i++ {
						if i == 1 || i == x {
							z01.PrintRune('A')
						} else {
							z01.PrintRune('B')
						}
					}
					z01.PrintRune('\n')
					for j := 1; j <= y-2; j++ {
						for k := 1; k <= x; k++ {
							if k == 1 || k == x {
								z01.PrintRune('B')
							} else {
								z01.PrintRune(' ')
							}

						}
						z01.PrintRune('\n')
					}
					for i := 1; i <= x; i++ {
						if i == 1 || i == x {
							z01.PrintRune('C')
						} else {
							z01.PrintRune('B')
						}
					}
					z01.PrintRune('\n')
					
				}
			}
		}
	}
}
func main() {
	Raid1c(5, 5)
}
