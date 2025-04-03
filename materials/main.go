package main

import (
	"apiserver/arrays"
	"apiserver/pointers"
	"apiserver/structs"
	"apiserver/times"
)

func main() {
	structs.GetStruct()
	times.GetTime()
	pointers.Pointers()
	arrays.SetArray()
	arrays.GetArray()
	// input := getInput()
	// fmt.Println("Your rating is ", input+1)
}
