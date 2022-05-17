package main

import (
	"fmt"
	"strings"
)

func main() {
	var input = []string{"apple", "pie", "apple", "red", "red", "red"}

	fmt.Println(freq(input))

}

func freq(arr []string) string {
	element := ""
	count := 0
	for j := 0; j < len(arr); j++ {
		tempElement := arr[j]
		tempCount := 0

		for i := 0; i < len(arr); i++ {
			if strings.EqualFold(arr[i], tempElement) {
				tempCount++
			}

		}
		if tempCount > count {
			element = tempElement
			count = tempCount
		}
	}
	return element

}
