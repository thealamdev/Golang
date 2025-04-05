package main

import (
	"apiserver/arrays"
	"apiserver/maps"
	"apiserver/pointers"
	"apiserver/slice"
	"apiserver/structs"
	"apiserver/times"
)

func main() {
	maps.GetMap()
	slice.GetSlice()
	structs.GetStruct()
	times.GetTime()
	pointers.Pointers()
	arrays.SetArray()
	arrays.GetArray()
	// input := getInput()
	// fmt.Println("Your rating is ", input+1)
}
