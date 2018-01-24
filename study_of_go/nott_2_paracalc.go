package main

import "fmt"

func sum(values [] int, resultchan chan int) {
	sum := 0
	for _,value := range values{
		sum += value
	}
	resultchan <- sum
}

func main() {
	values := [] int{1,2,3,4,5,6,7,8,9,10}

	resultchan := make(chan int, 2)
	go sum(values[:len(values)/2], resultchan)
	go sum(values[len(values)/2:], resultchan)
	sum1, sum2 := <-resultchan, <-resultchan

	fmt.Println("Result:",sum1, sum2, sum1+sum2)
}