package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Employee struct{
	ID 			int
	Name   		string
	Address 	string
	DoB 		time.Time
	Position 	string
	Salary 		int
	ManagerId 	int
}

var dilbert Employee

func EmployeeByID(id int) *Employee{
	return &dilbert
}

type Point struct{
	X,Y int
}

type Circle struct{
	Point
	Radius int
}

type Wheel struct{
	Circle
	Spoke int
}

type Movie struct{
	Title 	string
	Year 	int  `json:"released"`
	Color 	bool `json:"color, omitempty"`
	Actors	[]string
}

var movies = []Movie{
	{Title:"Casablanca", Year:1942,  Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title:"Cool Hand Luke", Year:1967, Actors:[]string{"Paul Newman"}},
}

func main(){
	//w := Wheel{
	//	Circle{
	//		Point{8,8},
	//		5,
	//	},
	//20,
	//}
	//
	//fmt.Printf("%#v\n", w)

	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil{
		log.Fatal("Json marshaling failed:%s", err)
	}

	fmt.Printf("%s\n", data)

	var moviess []Movie
	if err := json.Unmarshal(data, &moviess); err != nil {
		log.Fatal("Json Unmarshal failed:%s", err)
	}

	fmt.Println(moviess)
}
