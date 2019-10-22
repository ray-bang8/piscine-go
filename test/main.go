package main

import (
	"fmt"

	piscine ".."
)

func main() {
	str := "Cox 78 World!    4455 /"
	nb := piscine.AlphaCount(str)
	fmt.Println(nb)
}
