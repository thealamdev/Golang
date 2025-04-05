package slice

import (
	"fmt"
	"sort"
)

func GetSlice() {
	person := []string{"One", "two", "three", "four", "five", "six", "seven"}
	var remove = 3
	newPerson := append(person[:remove], person[remove+1:]...)
	fmt.Println(newPerson)
	highScroes := make([]int, 4)
	highScroes[0] = 345
	highScroes[1] = 453
	highScroes[2] = 788
	highScroes[3] = 474
	highScroes = append(highScroes, 454, 333, 343)
	sort.Ints(highScroes)
	fmt.Println(highScroes)

}
