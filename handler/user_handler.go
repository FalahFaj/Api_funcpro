package handler

import (
	"encoding/json"
	"net/http"
	"projek_funcpro_kel12/service"
	"strconv"
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
		return
	}
	var input service.LoginInput
	err := json.NewDecoder(terima.Body).Decode(&input)
	if err != nil {
		responError(balas, http.StatusBadRequest, "Format tidak sesuai")
		return
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

func (h *Userhandler) KelolaAkun(balas http.ResponseWriter, terima *http.Request) {
	idStr := terima.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responError(balas, http.StatusBadRequest, "ID pengguna tidak valid")
		return
	}

	userClaims := GetUserFromContext(terima.Context())
	if userClaims == nil {
		responError(balas, http.StatusUnauthorized, "Gagal mendapatkan info user dari token")
		return
	}

	if userClaims.UserId != id {
		responError(balas, http.StatusForbidden, "Akses ditolak. Anda hanya dapat mengelola akun Anda sendiri.")
		return
	}

	switch terima.Method {
	case http.MethodGet:
		h.GetUserById(balas, terima, id)
	case http.MethodPut:
		h.UpdateUser(balas, terima, id)
	case http.MethodDelete:
		h.DeleteUser(balas, terima, id)
	default:
		responError(balas, http.StatusMethodNotAllowed, "Method tidak valid")
	}
}

func (h *Userhandler) GetUserById(balas http.ResponseWriter, terima *http.Request, id int64) {
	user, err := h.userService.GetUserById(terima.Context(), id)
	if err != nil {
		responError(balas, http.StatusNotFound, err.Error())
		return
	}
	responJSON(balas, http.StatusOK, user)
}

func (h *Userhandler) UpdateUser(balas http.ResponseWriter, terima *http.Request, id int64) {
	var input service.RegisterInput
	err := json.NewDecoder(terima.Body).Decode(&input)
	if err != nil {
		responError(balas, http.StatusBadRequest, "Format tidak sesuai")
		return
	}
	user, err := h.userService.UpdateUser(terima.Context(), id, input)
	if err != nil {
		responError(balas, http.StatusInternalServerError, err.Error())
		return
	}
	responJSON(balas, http.StatusOK, user)
}

func (h *Userhandler) DeleteUser(balas http.ResponseWriter, terima *http.Request, id int64) {
	err := h.userService.DeleteUser(terima.Context(), id)
	if err != nil {
		responError(balas, http.StatusInternalServerError, err.Error())
		return
	}
	responJSON(balas, http.StatusOK, map[string]string{"message": "Pengguna berhasil dihapus"})
}