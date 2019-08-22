package main
import (
	"fmt"
	"unicode"
)

func string_check (s string) string {
	var check []int
	for _, r := range s {
        if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
            return "Invalid Input"
		} else {
			if unicode.IsLower(r) {
				if len(check) ==0 || check[len(check)-1] == 0 {
					check = append(check,0)
				}else{
					return "Mix"
				}
					
			} else {
				if len(check) == 0 || check[len(check)-1] == 1 {
					check = append(check,1)
				} else {
					return "Mix"
				}
			}
		}
    }
	
	if check[len(check)-1] == 0 {
		return "All Small Letter"
	} else {
		return "All Capital Letter"
	}
	
}

func main() {
	fmt.Println(string_check("ABCDEZXY"))
	fmt.Println(string_check("abcdezxy"))
	fmt.Println(string_check("AbCdEzXy"))
	fmt.Println(string_check("A1b2C3d4/+*"))

}