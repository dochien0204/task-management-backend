package repository

import (
	"log"
	"net/smtp"

	"gorm.io/gorm"
)

const (
	From = "doxuanchienh02042002@gmail.com"
    Password = "gmuxctrfivvetdda"

    // SMTP server configuration.
    SmtpHost = "smtp.gmail.com"
    SmtpPort = "587"
)

type EmailRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) *EmailRepository {
	return &EmailRepository{
		db: db,
	}
}

func (r EmailRepository) WithTrx(trxHandle *gorm.DB) EmailRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return r
	}
	r.db = trxHandle
	return r
}

func (r EmailRepository) SendMailPasswordForUser(body string, to []string, subject string) error {

	message := []byte("To: " + to[0] + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" +
        body + "\r\n")

	auth := smtp.PlainAuth("", From, Password, SmtpHost)
	err := smtp.SendMail(SmtpHost+":"+SmtpPort, auth, From, to, message)
	if err != nil {
		return err
	}

	return nil
}

func (r EmailRepository) SendMailForUsers(body string, to []string, subject string) error {

	message := []byte("To: " + to[0] + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" +
        body + "\r\n")

	auth := smtp.PlainAuth("", From, Password, SmtpHost)
	err := smtp.SendMail(SmtpHost+":"+SmtpPort, auth, From, to, message)
	if err != nil {
		return err
	}

	return nil
}