package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	"oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
	"os"
	"path/filepath"
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
	//Registering the user_pictures
	c.router.POST("/api/register", c.registerHandler)
	c.router.GET("/api/login", c.Login)
	c.router.GET("/api/user/", c.GetUser)
	c.router.POST("/api/user/upload-image", c.uploadImageHandler)
	c.router.ServeFiles("/static/*filepath", http.Dir("static"))
}

func (c *UserController) uploadImageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(4 * 1024)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	var user entity.User
	for key, value := range r.Form {
		switch key {
		case "user_email":
			user.Email = value[0]
		case "user_password":
			user.Password = value[0]
		}
	}

	user, err = c.service.GetUser(r.Context(), user)

	if err != nil {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	image, imageHandler, err := r.FormFile("image_data")

	if err != nil {
		http.Error(w, helper.MarshalThis(response.APIResponse[entity.User]{
			Message: "Error gambar tidak ditemukan saat upload",
			Data:    user,
		}), http.StatusBadRequest)
		return
	}
	defer image.Close()

	// Hapus gambar lama jika ada
	if user.Picture != "" {
		oldFilePath := filepath.Join("static/user_pictures", filepath.Base(user.Picture))
		os.Remove(oldFilePath)
	}

	// Simpan gambar baru
	dir, _ := os.Getwd()

	fileName := fmt.Sprintf("%d_%s", user.Id, imageHandler.Filename)
	fileLocation := filepath.Join(dir, "static/user_pictures", fileName)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, image)

	if err != nil {
		http.Error(w, helper.MarshalThis(response.APIResponse[entity.User]{
			Message: "Gagal menyimpan profile picture",
			Data:    user,
		}), http.StatusInternalServerError)
		return
	}

	// Simpan URL ke database
	publicURL := fmt.Sprintf(os.Getenv("LOCALHOST")+"/static/user_pictures/%s", fileName)
	user.Picture = publicURL

	updateUser, err := c.service.UpdateUser(r.Context(), user)

	if err != nil {
		http.Error(w, helper.MarshalThis(response.APIResponse[entity.User]{
			Message: err.Error(),
			Data:    user,
		}), http.StatusInternalServerError)
		return
	}

	// Kirim response ke klien
	_ = json.NewEncoder(w).Encode(response.APIResponse[entity.User]{
		Message: "Sukses mengubah profile picture",
		Data:    updateUser,
	})

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
