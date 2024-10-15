package utils

import (
	"fmt"
	"os"
)

type InputType interface {
	float64 | string
}

func RequestInput[T InputType](q string, t T) (T, error) {
	var input T
	fmt.Print(q)
	switch any(t).(type) {
	case float64:
		_, err := fmt.Scanf("%f", &input)
		return T(input), err
	default:
		_, err := fmt.Scanln(&input)
		return T(input), err
	}
}

func CallFuncExitOnError[T InputType](f func(string, T) (T, error), message string, t T) T {
	val, err := f(message, t)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return val
}
