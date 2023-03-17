package service

import (
	"context"
	"database/sql"
	"log"

	"newsletter/service/errors"
	"newsletter/service/model"
)

var (
	newsletters = map[string]model.Newsletter{}
)

// CreateUser saves user in map under email as a key.
func (Service) CreateNewsletter(_ context.Context, newsletter model.Newsletter, db *sql.DB) error {
	if _, exists := newsletters[newsletter.ID]; exists {
		return errors.ErrNewsletterAlreadyExists
	}

	newsletters[newsletter.ID] = newsletter

	return nil
}

// ListUsers returns list of users in array of users.
func (Service) ListNewsletters(_ context.Context, db *sql.DB) []model.Newsletter {
	var newsletter model.Newsletter
	var newsletters []model.Newsletter
	rows, errrows := db.Query("SELECT * FROM newsletters")
	if errrows != nil {
		log.Fatal(errrows)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&newsletter.ID, &newsletter.Name, &newsletter.Desc, &newsletter.EditorId)
		if err != nil {
			log.Fatal(err)
		}
		newsletters = append(newsletters, newsletter)
	}
	return newsletters
}

// GetUser returns an user with specified email.
func (Service) GetNewsletter(_ context.Context, id string, db *sql.DB) (model.Newsletter, error) {
	var newsletter model.Newsletter
	row := db.QueryRow("SELECT * FROM newsletters WHERE newsletter_id = $1", id)

	err := row.Scan(&newsletter.ID, &newsletter.Name, &newsletter.Desc, &newsletter.EditorId)
	if err != nil {
		return newsletter, errors.ErrNewsletterDoesntExists
	}
	return newsletter, nil
}

// UpdateUser updates attributes of a specified user.
func (Service) UpdateNewsletter(_ context.Context, id string, newsletter model.Newsletter, db *sql.DB) (model.Newsletter, error) {
	oldUser, exists := newsletters[id]

	if !exists {
		return model.Newsletter{}, errors.ErrNewsletterDoesntExists
	}

	if oldUser.ID == newsletter.ID {
		newsletters[id] = newsletter
	} else {
		newsletters[newsletter.ID] = newsletter

		delete(newsletters, id)
	}

	return newsletter, nil
}

// DeleteUser deletes user from memory.
func (Service) DeleteNewsletter(_ context.Context, id string, db *sql.DB) error {
	if _, exists := newsletters[id]; !exists {
		return errors.ErrNewsletterDoesntExists
	}

	delete(newsletters, id)

	return nil
}
