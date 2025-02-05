package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
	"time"
)

type WhatsappController struct {
	router          *httprouter.Router
	whatsappService *service.WhatsAppService
}

func NewWhatsappController(router *httprouter.Router, whatsAppService *service.WhatsAppService) *WhatsappController {
	return &WhatsappController{
		whatsappService: whatsAppService,
		router:          router,
	}
}

func (wc *WhatsappController) Serve() {
	wc.router.GET("/api/whatsapp/verify", wc.getVerificationCode)
}

func (wc *WhatsappController) getVerificationCode(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var requestBody request.WhatsappVerificationRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	fmt.Println("[WHATSAPP] : Requesting for verification code. ")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		resJson, _ := json.Marshal(response.EmailVerificationResponse{
			Messsage: "Request error, check user body",
		})

		_, _ = fmt.Fprintf(w, string(resJson))
		return
	}

	//If the user_pictures full name or phone number is empty return bad request
	//We don't use validate at this moment, we'll use later i think
	if requestBody.User.FullName == "" || requestBody.User.PhoneNumber == "" {
		w.WriteHeader(http.StatusBadRequest)

		resJson, _ := json.Marshal(response.EmailVerificationResponse{
			Messsage: "User full name or phone number is empty",
		})

		_, _ = fmt.Fprintf(w, string(resJson))

		return
	}

	//Send the verification code to User WhatsApp
	uniqueCode, err := wc.whatsappService.SendVerificationCode(requestBody.User)

	//Get the now time
	expireTime := time.Now().Add(5 * time.Minute)

	if err != nil {
		return
	}
	resJson, _ := json.Marshal(request.GeneralVerificationResponse{
		Message: "Success send user verification code",
		Payload: request.GeneralVerificationCodePayload{
			UniqueCode:     uniqueCode,
			ExpireTimeUnix: time.Now().Add(5 * time.Minute).Unix(),
			User:           requestBody.User,
		},
	})

	//send back the payload if everything is okay
	http.SetCookie(w, &http.Cookie{
		Name:    "verificationCode",
		Value:   uniqueCode,
		Expires: expireTime,
	})

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, string(resJson))
}
