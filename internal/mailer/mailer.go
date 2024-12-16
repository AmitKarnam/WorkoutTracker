package mailer

import (
	"log"
	"net/smtp"
)

func mailer() {
	// Set up email details
	from := "jd6219996@gmail.com"
	to := "amitkarnam6@gmail.com"
	subject := "Monthly report"
	body := "Here is your monthly workout report"

	// Construct the email message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Set up the SMTP server details
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	password := "bxms jlgs szfp qpgw"
	smtpAuth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, smtpAuth, from, []string{to}, []byte(msg))
	if err != nil {
		log.Println("Error sending email:", err)
		return
	}

	log.Println("Email sent successfully!")
}
