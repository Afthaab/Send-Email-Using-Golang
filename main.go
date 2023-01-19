package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	loaderror := godotenv.Load()
	if loaderror != nil {
		fmt.Println(loaderror)
	}

	// Sender data.

	//EMAIL and PASSWORD has been saved in the ENV file for the enhanced security
	//And the password of the email should be the app password.

	from := os.Getenv("EMAIL")        //insert your from email here instead of os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD") //insert your from email password here instead of os.Getenv("EMAIL")

	// Receiver email address.
	to := []string{
		"afthab606@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("Enter your massage here")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
