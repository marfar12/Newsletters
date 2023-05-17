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

	SignIn(ctx context.Context, editor model.Editor, db *sql.DB) (model.Editor, error)
	SignUp(ctx context.Context, editor model.Editor, db *sql.DB) (model.Editor, error)

	Subscribe(ctx context.Context, subscription model.Subscription, db *sql.DB) (model.Subscription, error)
	Unsubscribe(ctx context.Context, unsubscribe_code string, db *sql.DB) error

	Publish(ctx context.Context, issue model.Issue, db *sql.DB) error
}
