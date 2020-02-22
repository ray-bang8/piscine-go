
	"fmt"
	"os"
	"strconv"
)

func FPRIME() {
	div := 2 //Объявить делитель

	arr := os.Args[1:]

	if len(arr) > 1 || len(arr) < 1 {
		fmt.Println()
		return
	}
	x, _ := strconv.Atoi(arr[0])

	if x == 1 { //Если приходит 1, то отвечаем 1
		fmt.Println()
		return
	}
	if x < 1 { //Если меньше 1, то выходим
		fmt.Println()
		return
	}
	for div <= x { //Пока делитель не равен и меньше числа
		if x%div == 0 { //Если число делится на делитель без остатка
			fmt.Print(div) //Напечатать число
			if x == div {  //Если делитель равен числу
				fmt.Println() //Новая строка
				return        //Выход
			}
			fmt.Print("*") //Пишем знак умножения
			x /= div       //Число делим на делитель
			div = 1        //Делитель равняем 1
		}
		div++ //Увеличиваем делитель
	}
}
