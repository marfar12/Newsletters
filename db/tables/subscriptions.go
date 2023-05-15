package tables

import (
	"database/sql"
	dbmodel "newsletter/db/model"
	"newsletter/service/errors"
)

func AddSubscription(db *sql.DB, subscription dbmodel.Subscription) (dbmodel.Subscription, error) {
	row := db.QueryRow(`INSERT INTO subscriptions (newsletter_id, email, unsubscribe_code) VALUES($1, $2, $3) RETURNING *`, subscription.NewsletterId, subscription.Email, subscription.UnsubscribeCode)
	err := row.Scan(&subscription.NewsletterId, &subscription.Email, &subscription.UnsubscribeCode)
	if err != nil {
		return dbmodel.Subscription{}, errors.ErrCreatingSubscription
	}
	return dbmodel.Subscription{NewsletterId: subscription.NewsletterId, Email: subscription.Email, UnsubscribeCode: subscription.UnsubscribeCode}, nil
}

func RemoveSubscription(db *sql.DB, unsubscribe_code string) error {
	res, err := db.Exec(`DELETE FROM subscriptions WHERE unsubscribe_code = $1`, unsubscribe_code)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	return errors.ErrRemovingSubscription

}
