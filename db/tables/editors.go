package tables

import (
	"database/sql"
	dbmodel "newsletter/db/model"
	"newsletter/service/errors"
)

func AddEditor(db *sql.DB, editor dbmodel.Editor) (dbmodel.Editor, error) {
	row := db.QueryRow(`INSERT INTO editors (editor_id, email, password) VALUES($1, $2, $3) RETURNING editor_id, email`, editor.ID, editor.Email, editor.Password)
	err := row.Scan(&editor.ID, &editor.Email)
	if err != nil {
		return dbmodel.Editor{}, errors.ErrCreatingEditor
	}
	return dbmodel.Editor{ID: editor.ID, Email: editor.Email}, nil
}

func GetEditorByEmail(db *sql.DB, editor dbmodel.Editor) (dbmodel.Editor, error) {
	row := db.QueryRow("SELECT * FROM editors WHERE email = $1", editor.Email)
	err := row.Scan(&editor.ID, &editor.Email, &editor.Password)
	if err != nil {
		return dbmodel.Editor{}, errors.ErrNoUserFound
	}
	return dbmodel.Editor{ID: editor.ID, Email: editor.Email, Password: editor.Password}, nil
}
