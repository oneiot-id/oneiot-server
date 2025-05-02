package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"oneiot-server/helper"
	"oneiot-server/middleware"
	"oneiot-server/model/entity"
	"oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
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
		db:      db,
	}

	userController.Serve()

	return userController
}

func (c *UserController) Serve() {
	//Registering the user_pictures
	c.router.POST("/api/register", c.registerHandler)
	c.router.POST("/api/login", c.Login)
	c.router.POST("/api/logout", c.Logout)
	c.router.POST("/api/user/upload-image", middleware.JWTMiddleware(c.uploadImageHandler))
	c.router.POST("/api/user", middleware.JWTMiddleware(c.GetUser))
}

func (c *UserController) uploadImageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	claims, _ := middleware.GetClaimsFromContext(r.Context())

	err := r.ParseMultipartForm(4 * 1024 * 1024)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	user, err := c.service.GetUserByID(r.Context(), claims.UserID)
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

	updateUser.Password = ""

	// Kirim response ke klien
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.APIResponse[entity.User]{
		Message: "Sukses mengubah profil picture",
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

	tokenString, expirationTime, err := helper.GenerateJWT(registeredUser)
	if err != nil {
		http.Error(w, helper.MarshalThis(response.SimpleResponse{Message: "User registered, but failed to generate session token", Data: nil}), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     middleware.CookieName,
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("APP_ENV") == "production",
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.SimpleResponse{
		Message: "Successfully registered user",
		Data:    registeredUser,
	})
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userToLoginRequest request.UserLoginRequest
	w.Header().Set("content-type", "application/json")
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
		http.Error(w, helper.MarshalThis(response.SimpleResponse{Message: err.Error(), Data: nil}), http.StatusUnauthorized) // Use 401 for login failure
		return
	}

	tokenString, expirationTime, err := helper.GenerateJWT(loginUser)
	if err != nil {
		http.Error(w, helper.MarshalThis(response.SimpleResponse{Message: "Failed to generate token", Data: nil}), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     middleware.CookieName,
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/", // Important for cookie scope
		HttpOnly: true,
		Secure:   os.Getenv("APP_ENV") == "production", // Use Secure flag in production
		SameSite: http.SameSiteLaxMode,                 // Or StrictMode
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SimpleResponse{
		Message: "Successfully logged in",
		Data:    loginUser,
	})
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	claims, _ := middleware.GetClaimsFromContext(r.Context())

	getUserData, err := c.service.GetUserByID(r.Context(), claims.UserID)

	if err != nil {
		http.Error(w, helper.MarshalThis(response.SimpleResponse{
			Message: "Failed to retrieve user data: " + err.Error(),
			Data:    nil,
		}), http.StatusNotFound)
		return
	}

	getUserData.Password = ""

	json.NewEncoder(w).Encode(response.SimpleResponse{
		Message: "Successfully get user",
		Data:    getUserData,
	})
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	middleware.ClearAuthCookie(w)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SimpleResponse{
		Message: "Successfully logged out",
		Data:    nil,
	})
}
