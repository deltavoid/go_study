package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(&arr))
}

// pointer
func sum(arr *[5]int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}
