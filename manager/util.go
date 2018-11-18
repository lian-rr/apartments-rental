package manager

import "fmt"

func newError(m string, err error) error {
	fmt.Printf("Error: %s", err.Error())
	return fmt.Errorf("%s\n", m)
}