package arrays

import (
	"fmt"
	"reflect"
)

func SetArray() {
	var country string
	countires := []string{"Bangladesh", "Pakhisthan"}
	countires[0] = "Nepal"
	fmt.Println(countires)
	countires = append(countires, "China")
	for i := 0; i <= len(countires)-1; i++ {
		country = countires[i]
	}
	fmt.Println(country)
}
func GetArray() {
	var person = []int{10}
	fmt.Println("Type of person ", reflect.TypeOf(person).Kind())
	person = append(person, 2)
	fmt.Println(len(person), cap(person))
	location := [5]string{2: "Desi", 3: "Bedesi"}
	fmt.Println(location)
}
