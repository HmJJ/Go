package main

import "fmt"

func main() {
	a, b := 1, 2
	res, _ := Add(a, b)
	fmt.Println(res)
}


func Add(a, b int) (res int ,err error){
	return a+b, nil
}