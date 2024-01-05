package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}
	isMeaningful := false
	var min, max float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if !isMeaningful && err == nil {
			isMeaningful = true
			min = n
			max = n
		}
		if err != nil {
			continue
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	// 如果用户没有包括任何的有意义的数字
	if !isMeaningful {
		fmt.Println("Please include as least one number in the argument")
		return
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
