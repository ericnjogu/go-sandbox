package lib

import (
	"fmt"
	"runtime"
)

// range
func OddNumbers(to int) []int {
	var odd []int
	for n := range to + 1 {
		if n%2 != 0 {
			odd = append(odd, n)
		}
	}
	return odd
}

// range, strings
func Reverse(str string) string {
	n := len(str)
	reverse := ""
	for n > 0 {
		n -= 1
		reverse += string(str[n])
	}
	return reverse
}

// switch
func Runtime() string {
	switch runtime.GOOS {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		return runtime.GOOS
	}
}

func GetType(i interface{}) string {
	switch t := i.(type) {
	case bool:
		return "boolean"
	case string:
		return "string"
	default:
		return fmt.Sprintf("%T", t)
	}
}

// variadic args
func Mean(nums ...float64) float64 {
	var total float64 = 0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

// closures
func Increment(x int) func() int {
	store := x
	return func() int {
		store++
		return store
	}
}

// recursion
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
