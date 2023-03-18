package newsletters

import (
	"database/sql"
	dbmodel "newsletter/db/model"
	"newsletter/service/errors"
)

func AddNewsletter(db *sql.DB, newsletter dbmodel.Newsletter) (dbmodel.Newsletter, error) {
	row := db.QueryRow(`INSERT INTO newsletters (name, "desc", editor_id) VALUES($1, $2, $3) RETURNING *`, newsletter.Name, newsletter.Desc, newsletter.EditorId)
	err := row.Scan(&newsletter.ID, &newsletter.Name, &newsletter.Desc, &newsletter.EditorId)
	if err != nil {
		return dbmodel.Newsletter{}, errors.ErrCreatingNewsletter
	}
	return newsletter, nil
}

func GetAllNewsletters(db *sql.DB) ([]dbmodel.Newsletter, error) {
	var newsletter dbmodel.Newsletter
	var newsletters []dbmodel.Newsletter
	rows, errQuery := db.Query("SELECT * FROM newsletters")
	if errQuery != nil {
		return newsletters, errors.ErrRetrievingNewsletters
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&newsletter.ID, &newsletter.Name, &newsletter.Desc, &newsletter.EditorId)
		if err != nil {
			return newsletters, errors.ErrRetrievingNewsletters
		}
		newsletters = append(newsletters, newsletter)
	}
	return newsletters, nil
}

func GetNewsletterById(db *sql.DB, id string) (dbmodel.Newsletter, error) {
	var newsletter dbmodel.Newsletter
	row := db.QueryRow("SELECT * FROM newsletters WHERE newsletter_id = $1", id)
	err := row.Scan(&newsletter.ID, &newsletter.Name, &newsletter.Desc, &newsletter.EditorId)
	if err != nil {
		return newsletter, errors.ErrRetrievingNewsletter
	}
	return newsletter, nil
}

func GetNewslettersByEditorId(db *sql.DB, editorId string) ([]dbmodel.Newsletter, error) {
	var newsletter dbmodel.Newsletter
	var newsletters []dbmodel.Newsletter
	rows, errQuery := db.Query("SELECT * FROM newsletters WHERE editor_id = $1", editorId)
	if errQuery != nil {
		return newsletters, errors.ErrRetrievingNewsletters
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&newsletter.ID, &newsletter.Name, &newsletter.Desc, &newsletter.EditorId)
		if err != nil {
			return newsletters, errors.ErrRetrievingNewsletters
		}
		newsletters = append(newsletters, newsletter)
	}
	return newsletters, nil
}

func DeleteNewsletterById(db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM newsletters WHERE newsletter_id = $1", id)
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count > 0 {
				return nil
			}
		}

	}
	return errors.ErrDeletingNewsletter
}
