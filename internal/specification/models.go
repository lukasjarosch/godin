package specification

type Models struct {
	Structs []Structure
	Enums []Enumeration
}

type Structure struct {
	Name string
	Comment []string
	Fields []StructField
}

type StructField struct {
	Name string
	Type string
}

type Enumeration struct {
	Name string
	Comment []string
	Fields []string
}
