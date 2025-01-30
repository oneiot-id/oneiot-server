package main

import (
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/controller"
	"oneiot-server/email"
	"oneiot-server/service"
)

func main() {
	//Load the environment variable
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file aaa")
		panic(err)
		return
	}

	//Initialize router
	router := httprouter.New()

	//Services
	whatsappHandler := service.NewWhatsAppService()
	emailHandler := &email.Email{}

	//Controller
	whatsappController := controller.NewWhatsappController(router, whatsappHandler)
	emailController := controller.NewEmailController(router, emailHandler)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	emailController.Serve()
	whatsappController.Serve()

	err = server.ListenAndServe()

	if err != nil {
		return
	}

}
