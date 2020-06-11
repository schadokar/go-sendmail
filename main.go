package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	// Import godotenv
	"github.com/joho/godotenv"
)

func main() {
	// From address -- change the from to your sendinblue account mail id
	from := "hello@schadokar.dev"

	// Array of recipients address
	to := []string{"shubham@schadokar.dev"}

	// Create a message and convert it into bytes
	msg := []byte("To: shubham@schadokar.dev\r\n" +
		"From: hello@schadokar.dev\r\n" +
		"Subject: Hello Gophers!\r\n" +
		"\r\n" +
		"This is the email is sent using golang and sendinblue.\r\n")

	status := sendEmail(from, to, msg)

	if status {
		fmt.Printf("Email sent successfully.")
	} else {
		fmt.Printf("Email sent failed.")
	}
}

// send mail function
func sendEmail(from string, to []string, msg []byte) bool {
	// Load .env file to use the environment variable
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, os.Getenv("PASSWORD"), os.Getenv("SMTP_HOST"))

	// format smtp address
	smtpAddress := fmt.Sprintf("%s:%v", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail(smtpAddress, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
		return false
	}

	// return true on success
	return true
}
