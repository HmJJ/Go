package main

import "fmt"

func main() {
	Arr()
	Slice()
	Map()
}

func Arr(){
	fmt.Println("\nArray\n")
	arr1 := [...] int {1:2,3:4}

	for _,value := range arr1{
		fmt.Print(value," ")
	}
	fmt.Println()
}

func Slice(){
	fmt.Println("\nSlice\n")
	//先定义一个数组
	var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}

	//基于数组创建一个数组切片
	var mySlice []int = myArray[1:6]

	fmt.Println("Elements of myArray: ")
	for _,arr := range myArray{
		fmt.Print(arr," ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, arr := range mySlice{
		fmt.Print(arr," ")
	}

	fmt.Println("\nElements of mySlice2: ")
	mySlice2 := make([]int,5,10)
	for _,s := range mySlice2{
		fmt.Print(s," ")
	}
	fmt.Println()
}

type PersonInfo struct {
	ID string
	Name string 
	Address string
}
func Map(){
	fmt.Println("\nMap\n")

	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	//插数据
	personDB["1"] = PersonInfo{"1","nott","莲溪路"}
	personDB["2"] = PersonInfo{"2","jenny","曹杨村"}

	//查数据
	person1, ok := personDB["1"]
	person2, ok := personDB["2"]
	//person2, ok := personDB["3"] 公用一个ok则是‘与’的关系
	if ok {
		fmt.Print(person1.Name," ",person2.Name)
		fmt.Println()
	} else {
		fmt.Println("not found")
	}

	for _,person := range personDB{
		fmt.Print(person.Name," ")
	}
	fmt.Println()

	delete(personDB,"2")

	for _,person := range personDB{
		fmt.Print(person.Name," ")
	}
}