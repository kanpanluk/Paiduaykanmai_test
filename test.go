package main
import (
	"fmt"
	"strings"
)

func main() {
	find_age := "11-20"

	find_age_min := strings.Split(find_age,"-")[:1][0]
	var find_age_max string

	if find_age_max = find_age_min ; len(strings.Split(find_age,"-")[1:]) > 0 {
		find_age_max = strings.Split(find_age,"-")[1:][0]
	}

	fmt.Println(find_age_min,find_age_max)
}