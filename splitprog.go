import (
	"fmt"
	"os"
)

func SPLITPROG() {
	if len(os.Args) == 3 {
		s := Split(os.Args[1], os.Args[2])
		fmt.Println(s)
	} else {
		fmt.Println()
	}
}

func Split(s1, s2 string) []string {
	resp := []string{}
	new := ""
	for i := 0; i < len(s1); i++ {
		if Check(s1, s2, i) {
			if new != "" {
				resp = append(resp, new)
				new = ""
				i = i + len(s2) - 1
			}
		} else {
			new = new + string(s1[i])
		}
	}
	if new != "" {
		resp = append(resp, new)
	}
	return resp
}
func Check(s1, s2 string, i int) bool {
	j := 0
	for j < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			return false
		}
		i++
		j++
	}
	return true
}