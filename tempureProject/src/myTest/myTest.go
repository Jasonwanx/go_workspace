package main

import (
	"fmt"
)

func reverse_slice(slice []int){
	for i:=0; i < len(slice)/2; i++{
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
}

func reverse_arr(arr *[4]int){
	for i:=0; i < len(arr)/2; i++{
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

func equal(x, y map[string]int) bool{
	if len(x) != len(y){
		return  false
	}

	for k, xv := range x{
		if yv, ok := y[k]; !ok || yv != xv{
			return false
		}
	}

	return true
}


var m = make(map[string]int)

func k(list []string) string {
	str := fmt.Sprintf("%q", list)
	return str
}

func Add(list []string){
	m[k(list)] ++
}

func Count(list []string) int{
	return m[k(list)]
}

func GetNameAndAge(name string, age int) (myName string, myAge int){
	myName = name
	myAge = age
	return myName, myAge
}

func main(){
	s:= []int{1,2, 3, 4}
	reverse_slice(s)
	fmt.Println(s)

	arr := [...]int{1,2,3,4}
	reverse_arr(&arr)
	fmt.Println(arr)

	var ages map[string]int

	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)

	fmt.Println(equal(map[string]int {"B":42}, map[string]int{"B":42}))

	Add([]string{"12", "34", "56"})
	Add([]string{"11", "33", "55"})
	fmt.Println(Count([]string{"12", "34", "56"}))

	name ,age := GetNameAndAge("Json", 29)
	fmt.Printf("%s:%d\n", name, age)
}