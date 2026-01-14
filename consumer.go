package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {

	defer wg.Done()

	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "Just Testing our Email Campaign.")
		// msg := []byte(formattedMsg)

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker :%d Error parsing template for %s", id, recipient.Email)
			continue
		}

		fmt.Printf("Worker %d: Sending mail to %s \n", id, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "sianmol14@gmail.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: Sent mail to %s \n", id, recipient.Email)
	}
}
