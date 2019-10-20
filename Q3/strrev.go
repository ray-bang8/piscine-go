package piscine

func StrRev(s string) string {
	var cox string
	for _, shmox := range s {
		cox = string(shmox) + cox
	}
	return cox

}

/*func main() {
	s := "Hello World!"
	s := StrRev(s)
	Swap(&a, &b)
	fmt.Println(s)

}*/
