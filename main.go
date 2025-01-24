package main

import (
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/controller"
	"oneiot-server/email"
)

func main() {
	//Load the environment variable
	err := godotenv.Load()

	if err != nil {
		return
	}

	//Initialize router
	router := httprouter.New()

	//Initialize email handler
	emailHandler := &email.Email{}

	//Create constructor of email controller
	emailController := controller.NewEmailController(router, emailHandler)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	emailController.Serve()
	
	err = server.ListenAndServe()

	if err != nil {
		return
	}

}
