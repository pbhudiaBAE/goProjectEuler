package main

import ("fmt"
		"math")

func main() {

	//found:=false

Loop:for j:=1;j<1021;j++{
		for k:=1;k<2168;k++{
			query1:= Pentagonal(j) + Pentagonal(k)
			query2:= Pentagonal(k)-Pentagonal(j)
			if (isPentagonal(query1)==true && isPentagonal(query2)==true) {
				fmt.Println(j,Pentagonal(j))
				fmt.Println(k,Pentagonal(k))
				fmt.Println(Pentagonal(k)-Pentagonal(j))
				break Loop
			}

		}
	}

fmt.Println(isPentagonal(8602840))
fmt.Println(isPentagonal(1147))


}

//boolean function to check if an integer is a Pentagonal number 
func isPentagonal(n int) bool{

	var temp float64= (1+math.Sqrt(1+24*float64(n)))/6
	if (temp==float64(int(temp))) {
		return true
	}else{
		return false
	}
}

//function to return the nth Pentagonal number 
func Pentagonal(n int) int{
	return n*((3*n)-1)/2
}

