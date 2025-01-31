package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/helper"
	"oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
)

type UserController struct {
	router  *httprouter.Router
	service service.IUserService
	db      *sql.DB
}

func NewUserController(router *httprouter.Router, userService *service.UserService, db *sql.DB) *UserController {
	userController := &UserController{
		service: userService,
		router:  router,
	}

	userController.Serve()

	return userController
}

func (c *UserController) Serve() {
	//Registering the user
	c.router.POST("/api/register", c.registerHandler)
	c.router.GET("/api/login", c.Login)
	c.router.GET("/api/user/", c.GetUser)
}

// todo: We do the validation on the frontend only, next we will try to validate on the backend
func (c *UserController) registerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userToRegister request.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&userToRegister)

	//If something when wrong at the parsing then do return
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		out := helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		})

		_, _ = fmt.Fprint(w, out)
	}

	//else do the registering
	registeredUser, err := c.service.RegisterNewUser(r.Context(), userToRegister.User)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		out := helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		})
		_, _ = fmt.Fprint(w, out)
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprint(w, helper.MarshalThis(response.SimpleResponse{
		Message: "Successfully registered user",
		Data:    registeredUser,
	}))
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userToLoginRequest request.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&userToLoginRequest)

	//If the decode is error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		out := helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		})

		_, _ = fmt.Fprint(w, out)
	}

	loginUser, err := c.service.LoginUser(r.Context(), userToLoginRequest.User)

	//If something went wrong
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		}))
		if err2 != nil {
			return
		}
	}

	_, _ = fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
		Message: "Successfully logged in",
		Data:    loginUser,
	}))
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userToLoginRequest request.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&userToLoginRequest)

	//If the decode is error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		out := helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		})

		_, _ = fmt.Fprint(w, out)
	}

	getUser, err := c.service.GetUser(r.Context(), userToLoginRequest.User)

	//If something went wrong
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		}))
		if err2 != nil {
			return
		}
	}

	_, _ = fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
		Message: "Successfully get user",
		Data:    getUser,
	}))
}
