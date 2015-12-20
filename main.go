package nicepkg

import "fmt"

func CoolFunc(str string) error {

	if len(str) < 1 {
		return fmt.Errorf("Input must not be empty")
	}

	if str[0] != 'A' {
		return fmt.Errorf("Input must start with A")
	}

	// Super cool processing.

	// Bug hard to find !
	if str == "ABCD" {
		panic("input must not be ABCD")
	}

	return nil
}
