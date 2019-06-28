package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type Select struct {
	prompt promptui.Select
}

func NewSelect(label string, selectValues []string) (string, error) {
	prompt := promptui.Select{
		Label:    fmt.Sprintf("%s ", label),
		Items: selectValues,
	}

	s :=  &Select{
		prompt: prompt,
	}

	_, result, err :=  s.prompt.Run()

	return result, err
}

