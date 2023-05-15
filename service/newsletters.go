package service

import (
	"context"
	"database/sql"

	dbmodel "newsletter/db/model"
	dbtables "newsletter/db/tables"
	"newsletter/service/errors"
	"newsletter/service/model"
	"newsletter/service/util"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

func (NewsletterService) CreateNewsletter(c context.Context, newsletter model.Newsletter, db *sql.DB) (model.Newsletter, error) {
	_, claims, _ := jwtauth.FromContext(c)
	newsletter.ID = uuid.New()
	newsletter.EditorId, _ = uuid.Parse(claims["ID"].(string))
	newNewsletter, err := dbtables.AddNewsletter(db, dbmodel.ToDBNewsletter(newsletter))

	return dbmodel.ToSvcNewsletter(newNewsletter), err
}

func (NewsletterService) ListNewsletters(_ context.Context, db *sql.DB) ([]model.Newsletter, error) {
	newsletters, err := dbtables.GetAllNewsletters(db)

	return dbmodel.ToSvcNewsletters(newsletters), err
}

func (NewsletterService) GetNewsletter(_ context.Context, id string, db *sql.DB) (model.Newsletter, error) {
	newsletter, err := dbtables.GetNewsletterById(db, id)

	return dbmodel.ToSvcNewsletter(newsletter), err
}

func (NewsletterService) UpdateNewsletter(c context.Context, id string, newsletter model.Newsletter, db *sql.DB) (model.Newsletter, error) {
	_, claims, _ := jwtauth.FromContext(c)
	resNewsletter, err := dbtables.UpdateNewsletterById(claims["ID"].(string), db, id, dbmodel.ToDBNewsletter(newsletter))
	if err != nil {
		return model.Newsletter{}, errors.ErrUpdatingNewsletter
	}

	return dbmodel.ToSvcNewsletter(resNewsletter), nil
}

func (NewsletterService) DeleteNewsletter(c context.Context, id string, db *sql.DB) error {
	_, claims, _ := jwtauth.FromContext(c)
	err := dbtables.DeleteNewsletterById(claims["ID"].(string), db, id)
	if err != nil {
		return errors.ErrDeletingNewsletter
	}
	return nil
}

func (NewsletterService) SignUp(_ context.Context, editor model.Editor, db *sql.DB) (model.Editor, error) {
	hashedPassword, err := util.HashPassword(editor.Password)
	if err != nil {
		return model.Editor{}, errors.ErrHashingPassword
	}
	editor.ID = uuid.New()
	editor.Password = hashedPassword
	newEditor, err := dbtables.AddEditor(db, dbmodel.ToDBEditor(editor))

	if err != nil {
		return model.Editor{}, errors.ErrCratingEditor
	}

	return dbmodel.ToSvcEditor(newEditor), err
}

func (NewsletterService) SignIn(_ context.Context, editor model.Editor, db *sql.DB) (model.Editor, error) {
	dbEditor, err := dbtables.GetEditorByEmail(db, dbmodel.ToDBEditor(editor))

	if err != nil {
		return model.Editor{}, errors.ErrLoggingIn
	}

	match := util.CheckPasswordHash(editor.Password, dbEditor.Password)
	if !match {
		return model.Editor{}, errors.ErrLoggingIn
	}

	dbEditor = util.SanitizeUser(dbEditor)
	return dbmodel.ToSvcEditor(dbEditor), nil
}

func (NewsletterService) Subscribe(c context.Context, subscription model.Subscription, db *sql.DB) (model.Subscription, error) {
	var templateId string = "d-c5b9fdcbbf0b4ca6a229a9c5204df7c9"
	unsubscribeCode, _ := util.GenerateRandomString(32)

	subscription.UnsubscribeCode = unsubscribeCode

	newSubscription, err := dbtables.AddSubscription(db, dbmodel.ToDBSubscription(subscription))

	if err != nil {
		return model.Subscription{}, errors.ErrCreatingSubscription
	}

	newsletter, err := dbtables.GetNewsletterById(db, subscription.NewsletterId.String())

	if err != nil {
		return model.Subscription{}, errors.ErrRetrievingNewsletter
	}

	util.SendEmail("Subscription confirmation", subscription.Email, newsletter.Name, unsubscribeCode, templateId)

	return dbmodel.ToSvcSubscription(newSubscription), nil
}

func (NewsletterService) Unsubscribe(c context.Context, unsubscribe_code string, db *sql.DB) error {
	err := dbtables.RemoveSubscription(db, unsubscribe_code)

	if err != nil {
		return errors.ErrRemovingSubscription
	}

	//util.SendEmail("Subscription confirmation", subscription.Email, newsletter.Name, unsubscribeCode, templateId)

	return nil
}
