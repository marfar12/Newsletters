package errors

import "errors"

var (
	ErrCreatingNewsletter    = errors.New("newsletter could not be created")
	ErrRetrievingNewsletter  = errors.New("newsletter could not be retrieved")
	ErrRetrievingNewsletters = errors.New("newsletters could not be retrieved")
	ErrUldatingNewsletter    = errors.New("newsletter could not be updated")
	ErrDeletingNewsletter    = errors.New("newsletter could not be deleted")
	ErrCreatingEditor        = errors.New("editor could not be created")
	ErrRetrievingEditors     = errors.New("editors could not be retrieved")
)
