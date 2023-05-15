package util

import (
	"newsletter/db/model"
	"os"

	"crypto/rand"
	"math/big"

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

func SendEmail(subject string, email string, newsletterName string, unsubscribe_code string, templateId string) (response *rest.Response, err error) {
	// to, err := mail.ParseEmail(email)
	// if err != nil {
	// 	return nil, err
	// }
	m := mail.NewV3Mail()
	from := mail.NewEmail("Newsletter", "farm05@vse.cz")
	m.SetFrom(from)
	m.SetTemplateID(templateId)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("", email),
	}
	p.AddTos(tos...)
	p.SetDynamicTemplateData("name", newsletterName)
	p.SetDynamicTemplateData("unsubscribe_code", unsubscribe_code)

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
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
