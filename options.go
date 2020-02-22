package main

import (
	"fmt"
	"os"
)

func main() {
	table := "00000000 00000000 00000000 00000000"
	options := "******zy xwvutsrq ponmlkji hgfedcba"
	help := "options: abcdefghijklmnopqrstuvwxyz"
	invalid := "Invalid Option"
	if len(os.Args) == 1 {
		fmt.Println(help)
		return
	}
	for _, k := range os.Args[1:] {
		for j, l := range k {
			if l == 'h' && j == 1 {
				fmt.Println(help)
				return
			}
			for i, m := range options {
				for range table {
					if l >= 'a' && l <= 'z' || l == '-' {
						if l == m {
							table = table[:i] + "1" + table[i+1:]
						}
					} else {
						fmt.Println(invalid)
						return
					}
				}
			}
		}
	}
	fmt.Println(table)
}
