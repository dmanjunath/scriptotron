package main

import (
  "log"
  "net/smtp"
)

func sendEmail(body string) {
  from := config.Gmail.Email
  pass := config.Gmail.Password
  to := config.Gmail.Email

  msg := "From: " + from + "\n" +
    "To: " + to + "\n" +
    "Subject: Script-O-tron notification!\n\n" +
    body

  err := smtp.SendMail("smtp.gmail.com:587",
    smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
    from, []string{to}, []byte(msg))

  if err != nil {
    log.Printf("smtp error: %s", err)
    return
  }
}