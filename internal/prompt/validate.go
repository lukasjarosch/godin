package prompt

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/manifoldco/promptui"
	"regexp"
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

func PositiveInteger() promptui.ValidateFunc {
	return func(s string) error {
		val, err := strconv.Atoi(s);
		if err != nil {
			return fmt.Errorf("only positive integers")
		}

		if val < 0 {
			return fmt.Errorf("only positive integers")
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

func CamelCase() promptui.ValidateFunc {
	return func(s string) error {
		ok, _ := regexp.Match(`([A-Z][a-z0-9]+)+`, []byte(s))
		if !ok {
			return fmt.Errorf("string is not CamelCase")
		}
		return nil
	}
}

func ProtoFileExtension() promptui.ValidateFunc {
	return func(s string) error {
		if !strings.HasSuffix(s, ".proto") {
			return fmt.Errorf("must point to a .proto file")
		}
		return nil
	}
}