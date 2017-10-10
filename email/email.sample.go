package email

import (
	"log"
	"net/smtp"
)

func SendMailExample(recipientEmail string, kode_verifikasi string) {
  // Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"mail@account.com",
		"yourpassword",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{recipientEmail}
	msg := []byte("To: "+ recipientEmail +"\r\n" +
		"Subject: Verifikasi Login!\r\n" +
		"\r\n" +
		"Masukkan kode verisikasi ini "+ kode_verifikasi +".\r\n")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"mail@account.com",
		to,
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}
}
