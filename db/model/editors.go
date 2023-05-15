package model

import (
	svcmodel "newsletter/service/model"

	"github.com/google/uuid"
)

type Editor struct {
	ID       uuid.UUID
	Email    string
	Password string
}

func ToSvcEditor(u Editor) svcmodel.Editor {
	return svcmodel.Editor{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToDBEditor(u svcmodel.Editor) Editor {
	return Editor{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToSvcEditors(editors []Editor) []svcmodel.Editor {
	svcEditors := make([]svcmodel.Editor, 0, len(editors))
	for _, editor := range editors {
		svcEditors = append(svcEditors, ToSvcEditor(editor))
	}
	return svcEditors
}
