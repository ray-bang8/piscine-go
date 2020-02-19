func GCD()
    if len(os.Args) == 3 {


    		s1 := os.Args[1]
    		s2 := os.Args[2]
    		fmt.Println(gcd(Atoi(os.Args[1]), Atoi(os.Args[2])))

    	} else {
    		z01.PrintRune('\n')
    	}
    }
    func gcd(a, b int) int {
    	if b == 0 {
    		return a
    	}
    	return gcd(b, a%b)
    }
    func Atoi(s string) int {
    	min := 0
    	plu := 0
    	sum := 0
    	for i, v := range s {
    		if v >= '0' && v <= '9' {
    			num := 0
    			for j := '0'; j < v; j++ {
    				num++
    			}
    			sum = sum*10 + num
    		} else if v == '-' && i == 0 {
    			min++
    		} else if v == '+' && i == 0 {
    			plu++
    		} else {
    			return 0
    		}

    	}
    	if min == 1 {
    		sum = -sum
    	}
    	return sum
    }