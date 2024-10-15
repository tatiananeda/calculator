package utils

import "errors"

var Operations = map[string]func(float64, float64) (float64, error){
	"+": func(x, y float64) (float64, error) {
		return x + y, nil
	},
	"-": func(x, y float64) (float64, error) {
		return x - y, nil
	},
	"*": func(x, y float64) (float64, error) {
		return x * y, nil
	},
	"/": func(x, y float64) (float64, error) {
		if y == 0 {
			return 0, errors.New("Can't divide by 0")
		}
		return x / y, nil
	},
}
