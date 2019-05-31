// Code generated by Godin v0.3.0; DO NOT EDIT.

package endpoint

type (
	HelloRequest struct {
		Name string `json:"name"`
	}

	HelloResponse struct {
		Greeting *Greeting `json:"greeting"`
		Err      error     `json:"-"`
	}

	Hello2Request struct {
		Name string `json:"name"`
	}

	Hello2Response struct {
		Greeting Greeting `json:"greeting"`
		Err      error    `json:"-"`
	}

	Hello3Request struct {
		Name string `json:"name"`
	}

	Hello3Response struct {
		Greeting string `json:"greeting"`
		Err      error  `json:"-"`
	}

	Hello4Request struct {
		Name []string `json:"name"`
	}

	Hello4Response struct {
		Greeting []Greeting `json:"greeting"`
		Err      error      `json:"-"`
	}

	Hello5Request struct {
		Name []string `json:"name"`
	}

	Hello5Response struct {
		Greeting []*Greeting `json:"greeting"`
		Err      error       `json:"-"`
	}

	Hello6Request struct {
		Name []*Greeting `json:"name"`
	}

	Hello6Response struct {
		Greeting []*Greeting `json:"greeting"`
		Err      error       `json:"-"`
	}

	Hello7Request struct {
		Name *Greeting `json:"name"`
	}

	Hello7Response struct {
		Greeting []string `json:"greeting"`
		Err      error    `json:"-"`
	}

	Hello8Request struct {
		Name *[]Greeting `json:"name"`
	}

	Hello8Response struct {
		Greeting []string `json:"greeting"`
		Err      error    `json:"-"`
	}

	Hello9Request struct {
		Name *[]Greeting `json:"name"`
		Foo  string      `json:"foo"`
		Bar  string      `json:"bar"`
	}

	Hello9Response struct {
		Greeting []string `json:"greeting"`
		Err      error    `json:"-"`
	}
)

// Implement the Failer interface for all responses
func (resp HelloResponse) Failed() error  { return r.Err }
func (resp Hello2Response) Failed() error { return r.Err }
func (resp Hello3Response) Failed() error { return r.Err }
func (resp Hello4Response) Failed() error { return r.Err }
func (resp Hello5Response) Failed() error { return r.Err }
func (resp Hello6Response) Failed() error { return r.Err }
func (resp Hello7Response) Failed() error { return r.Err }
func (resp Hello8Response) Failed() error { return r.Err }
func (resp Hello9Response) Failed() error { return r.Err }
