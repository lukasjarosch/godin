package generate

import (
	"github.com/lukasjarosch/godin/internal/template"
)

type Option func(*Options)

type Options struct {
	Template   string
	IsGoSource bool
	Context    template.Context
	TargetFile string
	Overwrite  bool
}

func Template(template string) Option {
	return func(options *Options) {
		options.Template = template
	}
}

func IsGoSource(isGoSource bool) Option {
	return func(options *Options) {
		options.IsGoSource = isGoSource
	}
}

func Context(ctx template.Context) Option {
	return func(options *Options) {
		options.Context = ctx
	}
}

func TargetFile(file string) Option {
	return func(options *Options) {
		options.TargetFile = file
	}
}

func Overwrite(overwrite bool) Option {
	return func(options *Options) {
		options.Overwrite = overwrite
	}
}