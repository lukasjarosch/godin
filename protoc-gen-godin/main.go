package main

import (
	"github.com/lukasjarosch/godin/protoc-gen-godin/generate"
	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("./template/data")

	gen := generate.NewGenerator(box)
	gen.Generate()
}
