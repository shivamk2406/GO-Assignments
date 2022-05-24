package mail

import (
	"fmt"
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/mail"
)

type MailService interface {
	CreateNewsLetter(news news.SingleNews) string
	CreateMail(To string, From string, subject string, newsCollection []news.SingleNews) mail.NewMail
	SendMail(m mail.NewMail) error
}

type mailService struct {
	mailsvc *mail.Mailer
}

func NewMailService(dialer *mail.Mailer) MailService {
	return mailService{mailsvc: dialer}
}

func (m mailService) CreateMail(to string, from string, subject string, newsCollection []news.SingleNews) mail.NewMail {

	var newsContent string
	for _, val := range newsCollection {
		newsContent = newsContent + m.CreateNewsLetter(val)
	}
	return mail.NewMail{To: to,
		From:        from,
		Subject:     subject,
		Description: newsContent}
}

func (m mailService) SendMail(mail mail.NewMail) error {
	err := m.mailsvc.SendMail(mail)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m mailService) CreateNewsLetter(news news.SingleNews) string {
	newsContent := fmt.Sprintf("%s\n%s\n-----------------------------------\n", news.Heading, news.Description)
	fmt.Println(newsContent)

	return newsContent
}
