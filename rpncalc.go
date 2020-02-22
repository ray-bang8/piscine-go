

import (
	"fmt"
	"os"
	"strconv"
)

func RPNCALC() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error")
		return
	}

	arg := args[0]

	// array := strings.Split(arg, " ")
	// fmt.Printf("%q\n", array)
	array := SplitWhiteSpaces(arg)
	countnum := 0
	countop := 0

	for _, x := range array {
		if IsNumber(x) {
			countnum++
		} else {
			countop++
		}
	}

	first := 0
	second := first + 1
	op := second + 1

	new := 0
	if countnum == countop+1 && len(array) >= 3 {
		for first < len(array)-2 {
			if IsNumber(array[first]) && IsNumber(array[second]) && IsOper(array[op]) {
				// operator := array[op]
				n1, _ := strconv.Atoi(array[first])
				n2, _ := strconv.Atoi(array[second])
				switch array[op] {
				case "+":
					new = n1 + n2
					array[op] = Itoa(n1 + n2)
				case "-":
					new = n1 - n2
					array[op] = Itoa(n1 - n2)
				case "/":
					new = n1 / n2
					array[op] = Itoa(n1 / n2)
				case "*":
					new = n1 * n2
					array[op] = Itoa(n1 * n2)
				}

				array = append(array[:op-2], array[op:]...)

				first = 0
				second = first + 1
				op = second + 1

			} else {
				first++
				second++
				op++
			}
		}

	} else {
		fmt.Println("Error")
	}
	fmt.Println(new)

}

func IsNumber(a string) bool {
	_, err := strconv.Atoi(a)
	if err != nil {
		return false
	}
	return true
}

func IsOper(a string) bool {
	if a == "+" || a == "-" || a == "/" || a == "*" {
		return true
	}
	return false
}

func Itoa(n int) string {
	res := ""
	for n > 0 {
		temp := n % 10
		res = string(temp+48) + res
		n /= 10
	}
	return res
}

func SplitWhiteSpaces(str string) []string {
	prev := ' '
	count := 0
	for _, v := range str {
		if v != ' ' && prev == ' ' {
			count++
		}
		prev = v
	}
	ar := make([]string, count)
	i := 0
	word := ""
	for _, v := range str {
		if v == ' ' && word != "" {
			ar[i] = word
			word = ""
			if count != 1 {
				i++
			}
		} else if v != ' ' {
			word += string(v)
		}
	}
	if word != "" {
		ar[i] = word
	}
	return ar
}
