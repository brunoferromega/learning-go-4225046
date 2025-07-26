package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	iamNum := 23

	fmt.Println(str1, str2, str3)
	strSize, err := fmt.Println("I'm the number", iamNum)
	if err == nil {
		fmt.Printf("The string size is %v\n", strSize)
	}

	fmt.Printf("Checking a var type %T\n", iamNum)
}
