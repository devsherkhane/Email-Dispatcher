package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipent, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipent := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipent.Email, "Just testing our email campaign")
		msg := []byte(formattedMsg)

		fmt.Printf("Worker %d: Sending email to %s \n", id, recipent.Email)

		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "devsherkhane6@gmail.com", []string{recipent.Email}, msg)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker %d: Sent email to %s \n", id, recipent.Email)

	}

}
