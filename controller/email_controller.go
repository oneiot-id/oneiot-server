package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/email"
	"oneiot-server/request"
	"oneiot-server/response"
)

type EmailController struct {
	router       *httprouter.Router
	emailHandler *email.Email
}

// NewEmailController construct new email controller
func NewEmailController(router *httprouter.Router, emailHandle *email.Email) *EmailController {
	return &EmailController{router: router, emailHandler: emailHandle}
}

func (e *EmailController) Serve() {
	e.router.GET("/api/email/verification", e.handleVerificationCodeRequest)
}

func (e *EmailController) handleVerificationCodeRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var requestBody request.EmailVerificationRequestBody
	var responseBody response.EmailVerificationResponse

	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	//This handle when client not providing the email
	if requestBody.User.Email == "" {
		responseBody.Messsage = "User email is empty"

		jsonResponse, _ := json.Marshal(responseBody)

		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "%s", string(jsonResponse))

		if err != nil {
			return
		}
	}

	if err != nil {
		return
	}

	//Send the email
	res, err := e.emailHandler.SendVerificationEmail(requestBody.User)

	resJson, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "%s", resJson)

	if err != nil {
		return
	}
}
