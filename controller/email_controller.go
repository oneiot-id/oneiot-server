package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/email"
	"oneiot-server/model/entity"
	"oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
	"time"
)

type EmailController struct {
	router       *httprouter.Router
	emailHandler *email.Email
	userService  service.IUserService
}

// NewEmailController construct new email controller
func NewEmailController(router *httprouter.Router, emailHandle *email.Email, userService service.IUserService) *EmailController {
	return &EmailController{router: router, emailHandler: emailHandle, userService: userService}
}

func (e *EmailController) Serve() {
	e.router.GET("/api/email/verification", e.handleVerificationCodeRequest)
}

func (e *EmailController) handleVerificationCodeRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var requestBody request.EmailVerificationRequestBody
	var responseBody response.EmailVerificationResponse

	//Decode the request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if requestBody.User.FullName == "" {
		responseBody.Messsage = "User full name cannot be empty"

		jsonResponse, _ := json.Marshal(responseBody)

		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "%s", string(jsonResponse))

		if err != nil {
			return
		}
		return
	}

	//This handle when client not providing the email
	if requestBody.User.Email == "" {
		responseBody.Messsage = "User email cannot be empty"

		jsonResponse, _ := json.Marshal(responseBody)

		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "%s", string(jsonResponse))

		if err != nil {
			return
		}

		return
	}

	if err != nil {
		return
	}

	//Cek dulu apakah email telah terpakai
	user := entity.User{
		Email:    requestBody.User.Email,
		Password: requestBody.User.Password,
	}

	isUserExisted, err := e.userService.CheckUserExistence(r.Context(), user)

	if isUserExisted {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.SimpleResponse{
			Message: "Pengguna dengan email ini telah ada, gunakan email lain",
			Data:    nil,
		})
		return
	}

	//Send the email
	res, err := e.emailHandler.SendVerificationEmail(requestBody.User)

	cookie := &http.Cookie{
		Name:    "EmailVerificationCode",
		Value:   res.Payload.UniqueCode,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)
	resJson, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)

	_, err = fmt.Fprintf(w, "%s", resJson)

	if err != nil {
		return
	}
}
