package maps

import "fmt"

func GetMap() {
	collection := make(map[string]string)
	collection["js"] = "Javascript"
	collection["ja"] = "Java"
	collection["py"] = "Python"
	collection["go"] = "Go Lang"
	for key, value := range collection {
		fmt.Printf("Collection for key %v is %v \n", key, value)
	}
	delete(collection, "py")
	fmt.Println(collection)
}
