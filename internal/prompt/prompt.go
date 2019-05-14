package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type Prompt struct {
	prompt promptui.Prompt
}

func NewPrompt(label, defaultValue string, validateFunc promptui.ValidateFunc) *Prompt {
	prompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s ", label),
		Default: defaultValue,
		Validate:validateFunc,
	}

	return &Prompt{
		prompt: prompt,
	}
}

func (p *Prompt) Run() (string, error) {
	return p.prompt.Run()
}
