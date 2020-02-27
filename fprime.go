
	"fmt"
	"os"
	"strconv"
)

func FPRIME() {
	div := 2 

	arr := os.Args[1:]

	if len(arr) > 1 || len(arr) < 1 {
		fmt.Println()
		return
	}
	x, _ := strconv.Atoi(arr[0])

	if x == 1 { 
		fmt.Println()
		return
	}
	if x < 1 { 
		fmt.Println()
		return
	}
	for div <= x { 
		if x%div == 0 {
			fmt.Print(div) 
			if x == div {  
				fmt.Println() 
				return       
			}
			fmt.Print("*") 
			x /= div      
			div = 1        
		}
		div++ 
	}
}
