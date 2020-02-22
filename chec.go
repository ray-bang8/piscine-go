package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "92233720368547758082222"
	t, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(0)
	} else {
		fmt.Println(t)
	}

}
b := 0
a, err := strconv.Atoi(os.Args[0])
if err != nil {
	b = 0
}
