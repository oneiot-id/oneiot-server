package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/controller"
	"oneiot-server/database"
	"oneiot-server/email"
	"oneiot-server/repository"
	"oneiot-server/service"
)

func main() {
	//Load the environment variable
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
		return
	}

	//Initializer
	router := httprouter.New()
	sqlDb := database.NewSqlConnection()

	//Repository
	userRepository := repository.NewUserRepository(sqlDb)
	orderRepository := repository.NewOrderRepository(sqlDb)
	buyerRepository := repository.NewBuyerRepository(sqlDb)
	orderDetailRepository := repository.NewOrderDetailRepository(sqlDb)

	//Services
	whatsappHandler := service.NewWhatsAppService()
	emailHandler := &email.Email{}
	userService := service.NewUserService(userRepository, sqlDb)
	orderService := service.NewOrderService(userService, buyerRepository, orderDetailRepository, orderRepository)

	//Controller
	whatsappController := controller.NewWhatsappController(router, whatsappHandler)
	emailController := controller.NewEmailController(router, emailHandler, userService)
	orderController := controller.NewOrderController(router, userService, orderService)
	_ = controller.NewUserController(router, userService, sqlDb)

	//ToDo: we have to implement the middleware for API Key checking
	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	//ToDo: we need to implement safer than this, using go wire or something
	emailController.Serve()
	whatsappController.Serve()
	orderController.Serve()
	//userController.Serve()

	fmt.Println("Server running at " + server.Addr)

	err = server.ListenAndServe()

	if err != nil {
		return
	}

}
