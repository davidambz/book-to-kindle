package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	toEmail := os.Getenv("TO_EMAIL")
	subject := "CONVERT"
	booksDir := os.Getenv("BOOKS_DIR")

	if booksDir == "" {
		log.Fatalf("Books directory not specified in .env file")
	}

	if len(os.Args) < 2 {
		log.Fatalf("Please specify the file name as an argument in quotes")
	}

	filename := os.Args[1]

	attachmentPath := filepath.Join(booksDir, filename)

	attachmentData, err := os.ReadFile(attachmentPath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	headers := map[string]string{
		"From":         email,
		"To":           toEmail,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": fmt.Sprintf(`multipart/mixed; boundary="%s"`, writer.Boundary()),
	}

	for key, value := range headers {
		body.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	body.WriteString("\r\n")

	part, err := writer.CreatePart(nil)
	if err != nil {
		log.Fatalf("Error creating text part: %s", err)
	}
	part.Write([]byte(""))

	part, err = writer.CreatePart(map[string][]string{
		"Content-Type":              {fmt.Sprintf(`application/octet-stream; name="%s"`, filepath.Base(attachmentPath))},
		"Content-Transfer-Encoding": {"base64"},
		"Content-Disposition":       {fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(attachmentPath))},
	})
	if err != nil {
		log.Fatalf("Error creating attachment part: %s", err)
	}

	encoder := base64.NewEncoder(base64.StdEncoding, part)
	_, err = encoder.Write(attachmentData)
	if err != nil {
		log.Fatalf("Error encoding file content: %s", err)
	}
	encoder.Close()

	writer.Close()

	auth := smtp.PlainAuth("", email, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{toEmail}, body.Bytes())
	if err != nil {
		log.Fatalf("Error sending email: %s", err)
	} else {
		log.Println("Email sent successfully!")
	}
}
