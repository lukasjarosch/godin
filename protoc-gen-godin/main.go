package main

import (
	"github.com/lukasjarosch/godin/protoc-gen-godin/generate"
)

func main() {
	gen := generate.NewGenerator()
	gen.Generate()
}
