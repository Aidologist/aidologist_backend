package modules

import (
	"fmt"
	"net/smtp"
	"wkBackEnd/utils/constants"
)

func SendVerficationEmail(email string)  {
	message := []byte("Welcome to linkmoo, this is a verification email for your email address")
	auth := smtp.PlainAuth("", constants.GmailAddr, constants.GmailPassword, constants.GmailHost)
	err := smtp.SendMail(constants.GmailHost+":"+constants.GamilPort, auth, constants.GmailAddr, []string{email}, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}