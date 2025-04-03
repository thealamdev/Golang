package main

import (
	"apiserver/arrays"
	"apiserver/pointers"
	"apiserver/times"
)

func main() {
	times.GetTime()
	pointers.Pointers()
	arrays.SetArray()
	arrays.GetArray()
	// input := getInput()
	// fmt.Println("Your rating is ", input+1)
}
