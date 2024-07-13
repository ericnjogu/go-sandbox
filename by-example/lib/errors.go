package lib

import (
	"errors"
	"fmt"
)

func Add3(arg int) (int, error) {
	if arg == 2 {
		return -1, errors.New("2 is forbidden")
	}
	return arg + 3, nil
}

// custom error
type MyError struct {
	Key  string
	Vars map[string]string
}

// implementing error interface
func (err *MyError) Error() string {
	return fmt.Sprintf("%+v", err)
}
