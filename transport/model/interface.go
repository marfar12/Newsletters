package model

import (
	"context"
	"database/sql"

	"newsletter/service/model"
)

type NewsletterService interface {
	CreateNewsletter(ctx context.Context, newsletter model.Newsletter, db *sql.DB) (model.Newsletter, error)
	ListNewsletters(ctx context.Context, db *sql.DB) ([]model.Newsletter, error)
	GetNewsletter(ctx context.Context, id string, db *sql.DB) (model.Newsletter, error)
	UpdateNewsletter(ctx context.Context, id string, newsletter model.Newsletter, db *sql.DB) (model.Newsletter, error)
	DeleteNewsletter(ctx context.Context, id string, db *sql.DB) error
}

type EditorService interface {
	SignIn(ctx context.Context, editor model.Editor, db *sql.DB) (model.Editor, error)
	SignUp(ctx context.Context, editor model.Editor, db *sql.DB) (model.Editor, error)
}
