package model

import (
	"context"
	"database/sql"

	svcmodel "newsletter/service/model"
)

type Service interface {
	CreateNewsletter(ctx context.Context, newsletter svcmodel.Newsletter, db *sql.DB) error
	ListNewsletters(ctx context.Context, db *sql.DB) []svcmodel.Newsletter
	GetNewsletter(ctx context.Context, id string, db *sql.DB) (svcmodel.Newsletter, error)
	UpdateNewsletter(ctx context.Context, id string, newsletter svcmodel.Newsletter, db *sql.DB) (svcmodel.Newsletter, error)
	DeleteNewsletter(ctx context.Context, id string, db *sql.DB) error
}
