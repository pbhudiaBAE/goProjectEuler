package main

import ("fmt"
		"strconv"
		"reflect"
		"sort")

func main() {

	for i:=1;;i++ {
		if (sixLoop(i)==true) {
			fmt.Println(i)
			break
		}
		
	}

}

// func that takes in a number and splits its digits into an ordered array
func numSplit(num int) []int {
	numLen := len(strconv.Itoa(num))
	digits:=[]int{}
	for i:=0; i<numLen; i++{
		digits=append(digits,num%10)
		num=num/10

	}
	sort.Ints(digits)
	return digits
}

//func to input a slice and compare individual digits with another slice

func compare(a []int, b []int) bool {
return reflect.DeepEqual(a,b)
	
}

//func that takes in an int to be tested, and returns true if its 6 multiples all contain the same digits
func sixLoop(a int) bool{
	for i:=2;i<7;i++{
		if(compare(numSplit(a),numSplit(a*i))==false){
			return false
		}
	}
	return true
}
