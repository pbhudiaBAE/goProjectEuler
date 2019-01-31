package main

import "fmt"

func main() {
	biggest:=0
	temp:=0
	for i:=2;i<1000000;i++{
		if (len(collatz(i))>biggest){
			biggest=len(collatz(i))
			temp=i
		}
	}
	fmt.Println(temp,biggest,collatz(temp))
}

func collatz(num int) []int{
	sequence:= [] int{num}

	for {
		if(num % 2 == 0){
			num=num/2
		} else{
			num=(3*num)+1
		}
		sequence=append(sequence,num)
		if(num==1){
			break
		}
	}
	return sequence
}
	