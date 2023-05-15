package model

import (
	svcmodel "newsletter/service/model"

	"github.com/google/uuid"
)

type Subscription struct {
	NewsletterId    uuid.UUID
	Email           string
	UnsubscribeCode string
}

func ToSvcSubscription(s Subscription) svcmodel.Subscription {
	return svcmodel.Subscription{
		NewsletterId: s.NewsletterId,
		Email:        s.Email,
	}
}

func ToDBSubscription(s svcmodel.Subscription) Subscription {
	return Subscription{
		NewsletterId:    s.NewsletterId,
		Email:           s.Email,
		UnsubscribeCode: s.UnsubscribeCode,
	}
}
