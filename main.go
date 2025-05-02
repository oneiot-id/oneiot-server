package main

import (
	"fmt"
	"log"
	"net/http"
	"oneiot-server/controller"
	"oneiot-server/database"
	"oneiot-server/email"
	"oneiot-server/helper"
	"oneiot-server/repository"
	"oneiot-server/service"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	//Load the environment variable
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = helper.LoadJWTConfig()
	if err != nil {
		log.Fatalf("Error loading JWT configuration: %v", err)
	}

	//Initializer
	router := httprouter.New()
	sqlDb := database.NewSqlConnection()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(router)
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	//Repository
	userRepository := repository.NewUserRepository(sqlDb)
	orderRepository := repository.NewOrderRepository(sqlDb)
	buyerRepository := repository.NewBuyerRepository(sqlDb)
	orderDetailRepository := repository.NewOrderDetailRepository(sqlDb)

	transactionRepository := repository.NewTransactionRepository(sqlDb)
	productionStatusRepository := repository.NewProductionStatusRepository(sqlDb)
	deliveryStatusRepository := repository.NewDeliveryStatusRepository(sqlDb)
	pricingRepository := repository.NewPricingRepository(sqlDb, 0.11)
	paymentRepository := repository.NewPaymentRepository(sqlDb)

	//Services
	whatsappHandler := service.NewWhatsAppService()
	emailHandler := &email.Email{}
	userService := service.NewUserService(userRepository, sqlDb)
	orderService := service.NewOrderService(userService, buyerRepository, orderDetailRepository, orderRepository)
	transactionService := service.NewTransactionService(sqlDb, transactionRepository, paymentRepository, pricingRepository, productionStatusRepository, deliveryStatusRepository)

	//Controller
	whatsappController := controller.NewWhatsappController(router, whatsappHandler)
	emailController := controller.NewEmailController(router, emailHandler, userService)
	orderController := controller.NewOrderController(router, userService, orderService)
	_ = controller.NewUserController(router, userService, sqlDb)
	transactionController := controller.NewTransactionController(router, userService, transactionService, orderService)

	//ToDo: we have to implement the middleware for API Key checking
	server := http.Server{
		Addr:    ":8000",
		Handler: corsHandler,
	}

	//ToDo: we need to implement safer than this, using go wire or something
	emailController.Serve()
	whatsappController.Serve()
	orderController.Serve()
	transactionController.Serve()

	fmt.Println(corsHandler)
	fmt.Println("[INFO] : Server running at  " + server.Addr)

	err = server.ListenAndServe()

	if err != nil {
		return
	}

}
