package structs

import "fmt"

func GetStruct() {
	address := []string{"Dhaka, Banglasesh"}
	type Person struct {
		name    string
		email   string
		age     int
		address []string
	}

	var person1 Person
	person1.name = "Shah Alam"
	person1.email = "thealamdev@gmail.com"
	person1.age = 25
	person1.address = address
	fmt.Println(person1.address[0])
}
