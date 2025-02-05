package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	request2 "oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
	"time"
)

type OrderController struct {
	router       *httprouter.Router
	userService  service.IUserService
	orderService service.IOrderService
}

func (controller *OrderController) Serve() {
	controller.router.GET("/api/order", controller.getOrderHandler)
	controller.router.GET("/api/orders", controller.getAllUserOrders)
	controller.router.POST("/api/order", controller.createOrderHandler)
}

func (controller *OrderController) createOrderHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.CreateOrderRequest]

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := controller.userService.GetUser(r.Context(), request.Data.User)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		}))

		return
	}

	order := entity.Order{
		IsActive:  false,
		CreatedAt: time.Now(),
	}

	createdOrder, err := controller.orderService.CreateOrder(r.Context(), order, user, request.Data.OrderDetail, request.Data.Buyer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		}))
	}

	//Jika semua berhasil maka kirim
	fmt.Fprintf(w, helper.MarshalThis(response.APIResponse[response.CreateOrderResponse]{
		Message: "Sukses membuat order",
		Data: response.CreateOrderResponse{
			Order: createdOrder,
		},
	}))

	fmt.Println(createdOrder)
}

func (controller *OrderController) getAllUserOrders(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestData request2.APIRequest[request2.GetOrdersRequest]

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orders, err := controller.orderService.GetAllUserOrder(r.Context(), requestData.Data.User)

	if err != nil {
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil}))

		return
	}

	fmt.Fprintf(w, helper.MarshalThis(response.APIResponse[response.GetAllOrdersResponse]{
		Message: "Sukses mendapatkan data",
		Data: response.GetAllOrdersResponse{
			Orders: orders,
		},
	}))

	fmt.Println(requestData)
	fmt.Println(orders)
}

func (controller *OrderController) getOrderHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var requestData request2.APIRequest[request2.GetOrderRequest]

	err := json.NewDecoder(request.Body).Decode(&requestData)

	if err != nil {
		panic(err)
	}

	//ToDo:
	// 1. we need to verify if the user_pictures is valid to get the order [x]
	// 2. then we can get the order [x]
	// 3. Hmm kayaknya butuh authorisasi user_pictures, jika user_pictures berbeda dengan order user_pictures id maka batalkan

	user, err := controller.userService.GetUser(request.Context(), requestData.Data.User)

	//Check if the user_pictures is logged in / valid or not
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintf(writer, helper.MarshalThis(response.SimpleResponse{
			Message: "Unauthorized, pengguna ini tidak dapat mengakses data, karena terdapat kesalahan pada data yang diberikan",
			Data:    nil}))

		return
	}

	orderDTOResponse, err := controller.orderService.GetOrderById(request.Context(), requestData.Data.Order)

	if int(orderDTOResponse.Order.UserId) != user.Id {
		writer.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintf(writer, helper.MarshalThis(response.SimpleResponse{
			Message: "Unauthorized, pengguna ini tidak dapat mengakses data, karena order ini bukan milik pengguna ini",
			Data:    nil}))

		return
	}

	//Check if the order id is valid or not
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil}))
		return
	}

	res := helper.MarshalThis(orderDTOResponse)
	_, err = fmt.Fprint(writer, res)

	if err != nil {
		return
	}
}

func NewOrderController(router *httprouter.Router, userService service.IUserService, orderService service.IOrderService) *OrderController {
	return &OrderController{
		router:       router,
		userService:  userService,
		orderService: orderService,
	}
}
