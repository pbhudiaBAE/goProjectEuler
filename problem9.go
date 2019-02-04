package main

import (
	"fmt"
	"math"
)

func main() {
	squares := []int{}
	for i := 1; i < 1000000; i++ {
		s := i * i
		squares = append(squares, s)
	}
	All: for j := range squares {
		for k := range squares {
			cSquared := squares[j] + squares[k]
			if isInSlice(cSquared, squares) == true {
				if (math.Sqrt(float64(squares[j]))+math.Sqrt(float64(squares[k]))+math.Sqrt(float64(cSquared)))==1000 {
					fmt.Println(squares[j], squares[k], squares[j]+squares[k])
					break All
				}
			}
		}
		fmt.Println(isInSlice(3, squares))
	}
}

//function which takes a num and checks if its in a slice
func isInSlice(num int, s []int) bool {
	for i := range s {
		if num == s[i] {
			return true
		}
	}
	return false
}