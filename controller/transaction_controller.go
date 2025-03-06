package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	request2 "oneiot-server/request"
	"oneiot-server/service"
)

type TransactionController struct {
	router             *httprouter.Router
	userService        service.IUserService
	orderService       service.IOrderService
	transactionService service.ITransactionService
}

func (controller *TransactionController) Serve() {
	controller.router.POST("/api/transaction/create", controller.CreateTransactionOrder)
	controller.router.POST("/api/transaction/", controller.GetTransactionHandler)
	controller.router.POST("/api/transactions/", controller.GetUserTransactionsHandler)
}

func (controller *TransactionController) CreateTransactionOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.CreateTransactionRequest]

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return
	}

	//Get the user first
	user, err := controller.userService.LoginUser(r.Context(), request.Data.User)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//Get the order if it is available
	orderDto, err := controller.orderService.GetOrderById(r.Context(), request.Data.Order)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//return if the user and order is not same
	if user.Id != int(orderDto.Order.UserId) {
		fmt.Fprintf(w, "User dan order tidak valid ")
		return
	}

	//Jika keduanya ditemukan maka buat order
	transactionDto, err := controller.transactionService.CreateTransaction(r.Context(), request.Data.TransactionDto)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(transactionDto)
}

func (controller *TransactionController) GetTransactionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.GetTransactionRequest]
	
	w.Header().Set("content-type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&request)

	//Get the user first
	user, err := controller.userService.LoginUser(r.Context(), request.Data.User)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	transactionDto, err := controller.transactionService.GetTransaction(r.Context(), request.Data.Transaction)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	if user.Id != int(transactionDto.Transaction.UserId) {
		fmt.Fprintf(w, "Anda tidak dapat mengakses transaksi ini ")
		return
	}

	json.NewEncoder(w).Encode(transactionDto)
}

func (controller *TransactionController) GetUserTransactionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.GetTransactionsRequest]

	_ = json.NewDecoder(r.Body).Decode(&request)

	//Get the user first
	user, err := controller.userService.LoginUser(r.Context(), request.Data.User)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	transactions, err := controller.transactionService.GetAllUserTransactions(r.Context(), int64(user.Id))

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(transactions)
}

func NewTransactionController(router *httprouter.Router, userService service.IUserService, transactionService service.ITransactionService, orderService service.IOrderService) TransactionController {
	return TransactionController{
		router:             router,
		userService:        userService,
		transactionService: transactionService,
		orderService:       orderService,
	}
}
