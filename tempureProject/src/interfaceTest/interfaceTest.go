package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Worker struct{
	empId int
	basicpay int
}

type Contract struct{
	Worker
}

type Permanent struct {
	Worker
	jj int
}

func (c Contract) CalculateSalary() int{
	return c.basicpay
}

func (p Permanent) CalculateSalary() int{
	return p.basicpay + p.jj
}

func totalExpense(s []SalaryCalculator) int{
	expense := 0
	for _, v := range s{
		expense += v.CalculateSalary()
	}
	return expense
}

func describe(i interface{}){
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}){
	v, ok := i.(int)

	fmt.Println(v, ok)
}

func main(){
	//permp1 := Permanent{Worker{1, 3000}, 10000}
	//permp2 := Permanent{Worker{2, 4000}, 15000}
	//
	//cemp1 := Contract{Worker{3, 4000}}
	//cemp2 := Contract{Worker{4,5000}}
	//
	//employees := []SalaryCalculator{permp1,permp2,cemp1,cemp2}
	//
	//fmt.Printf("total:%d\n", totalExpense(employees))
	//
	//s := "Hello "
	//i:= 55
	//strt := struct{
	//	name string
	//}{"Jaos"}
	//
	//describe(s)
	//describe(i)
	//describe(strt)

	var s interface{} = 55
	assert(s)
}