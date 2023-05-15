package errors

import "errors"

var (
	ErrCreatingNewsletter      = errors.New("newsletter could not be created")
	ErrRetrievingNewsletter    = errors.New("newsletter could not be retrieved")
	ErrRetrievingNewsletters   = errors.New("newsletters could not be retrieved")
	ErrUldatingNewsletter      = errors.New("newsletter could not be updated")
	ErrUpdatingNewsletter      = errors.New("newsletter could not be updated")
	ErrDeletingNewsletter      = errors.New("newsletter could not be deleted")
	ErrCreatingEditor          = errors.New("editor could not be created")
	ErrRetrievingEditors       = errors.New("editors could not be retrieved")
	ErrLoggingIn               = errors.New("email or password is incorrect")
	ErrHashingPassword         = errors.New("password could not be hashed")
	ErrCratingEditor           = errors.New("editor could not be created")
	ErrNoUserFound             = errors.New("no user found")
	ErrUnexpectedSigningMethot = errors.New("unexpected jwt signing method")
	ErrInvalidJWTToken         = errors.New("invalid jwt token")
	ErrCreatingSubscription    = errors.New("subscription could not be created")
	ErrRemovingSubscription    = errors.New("subscription could not be removed")
)
