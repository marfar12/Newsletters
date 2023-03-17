package model

import "database/sql"

type Newsletter struct {
	ID       string
	Name     string
	Desc     sql.NullString
	EditorId string
}
