package service

import (
	"io"
	"net/http"
	"net/url"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	"os"
	"strings"
)

type WhatsAppService struct {
	baseUrl     string
	accountSid  string
	accountAuth string
	client      *http.Client
}

func NewWhatsAppService() *WhatsAppService {
	return &WhatsAppService{
		accountSid:  os.Getenv("TWILIO_SID"),
		accountAuth: os.Getenv("TWILIO_AUTH"),
		baseUrl:     "https://api.twilio.com/2010-04-01/Accounts/" + os.Getenv("TWILIO_SID") + "/Messages.json",
		client:      &http.Client{},
	}
}

func (w *WhatsAppService) SendVerificationCode(to entity.User) (string, error) {
	uniqueCode := helper.GenerateRandomVerificationCode(4)

	data := url.Values{}

	data.Set("From", "whatsapp:+"+os.Getenv("TWILIO_FROM_TEST"))

	if strings.Contains(to.PhoneNumber, "+") {
		data.Set("To", "whatsapp:"+to.PhoneNumber)
	} else {
		data.Set("To", "whatsapp:+"+to.PhoneNumber)
	}

	message := "Halo " + to.FullName + ", selamat datang di layanan OneIoT. Kode verifikasi akun Anda adalah *" + uniqueCode + "*. Demi keamanan, jangan bagikan kode ini kepada siapa pun. Harap diingat bahwa kode ini hanya berlaku selama 5 menit."

	data.Set("Body", message)

	request, err := http.NewRequest(http.MethodPost, w.baseUrl, strings.NewReader(data.Encode()))

	request.SetBasicAuth(w.accountSid, w.accountAuth)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := w.client.Do(request)

	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	return uniqueCode, nil
}
