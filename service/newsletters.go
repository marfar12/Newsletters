package service

import (
	"context"
	"database/sql"

	dbmodel "newsletter/db/model"
	dbNewsletters "newsletter/db/tables/newsletters"
	"newsletter/service/errors"
	"newsletter/service/model"
)

/*var (
	newsletters = map[string]model.Newsletter{}
)*/

/*func getField(v *model.Newsletter, field string) string{

}*/

// CreateUser saves user in map under email as a key.
func (NewsletterService) CreateNewsletter(_ context.Context, newsletter model.Newsletter, db *sql.DB) (model.Newsletter, error) {
	newNewsletter, err := dbNewsletters.AddNewsletter(db, dbmodel.ToDBNewsletter(newsletter))

	return dbmodel.ToSvcNewsletter(newNewsletter), err
}

// ListUsers returns list of users in array of users.
func (NewsletterService) ListNewsletters(_ context.Context, db *sql.DB) ([]model.Newsletter, error) {
	newsletters, err := dbNewsletters.GetAllNewsletters(db)

	return dbmodel.ToSvcNewsletters(newsletters), err
}

// GetUser returns an user with specified email.
func (NewsletterService) GetNewsletter(_ context.Context, id string, db *sql.DB) (model.Newsletter, error) {
	newsletter, err := dbNewsletters.GetNewsletterById(db, id)

	return dbmodel.ToSvcNewsletter(newsletter), err
}

// UpdateUser updates attributes of a specified user.
func (NewsletterService) UpdateNewsletter(_ context.Context, id string, newsletter model.Newsletter, db *sql.DB) (model.Newsletter, error) {

	return newsletter, nil
}

// DeleteUser deletes user from memory.
func (NewsletterService) DeleteNewsletter(_ context.Context, id string, db *sql.DB) error {
	err := dbNewsletters.DeleteNewsletterById(db, id)
	if err != nil {
		return errors.ErrDeletingNewsletter
	}
	return nil
}
