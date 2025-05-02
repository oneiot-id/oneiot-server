package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/helper"
	"oneiot-server/middleware"
	"oneiot-server/model/dto"
	request2 "oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
)

type TransactionController struct {
	router             *httprouter.Router
	userService        service.IUserService
	orderService       service.IOrderService
	transactionService service.ITransactionService
}

func (controller *TransactionController) Serve() {
	controller.router.POST("/api/transaction/create", middleware.JWTMiddleware(controller.CreateTransactionOrder))
	controller.router.POST("/api/transaction/", middleware.JWTMiddleware(controller.GetTransactionHandler))
	controller.router.POST("/api/transactions/", middleware.JWTMiddleware(controller.GetUserTransactionsHandler))
}

func (controller *TransactionController) CreateTransactionOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.CreateTransactionRequest]
	w.Header().Set("Content-Type", "application/json")

	// 1. Get claims
	claims, _ := middleware.GetClaimsFromContext(r.Context())
	userId := claims.UserID

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return
	}

	//Get the order if it is available
	orderDto, err := controller.orderService.GetOrderById(r.Context(), request.Data.Order)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//return if the user and order is not same
	if int(orderDto.Order.UserId) != userId {
		middleware.UnauthorizedResponse(w, "User tidak memiliki izin untuk melakukan transaksi order ini")
		return
	}

	transactionToCreate := request.Data.TransactionDto
	transactionToCreate.Transaction.UserId = int64(userId)
	transactionToCreate.Transaction.OrderId = orderDto.Order.Id

	//Jika keduanya ditemukan maka buat order
	transactionDto, err := controller.transactionService.CreateTransaction(r.Context(), transactionToCreate)
	if err != nil {
		http.Error(w, helper.MarshalThis(response.SimpleResponse{Message: "Failed to create transaction: " + err.Error(), Data: nil}), http.StatusInternalServerError) // Or Bad Request
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.APIResponse[dto.TransactionDto]{
		Message: "Sukses membuat transaksi",
		Data:    transactionDto,
	})
}

func (controller *TransactionController) GetTransactionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.GetTransactionRequest]

	w.Header().Set("content-type", "application/json")

	claims, _ := middleware.GetClaimsFromContext(r.Context())
	userId := claims.UserID

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return
	}

	transactionDto, err := controller.transactionService.GetTransaction(r.Context(), request.Data.Transaction)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	if int(transactionDto.Transaction.UserId) != userId {
		middleware.UnauthorizedResponse(w, "Anda tidak dapat mengakses transaksi ini")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.APIResponse[dto.TransactionDto]{
		Message: "Sukses mendapatkan transaksi",
		Data:    transactionDto,
	})
}

func (controller *TransactionController) GetUserTransactionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.GetTransactionsRequest]

	claims, _ := middleware.GetClaimsFromContext(r.Context())
	userId := claims.UserID

	_ = json.NewDecoder(r.Body).Decode(&request)

	transactions, err := controller.transactionService.GetAllUserTransactions(r.Context(), int64(userId))
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	if transactions == nil {
		transactions = []dto.TransactionDto{}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.APIResponse[[]dto.TransactionDto]{
		Message: "Sukses mendapatkan transaksi user",
		Data:    transactions,
	})
}

func NewTransactionController(router *httprouter.Router, userService service.IUserService, transactionService service.ITransactionService, orderService service.IOrderService) TransactionController {
	return TransactionController{
		router:             router,
		userService:        userService,
		transactionService: transactionService,
		orderService:       orderService,
	}
}
