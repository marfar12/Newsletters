package model

import (
	"database/sql"
	svcmodel "newsletter/service/model"
)

type Newsletter struct {
	ID       string
	Name     string
	Desc     sql.NullString
	EditorId string
}

func ToSvcNewsletter(u Newsletter) svcmodel.Newsletter {
	return svcmodel.Newsletter{
		ID:       u.ID,
		Name:     u.Name,
		Desc:     u.Desc.String,
		EditorId: u.EditorId,
	}
}

func ToDBNewsletter(u svcmodel.Newsletter) Newsletter {
	return Newsletter{
		ID:       u.ID,
		Name:     u.Name,
		Desc:     sql.NullString{String: u.Desc, Valid: true},
		EditorId: u.EditorId,
	}
}

func ToSvcNewsletters(newsletters []Newsletter) []svcmodel.Newsletter {
	svcNewsletters := make([]svcmodel.Newsletter, 0, len(newsletters))
	for _, newsletter := range newsletters {
		svcNewsletters = append(svcNewsletters, ToSvcNewsletter(newsletter))
	}
	return svcNewsletters
}
