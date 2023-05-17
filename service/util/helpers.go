package util

import (
	"fmt"
	"newsletter/db/model"
	"os"
	"path/filepath"

	"crypto/rand"
	"math/big"
	svcmodel "newsletter/service/model"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SanitizeUser(editor model.Editor) model.Editor {
	return model.Editor{ID: editor.ID, Email: editor.Email, Password: ""}
}

func NewSubscriptionEmail(email string, newsletterName string, unsubscribe_code string) (*rest.Response, error) {
	var personalizations []*mail.Personalization

	unsubscribeUrl := fmt.Sprintf("%s/unsubscribe/%s", os.Getenv("BASE_URL"), unsubscribe_code)
	subject := "Subscription confirmation"
	to := mail.NewEmail("", email)

	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = subject
	personalization.SetSubstitution("{{.NewsletterName}}", newsletterName)
	personalization.SetSubstitution("{{.UnsubscribeUrl}}", unsubscribeUrl)
	personalizations = append(personalizations, personalization)

	filePath, _ := filepath.Abs("../service/templates/new_subscription.html")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := mail.NewContent("text/html", string(data))

	response, err := SendEmail(newsletterName, content, personalizations)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewIssueEmail(subscriptions []model.Subscription, issue svcmodel.Issue, newsletter svcmodel.Newsletter) (*rest.Response, error) {

	var personalizations []*mail.Personalization

	for _, subscription := range subscriptions {
		unsubscribeUrl := fmt.Sprintf("%s/unsubscribe/%s", os.Getenv("BASE_URL"), subscription.UnsubscribeCode)
		to := mail.NewEmail("", subscription.Email)
		personalization := mail.NewPersonalization()
		personalization.AddTos(to)
		personalization.Subject = fmt.Sprintf("%s - %s", newsletter.Name, issue.Title)
		personalization.SetSubstitution("{{.Title}}", issue.Title)
		personalization.SetSubstitution("{{.Content}}", issue.Content)
		personalization.SetSubstitution("{{.UnsubscribeUrl}}", unsubscribeUrl)
		personalizations = append(personalizations, personalization)
	}

	filePath, _ := filepath.Abs("../service/templates/issue.html")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := mail.NewContent("text/html", string(data))

	response, err := SendEmail(newsletter.Name, content, personalizations)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SendEmail(newsletterName string, content *mail.Content, personalizations []*mail.Personalization) (response *rest.Response, err error) {
	from := mail.NewEmail(newsletterName, os.Getenv("SENDGRID_EMAIL"))

	m := mail.NewV3Mail()

	m.SetFrom(from)
	m.AddContent(content)
	m.AddPersonalizations(personalizations...)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err = sendgrid.API(request)

	return
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
