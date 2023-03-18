package service

import (
	"context"
	"database/sql"

	dbmodel "newsletter/db/model"
	dbEditors "newsletter/db/tables/editors"
	"newsletter/service/model"
)

// CreateUser saves user in map under email as a key.
func (EditorService) SignIn(_ context.Context, editor model.Editor, db *sql.DB) (model.Editor, error) {
	signedInEditor, err := dbEditors.GetEditorByEmailAndPassword(db, dbmodel.ToDBEditor(editor))

	return dbmodel.ToSvcEditor(signedInEditor), err
}

// ListUsers returns list of users in array of users.
func (EditorService) SignUp(_ context.Context, editor model.Editor, db *sql.DB) (model.Editor, error) {
	newEditor, err := dbEditors.AddEditor(db, dbmodel.ToDBEditor(editor))

	return dbmodel.ToSvcEditor(newEditor), err
}
