package main

import "fmt"

type data struct {
	Name string
	age  uint32
	grade map[string]uint32
}


func main()  {
	m := make(map[int]*data, 4)
	d1 := &data{
		Name: "wtq",
		age:  12,
	}
	m[0] = d1
	m[0].grade = make(map[string]uint32 ,10)
	m[0].grade["chinese"] = 100
	for i, d := range m {
		fmt.Printf("index:%d  %v \n",i, d)
	}

}
