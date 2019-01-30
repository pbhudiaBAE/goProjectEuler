package main

import "fmt"

func main() {

amicable := []int{}

for i:=0; i<10000; i++ {
	dofI:=d(i)
	if (d(dofI) == i && i != dofI && !contains(amicable,dofI)) {
		amicable = append(amicable,i,dofI)
	}
}

fmt.Println(amicable)

}

func contains (s []int, e int) bool {
	for i:=0; i< len(s); i++ {
		if s[i]==e {return true}
	}
	return false
}


//function d which takes an int and returns the sum of its proper divisors
func d(num int) int {

	sum:=0

	for i:=1; i<=num/2; i++ {
		if (num%i == 0){
			sum=sum+i
		} 
	}
	return sum
}

