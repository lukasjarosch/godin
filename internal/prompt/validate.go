package prompt

import (
	"fmt"
	"strings"
	"github.com/manifoldco/promptui"
)

func Validate(validators ...promptui.ValidateFunc) promptui.ValidateFunc {
	return func(s string) error {
		for _, validator := range validators {
			if err := validator(s); err != nil {
				return err
			}
		}
		return nil
	}
}

func MinLengthThree() promptui.ValidateFunc {
	return func(s string) error {
		if len(s) < 3 {
			return fmt.Errorf("must be at least 3 characters")
		}
		return nil
	}
}

func Lowercase() promptui.ValidateFunc {
	return func(s string) error {
		if strings.ToLower(s) != s {
			return fmt.Errorf("lowercase only")
		}
		return nil
	}
}

func GoSuffix() promptui.ValidateFunc {
	return func(s string) error {
		if !strings.HasSuffix(s, ".go") {
			return fmt.Errorf("with .go file-extension")
		}
		return nil
	}
}

