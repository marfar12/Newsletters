package model

import (
	svcmodel "newsletter/service/model"

	"github.com/google/uuid"
)

type Subscription struct {
	NewsletterId uuid.UUID `json:"newsletter_id"`
	Email        string    `json:"email" validate:"email"`
}

func ToSvcSubscription(s Subscription) svcmodel.Subscription {
	return svcmodel.Subscription{
		NewsletterId: s.NewsletterId,
		Email:        s.Email,
	}
}

func ToNetSubscription(s svcmodel.Subscription) Subscription {
	return Subscription{
		NewsletterId: s.NewsletterId,
		Email:        s.Email,
	}
}
