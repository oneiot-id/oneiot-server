package test

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"net/url"
	"oneiot-server/model/entity"
	"oneiot-server/service"
	"os"
	"strings"
	"testing"
)

func TestSendWhatsapp(t *testing.T) {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("Error loading .env file")
		t.Error("Error loading .env file")
	}

	accountSID := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_AUTH")
	from := "whatsapp:+" + os.Getenv("TWILIO_FROM_TEST")
	to := "whatsapp:+" + os.Getenv("TWILIO_TO_TEST")
	message := "Halo! Ini pesan dari Twilio WhatsApp API menggunakan Go."

	// Data yang akan dikirim
	data := url.Values{}
	data.Set("From", from)
	data.Set("To", to)
	data.Set("Body", message)

	// Buat request
	client := &http.Client{}
	req, _ := http.NewRequest("POST",
		"https://api.twilio.com/2010-04-01/Accounts/"+accountSID+"/Messages.json",
		strings.NewReader(data.Encode()))

	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Eksekusi request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))

}

func TestWhatsappService(t *testing.T) {
	_ = godotenv.Load("../.env")

	whatsAppService := service.NewWhatsAppService()

	code, err := whatsAppService.SendVerificationCode(entity.User{
		FullName:    "Vincent Kenutama Prasetyo",
		PhoneNumber: os.Getenv("TWILIO_TO_TEST"),
	})

	if err != nil {
		return

	}

	fmt.Println("Verification code: " + code)
}
