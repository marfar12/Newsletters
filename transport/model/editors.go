package model

import (
	svcmodel "newsletter/service/model"
)

type Editor struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password,omitempty"`
}

func ToSvcEditor(u Editor) svcmodel.Editor {
	return svcmodel.Editor{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToNetEditor(u svcmodel.Editor) Editor {
	return Editor{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToNetEditors(editors []svcmodel.Editor) []Editor {
	netEditors := make([]Editor, 0, len(editors))
	for _, editor := range editors {
		netEditors = append(netEditors, ToNetEditor(editor))
	}
	return netEditors
}
