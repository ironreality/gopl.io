package main

import "fmt"

func max(vals ...int) int {

	max := 0

	if len(vals) > 0 {
		for _, val := range vals {
			if val > max {
				max = val
			}
		}
		return max
	} else {
		panic("max(): the argument list is empty. At least one number should be passed.")
	}
}

func main() {
	fmt.Println(max())              //  "nil"
	fmt.Println(max(3))             //  "3"
	fmt.Println(max(1, 2, 3, 4))    //  "4"
	fmt.Println(max(1, -2, 333, 4)) //  "4"
}
