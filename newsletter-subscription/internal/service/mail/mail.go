package mail

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	gomail "gopkg.in/mail.v2"
)

type MailServiceConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type MailService interface {
	CreateNewsLetter(news news.SingleNews) string
	SendMail()
}

type mail struct {
	email string
	news  []news.SingleNews
}

func NewMailService(email string, news []news.SingleNews) MailService {
	return mail{email: email, news: news}
}

func (m mail) SendMail() {
	conf, err := config.LoadMailService()
	if err != nil {
		log.Fatal(err)
	}

	msg := gomail.NewMessage()

	msg.SetHeader("From", "shivam.1si18cs104@gmail.com")

	msg.SetHeader("To", m.email)

	msg.SetHeader("Subject", "Your Daily Feed")

	var newsContent string
	for _, val := range m.news {
		newsContent = newsContent + m.CreateNewsLetter(val)
	}

	msg.SetBody("text/plain", newsContent)

	d := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(msg); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}

func (m mail) CreateNewsLetter(news news.SingleNews) string {
	newsContent := fmt.Sprintf("%s\n%s\n-----------------------------------\n", news.Heading, news.Description)
	fmt.Println(newsContent)

	return newsContent
}
