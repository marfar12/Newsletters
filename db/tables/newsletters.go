package tables

import (
	"database/sql"
	dbmodel "newsletter/db/model"
	"newsletter/service/errors"
)

func AddNewsletter(db *sql.DB, newsletter dbmodel.Newsletter) (dbmodel.Newsletter, error) {
	row := db.QueryRow(`INSERT INTO newsletters (newsletter_id, name, "desc", editor_id) VALUES($1, $2, $3, $4) RETURNING *`, newsletter.ID, newsletter.Name, newsletter.Desc, newsletter.EditorId)
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

func UpdateNewsletterById(editorId string, db *sql.DB, id string, newsletter dbmodel.Newsletter) (dbmodel.Newsletter, error) {
	var resNewsletter dbmodel.Newsletter
	row := db.QueryRow("UPDATE newsletters SET name = $1 WHERE newsletter_id = $2 AND editor_id = $3 RETURNING *", newsletter.Name, id, editorId)
	err := row.Scan(&resNewsletter.ID, &resNewsletter.Name, &resNewsletter.Desc, &resNewsletter.EditorId)
	if err != nil {
		return newsletter, errors.ErrUpdatingNewsletter
	}
	return resNewsletter, nil
}

func DeleteNewsletterById(editorId string, db *sql.DB, id string) error {
	res, err := db.Exec("DELETE FROM newsletters WHERE newsletter_id = $1 AND editor_id = $2", id, editorId)
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
