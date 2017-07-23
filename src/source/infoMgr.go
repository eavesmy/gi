package source

import (
	"fmt"
)

func InsertINFO(infos *[]string) {
	// fmt.Println(infos)

	if len(*infos) != 0 {
		fmt.Println(infos)
	}
}
