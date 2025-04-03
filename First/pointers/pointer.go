package pointers

import "fmt"

func Pointers() {
	var modelYear int = 10
	var pointerVar = &modelYear
	*pointerVar = *pointerVar + 2
	fmt.Println(modelYear)
}
