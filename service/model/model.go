package model

type Newsletter struct {
	ID       string
	Name     string
	Desc     string
	EditorId string
}

type Editor struct {
	ID       string
	Email    string
	Password string
}
