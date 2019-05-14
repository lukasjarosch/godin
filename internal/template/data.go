package template

type Data struct {
	Project Project
	Service Service
}

type Project struct {
	RootPath string
}

type Service struct {
	Name      string
	Namespace string
	Module    string
}
