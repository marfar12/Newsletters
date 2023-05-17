package model

import (
	svcmodel "newsletter/service/model"

	"github.com/google/uuid"
)

type Issue struct {
	NewsletterId uuid.UUID `json:"newsletter_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
}

func ToSvcIssue(i Issue) svcmodel.Issue {
	return svcmodel.Issue{
		NewsletterId: i.NewsletterId,
		Title:        i.Title,
		Content:      i.Content,
	}
}

func ToNetIssue(i svcmodel.Issue) Issue {
	return Issue{
		NewsletterId: i.NewsletterId,
		Title:        i.Title,
		Content:      i.Content,
	}
}
