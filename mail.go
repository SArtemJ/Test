package main

import (
	"log"
	"net/smtp"
	"fmt"
)

func Send(fm string, fp string, tm string, m string) {

	var k []string
	k = append(k, tm)

//var msg = `From: fortest011@yandex.ru
//To: fortest01223@yandex.ru
//Subject: Error message
//MIME-Version: 1.0
//Content-Transfer-Encoding: 8bit
//Content-Type: text/plain; charset="UTF-8"`

	fmt.Println(m)
	fmt.Println(tm)
	fmt.Println(k)

	auth := smtp.PlainAuth("",
	fm,
	fp,
	"smtp.yandex.ru")

	err := smtp.SendMail("smtp.yandex.ru:465",
		auth,
		fm,
		k,
		[]byte(m),
	)

	if err != nil {
		log.Println("Error sending")
		return
	}

}