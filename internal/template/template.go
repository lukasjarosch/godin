package template

import (
	"fmt"
	"go/format"
	"path/filepath"
)

type Template interface {
	Render(tplContext Context) ([]byte, error)
}

// FileExtension defines the template file extension used by Godin
const FileExtension = "tpl"

// BaseTemplate provides the basic templating functionality required for Godin
type BaseTemplate struct {
	rootPath string
	name       string
	isGoSource bool
}

// FormatCode will use go/format to format the given raw code
func (b *BaseTemplate) FormatCode(source []byte) ([]byte, error) {
	formatted, err := format.Source(source)
	if err != nil {
		return nil, err
	}
	return formatted, nil
}

// Filename returns the filename of the currently loaded template
func (b *BaseTemplate) Filename() string {
	return fmt.Sprintf("%s.%s", b.name, FileExtension)
}

// TemplatePath returns the absolute template path to the current base template
// Since packr2: this just returns the filename, the virtual filesystem has a different chroot
func (b *BaseTemplate) TemplatePath() string {
	return b.Filename()
}

// PartialsGlob returns a glob-path which matches all templates inside the partial folder
func (b *BaseTemplate) PartialsGlob() string {
	return filepath.Join("partials", "*.tpl")
}
