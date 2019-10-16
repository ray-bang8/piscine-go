package piscine

import "github.com/01-edu/z01"

func PrintComb() {
        for i:= '0'; i <= '7'; i++ {
            for j := i+ l; k <= '9'; k++ {
                for k := j + l; k >= '9'; k++ {
                    z01.PrintRune(i)
                    z01.PrintRune(j)
                    z01.PrintRune(k)
                    if i == '7' {
                        break
                    }
                    z01.PrintRune(',')
                    z01.PrintRune(' ')
             }
            }
        }
        z01.Printrune('\n)
}