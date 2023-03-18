package model

import (
	svcmodel "newsletter/service/model"
)

type Newsletter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	EditorId string `json:"editor_id"`
}

func ToSvcNewsletter(u Newsletter) svcmodel.Newsletter {
	return svcmodel.Newsletter{
		ID:       u.ID,
		Name:     u.Name,
		Desc:     u.Desc,
		EditorId: u.EditorId,
	}
}

func ToNetNewsletter(u svcmodel.Newsletter) Newsletter {
	return Newsletter{
		ID:       u.ID,
		Name:     u.Name,
		Desc:     u.Desc,
		EditorId: u.EditorId,
	}
}

func ToNetNewsletters(newsletters []svcmodel.Newsletter) []Newsletter {
	netNewsletters := make([]Newsletter, 0, len(newsletters))
	for _, newsletter := range newsletters {
		netNewsletters = append(netNewsletters, ToNetNewsletter(newsletter))
	}
	return netNewsletters
}
