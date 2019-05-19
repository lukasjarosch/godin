package generate

import (
	. "github.com/dave/jennifer/jen"
	"github.com/lukasjarosch/godin/internal"
	"fmt"
)

type Endpoint struct {
	data *internal.Data
}

func NewEndpoint(data *internal.Data) *Endpoint {
    return &Endpoint{
    	data:data,
	}
}

func (e *Endpoint) Render() {
	f := NewFile("endpoint")
	f.HeaderComment("Code generated by Godin v" + e.data.Godin.Version + ". DO NOT EDIT.")

	fmt.Println(f.GoString())
}