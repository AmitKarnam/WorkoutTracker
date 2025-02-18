package mailer

// import (
// 	"log"
// 	"net/smtp"

// 	"github.com/joho/godotenv"
// 	"github.com/knadh/koanf"
// 	"github.com/knadh/koanf/providers/env"
// )

// func mailer() {
// 	// Load environment variables from .env file (optional)
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found, proceeding with environment variables")
// 	}

// 	// Create a new koanf instance
// 	k := koanf.New(".")

// 	// Load environment variables into koanf
// 	if err := k.Load(env.Provider("", ".", func(s string) string {
// 		return s // No transformation on environment variable keys
// 	}), nil); err != nil {
// 		log.Fatalf("Error loading environment variables: %v", err)
// 	}

// 	// Set up email details
// 	from := k.String("FROM_EMAIL")
// 	to := k.String("TO_EMAIL")
// 	subject := "Monthly report"
// 	body := "Here is your monthly workout report"

// 	// Construct the email message
// 	msg := "From: " + from + "\n" +
// 		"To: " + to + "\n" +
// 		"Subject: " + subject + "\n\n" +
// 		body

// 	// Set up the SMTP server details
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"
// 	password := k.String("EMAIL_APP_KEY")
// 	smtpAuth := smtp.PlainAuth("", from, password, smtpHost)

// 	// Send the email
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, smtpAuth, from, []string{to}, []byte(msg))
// 	if err != nil {
// 		log.Println("Error sending email:", err)
// 		return
// 	}

// 	log.Println("Email sent successfully!")
// }
