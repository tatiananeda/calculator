package main

import (
	"fmt"
	"os"

	"calculator/utils" //precede with package go.mod name
)

func main() {
	var x, y float64
	var o string

	x = utils.CallFuncExitOnError(utils.RequestInput, "Enter a number \n", x)
	y = utils.CallFuncExitOnError(utils.RequestInput, "Enter a number \n", x)
	o = utils.CallFuncExitOnError(utils.RequestInput, "Enter operator +, -, * or / \n", o)

	if utils.Operations[o] == nil {
		fmt.Printf("operation %s is not supported", o)
		os.Exit(1)
	}

	res, err := utils.Operations[o](x, y)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result: %f", res)
}
