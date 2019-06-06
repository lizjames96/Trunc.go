package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Enter a floating point number")
	var inp string
	fmt.Scan(&inp)
	parts := strings.Split(inp, ".")
	if len(parts) != 2 {
		fmt.Println("Invalid input!")
		return
	}
	_, err1 := strconv.Atoi(parts[0])
	_, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
    	fmt.Println("Invalid input!")
    	return
	}
	fmt.Printf("Truncated number is %s. \n", parts[0])
}