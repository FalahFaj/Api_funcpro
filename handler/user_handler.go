package handler

import (
	"encoding/json"
	"net/http"
	"projek_funcpro_kel12/service"
)

type Userhandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *Userhandler {
	return &Userhandler{userService}
}

func (h *Userhandler) Register(balas http.ResponseWriter, terima *http.Request) {
	if terima.Method != "POST" {
		responError(balas, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
		return
	}
	var input service.RegisterInput
	err := json.NewDecoder(terima.Body).Decode(&input)
	if err != nil {
		responError(balas, http.StatusBadRequest, "Gagal membaca input")
		return
	}
	user, err := h.userService.Register(terima.Context(), input)
	if err != nil {
		responError(balas, http.StatusBadRequest, err.Error())
		return
	}
	responJSON(balas, http.StatusCreated, user)
}

func (h *Userhandler) Login(balas http.ResponseWriter, terima *http.Request) {
	if terima.Method != "POST" {
		responError(balas, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
	}
	var input service.LoginInput
	err := json.NewDecoder(terima.Body).Decode(&input)
	if err != nil {
		responError(balas, http.StatusBadRequest, "Format tidak sesuai")
	}

	token, err := h.userService.Login(terima.Context(), input)
	if err != nil {
		responError(balas, http.StatusUnauthorized, err.Error())
		return
	}
	respon := map[string]string{"token": token}
	responJSON(balas, http.StatusOK, respon)
}

func responJSON(balas http.ResponseWriter, kode int, data any) {
	balas.Header().Set("Content-Type", "application/json")
	balas.WriteHeader(kode)
	json.NewEncoder(balas).Encode(data)
}

func responError(balas http.ResponseWriter, kode int, pesan string) {
	errorPayload := map[string]string{"error": pesan}
	responJSON(balas, kode, errorPayload)
}
