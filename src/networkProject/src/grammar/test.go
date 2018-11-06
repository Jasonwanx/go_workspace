package main

import (
	"fmt"
	"sort"
)

type Shaper interface {
	Area() float32
}

type Square struct{
	side float32
}

func (sq *Square) Area() float32{
	fmt.Println("Square Aread func")
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	fmt.Println("Rectangle Aread func")
	return r.length * r.width
}

type IntArray []int

func (p IntArray) Len() int{
	return len(p)
}

func (p IntArray) Less(i,j int) bool{
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int)  {
	p[i],p[j] = p[j],p[i]
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday    := day{0, "SUN", "Sunday"}
	Monday    := day{1, "MON", "Monday"}
	Tuesday   := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday  := day{4, "THU", "Thursday"}
	Friday    := day{5, "FRI", "Friday"}
	Saturday  := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

func main() {
	days()
}
