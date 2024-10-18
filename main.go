package main

import (
	"fmt"
	"os"

	u "github.com/tatiananeda/calculator/utils" //precede with package go.mod name
)

func main() {
	var x, y float64
	var o string

	x = u.CallFuncExitOnError(u.RequestInput, "Enter a number \n", x)
	y = u.CallFuncExitOnError(u.RequestInput, "Enter a number \n", x)
	o = u.CallFuncExitOnError(u.RequestInput, "Enter operator +, -, * or / \n", o)

	if u.Operations[o] == nil {
		fmt.Printf("operation %s is not supported", o)
		os.Exit(1)
	}

	res, err := u.Operations[o](x, y)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result: %f", res)
}
