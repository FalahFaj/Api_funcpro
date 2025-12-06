package handler

import (
	"encoding/json"
	"net/http"
	"projek_funcpro_kel12/service"
	"strconv"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService}
}

func (h *OrderHandler) CreateOrder(balas http.ResponseWriter, terima *http.Request) {
	var input service.CreateOrderInput
	err := json.NewDecoder(terima.Body).Decode(&input)
	if err != nil {
		responError(balas, http.StatusBadRequest, "Gagal membaca input")
		return
	}

	userClaims := GetUserFromContext(terima.Context())
	if userClaims == nil {
		responError(balas, http.StatusUnauthorized, "Gagal mendapatkan info user dari token")
		return
	}

	order, err := h.orderService.CreateOrder(terima.Context(), userClaims.UserId, input)
	if err != nil {
		responError(balas, http.StatusBadRequest, err.Error())
		return
	}
	responJSON(balas, http.StatusCreated, order)
}

func (h *OrderHandler) GetAllOrder(balas http.ResponseWriter, terima *http.Request) {
	orders, err := h.orderService.GetAllOrder(terima.Context())
	if err != nil {
		responError(balas, http.StatusInternalServerError, err.Error())
		return
	}
	responJSON(balas, http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderById(balas http.ResponseWriter, terima *http.Request) {
	idStr := terima.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responError(balas, http.StatusBadRequest, "ID order tidak valid")
		return
	}
	order, err := h.orderService.GetOrderById(terima.Context(), id)
	if err != nil {
		responError(balas, http.StatusNotFound, err.Error())
		return
	}

	userClaims := GetUserFromContext(terima.Context())
	if userClaims == nil {
		responError(balas, http.StatusUnauthorized, "Gagal mendapatkan info user dari token")
		return
	}

	if userClaims.Role != "admin" && order.PembeliId != userClaims.UserId {
		responError(balas, http.StatusForbidden, "Akses ditolak. Anda tidak berhak melihat pesanan ini.")
		return
	}

	responJSON(balas, http.StatusOK, order)
}
