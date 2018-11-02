package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func appendInt(x []int, y int) []int{
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x){
		z = x[:zlen]
	}else{
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}

	z[len(x)] = y
	return  z
}

func main(){

	defer fmt.Printf("end main;")

	var x []int
	for i:=0; i<10; i++{
		x = appendInt(x, i)
	}

	for _,v := range x{
		fmt.Printf("%d\n", v)
	}

	for _,arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil{
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

	var a = [...]int{1,2,3,4,5}

	for i,v := range a{
		fmt.Printf("%d : %d\n", i, v)
	}
}