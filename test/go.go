package main

import "fmt"

func main(){

	cox := "abcdefghijklmnopqrstuvwxyz"

	pox := []byte(cox)

	pox[1] = 'B'
	pox[3] = 'D'
	pox[5] = 'F'
	pox[7] = 'H'
	pox[9] = 'J'
	pox[11] = 'L'
	pox[13] = 'N'
	pox[15] = 'P'
	pox[17] = 'R'
	pox[19] = 'T'
	pox[21] = 'V'
	pox[23] = 'X'
	pox[25] = 'Z'

	shmox := string(pox)
	fmt.Println(shmox)
	
}