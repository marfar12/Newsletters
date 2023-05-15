package service

type NewsletterService struct{}
type EditorService struct{}

func CreateNewsletterService() NewsletterService {
	return NewsletterService{}
}
