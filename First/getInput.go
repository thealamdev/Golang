package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() float64 {
	fmt.Println("Enter your rating...")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fianlInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		return fianlInput
	}
	return fianlInput
}
