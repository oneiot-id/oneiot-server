package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	request2 "oneiot-server/request"
	"oneiot-server/response"
	"oneiot-server/service"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
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
	controller.router.PATCH("/api/order", controller.setOrderStatusHandler)
	controller.router.POST("/api/order/upload-brief", controller.uploadWorkBriefHandler)

}

func (controller *OrderController) uploadWorkBriefHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User
	var order entity.Order
	err := r.ParseMultipartForm(10 * 1024)

	//ToDo:
	// 0. Dapatkan semua data pada request [x]
	// 1. Cek terlebih dahulu apakah user valid [x]
	// 2. Dapatkan order
	// 3. Baru kita masukkan ke DTO file nya
	// 4. Buat order

	//0.
	//Data user
	user.Email = r.FormValue("user_email")
	user.Password = r.FormValue("user_password")

	//Data order
	order.Id, err = strconv.ParseInt(r.FormValue("order_id"), 10, 64)

	//1.
	//Login user
	user, err = controller.userService.LoginUser(r.Context(), user)

	//Jika tidak ada user dengan email dan password ini maka kembalikan
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response.APIResponse[response.UpdateOrderResponse]{
			Message: err.Error(),
			Data:    response.UpdateOrderResponse{},
		})
		return
	}

	//Cek order
	orderDTO, err := controller.orderService.GetOrderById(r.Context(), order)

	//Jika tidak ada order dengan id ini
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.APIResponse[response.UpdateOrderResponse]{
			Message: err.Error(),
			Data:    response.UpdateOrderResponse{},
		})
		return
	}

	//Return error jika order bukan milik user
	if user.Id != int(orderDTO.Order.UserId) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response.APIResponse[response.UpdateOrderResponse]{
			Message: "User tidak memiliki akses ke order ini",
			Data:    response.UpdateOrderResponse{},
		})
		return
	}

	//Logic disini untuk mengupdate brief file
	//Buka directory sekarang
	file, fileHandler, err := r.FormFile("brief_file")

	dir, _ := os.Getwd()

	//Buat file
	fileName := fmt.Sprintf("%d_%s_%s", user.Id, time.Now().Format("2006-01-02 15-04-05"), fileHandler.Filename)

	filePath := filepath.Join(dir, "static/order_briefs", fileName)

	newFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	_, err = io.Copy(newFile, file)

	if err != nil {
		return
	}

	defer newFile.Close()
	defer file.Close()

	//Update brief file
	orderDTO.OrderDetail.BriefFile = fmt.Sprintf("%s/static/order_briefs/%s", os.Getenv("LOCALHOST"), fileName)

	orderDTO, err = controller.orderService.UploadBriefFile(r.Context(), orderDTO, true)

	json.NewEncoder(w).Encode(response.APIResponse[response.UpdateBriefFile]{
		Message: "Sukses mengupdate brief file",
		Data: response.UpdateBriefFile{
			User:     user,
			OrderDTO: orderDTO,
		},
	})
}

func (controller *OrderController) setOrderStatusHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var request request2.APIRequest[request2.SetOrderRequest]
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Login user terlebih dahulu
	user, err := controller.userService.GetUser(r.Context(), request.Data.User)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	//Dapatkan terlebih dahulu ordernya
	order, err := controller.orderService.GetOrderById(r.Context(), request.Data.Order)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	//Update statusnya
	order, err = controller.orderService.SetStatus(r.Context(), request.Data.Order)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	//Jika berhasil kirim ke client
	json.NewEncoder(w).Encode(response.APIResponse[response.UpdateOrderResponse]{
		Message: "Sukses mengupdate status order",
		Data: response.UpdateOrderResponse{
			User:  user,
			Order: order.Order,
		},
	})
}

func (controller *OrderController) createOrderHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request request2.APIRequest[request2.CreateOrderRequest]

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = controller.userService.GetUser(r.Context(), request.Data.User)

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

	createdOrder, err := controller.orderService.CreateOrder(r.Context(), order, request.Data.User, request.Data.OrderDetail, request.Data.Buyer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, helper.MarshalThis(response.SimpleResponse{
			Message: err.Error(),
			Data:    nil,
		}))
		return
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

	w.Header().Set("Content-Type", "application-json")

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

	writer.Header().Set("Content-Type", "application/json")

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
