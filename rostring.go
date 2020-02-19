import (
	"fmt"
	"os"
	"strings"
)

func ROSTRING() {

	new := ""
	new2 := ""
	if len(os.Args) == 2 {
		str := os.Args[1]
		for i := 0; i < len(str); i++ {
			str = standart(str)
			if str[0] != 32 {
				for v := 0; v < len(str); v++ {

					if str[v] == 32 {
						new = str[:v]
						new2 = str[v:]
						str = new2 + " " + new
						fmt.Println(standart(str))
						return
					}
				}
			}
		}
	} else {
		fmt.Println()
	}
}
func standart(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}