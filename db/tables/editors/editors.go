package editors

import (
	"database/sql"
	dbmodel "newsletter/db/model"
	"newsletter/service/errors"
)

func AddEditor(db *sql.DB, editor dbmodel.Editor) (dbmodel.Editor, error) {
	row := db.QueryRow(`INSERT INTO editors (email, password) VALUES($1, $2) RETURNING editor_id, email`, editor.Email, editor.Password)
	err := row.Scan(&editor.ID, &editor.Email)
	if err != nil {
		return dbmodel.Editor{}, errors.ErrCreatingEditor
	}
	return dbmodel.Editor{ID: editor.ID, Email: editor.Email}, nil
}

func GetEditorByEmailAndPassword(db *sql.DB, editor dbmodel.Editor) (dbmodel.Editor, error) {
	row := db.QueryRow("SELECT editor_id, email FROM editors WHERE email = $1 AND password = $2", editor.Email, editor.Password)
	err := row.Scan(&editor.ID, &editor.Email)
	if err != nil {
		return dbmodel.Editor{ID: editor.ID, Email: editor.Email}, errors.ErrRetrievingNewsletter
	}
	return dbmodel.Editor{ID: editor.ID, Email: editor.Email}, nil
}
