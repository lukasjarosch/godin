package template

type Data struct {
	Project Project
	Service Service
	Godin Godin
}

type Project struct {
	RootPath string
}

type Service struct {
	Name      string
	Namespace string
	Module    string
}

type Godin struct {
	Version string
	Commit string
	Build string
}
