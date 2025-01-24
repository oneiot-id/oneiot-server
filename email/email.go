package email

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"math/rand"
	"oneiot-server/model/entity"
	"oneiot-server/response"
	"os"
	"strconv"
	"time"
)

type Email struct {
}

func GenerateRandomVerificationCode(length int) string {
	uniqueCode := ""

	for i := 0; i < length; i++ {
		random := strconv.Itoa(rand.Intn(10))
		uniqueCode += random
	}

	fmt.Println(uniqueCode)

	return uniqueCode
}

// This return the 5 minutest expired time in unix for the unique code
func getExpiredTimeUnix() int64 {
	return time.Now().Unix() + 5*60*1000
}

func (e *Email) SendVerificationEmail(to entity.User) (response.EmailVerificationResponse, error) {
	//Create new Status
	uniqueCode := GenerateRandomVerificationCode(4)
	message := gomail.NewMessage()

	message.SetHeader("From", os.Getenv("SMTP_USER"))
	message.SetHeader("To", to.Email)
	message.SetHeader("Subject", "Verifikasi Email Akun OneIoT")

	// Replace placeholders with actual values
	htmlBody := `
<html>
  <head>
    <style>
      /* Styles here */
    </style>
  </head>
  <body>
    <div class="container">
      <h1>Verifikasi Akun OneIoT</h1>
      <p>Halo ` + to.FullName + `,</p>
      <p>Yuk mulai aktivasi akun kamu!</p>
      <p>Cukup masukkan kode verifikasi di bawah ini untuk aktivasi akun kamu: <strong>` + uniqueCode + `</strong></p>
      <p>Kode hanya berlaku 5 menit. Mohon jangan menyebarkan kode ini ya!</p>
      <p>Terima kasih telah mendaftar ke OneIoT Partner, kami tunggu pesanan Anda!</p>
      
      <div class="footer">
        <p>Jika ada pertanyaan Anda dapat menghubungi <a>oneiot@gmail.com</a></p>
      </div>
    </div>
  </body>
</html>
`

	// Set email body
	message.SetBody("text/html", htmlBody)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)

		panic(err)
	} else {
		fmt.Println("[Email Handler] : Email sent successfully!")
	}

	payload := response.EmailVerificationBody{
		User:           to,
		UniqueCode:     uniqueCode,
		ExpireTimeUnix: getExpiredTimeUnix(),
	}

	res := response.EmailVerificationResponse{
		Payload:  payload,
		Messsage: "Successfully sent verification email",
	}

	return res, nil
}
