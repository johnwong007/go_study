package main

import (
	"fmt"
)

func main() {
	var pstr = new([3]int)
	fmt.Println(pstr)
	for i := 0; i < len(*pstr); i++ {
		pstr[i] = i
	}

	fmt.Println(pstr)

}
