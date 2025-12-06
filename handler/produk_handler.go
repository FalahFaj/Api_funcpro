package handler

import (
	"encoding/json"
	"net/http"
	"projek_funcpro_kel12/service"
	"strconv"
)

type ProdukHandler struct {
	produkService service.ProdukService
}

func NewProdukHandler(produkService service.ProdukService) *ProdukHandler {
	return &ProdukHandler{produkService}
}

func (h *ProdukHandler) KelolaProduk(balas http.ResponseWriter, terima *http.Request) {
	switch terima.Method {
	case http.MethodGet:
		h.GetAllProduk(balas, terima)
	case http.MethodPost:
		h.CreateProduk(balas, terima)
	default:
		responError(balas, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
	}
}

func (h *ProdukHandler) CreateProduk(balas http.ResponseWriter, terima *http.Request) {
	var input service.InputProduk
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

	produk, err := h.produkService.CreateProduk(terima.Context(), userClaims.UserId, input)
	if err != nil {
		responError(balas, http.StatusInternalServerError, err.Error())
		return
	}
	responJSON(balas, http.StatusCreated, produk)
}

func (h *ProdukHandler) GetAllProduk(balas http.ResponseWriter, terima *http.Request) {
	produks, err := h.produkService.GetAllProduk(terima.Context())
	if err != nil {
		responError(balas, http.StatusInternalServerError, err.Error())
		return
	}
	responJSON(balas, http.StatusOK, produks)
}

func (h *ProdukHandler) KelolaProdukById(balas http.ResponseWriter, terima *http.Request) {
	idStr := terima.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		responError(balas, http.StatusBadRequest, "ID produk tidak valid")
		return
	}

	switch terima.Method {
	case http.MethodGet:
		h.GetProdukById(balas, terima, id)
	case http.MethodPut:
		h.UpdateProduk(balas, terima, id)
	case http.MethodDelete:
		h.DeleteProduk(balas, terima, id)
	default:
		responError(balas, http.StatusMethodNotAllowed, "Metode tidak diizinkan")
	}
}

func (h *ProdukHandler) DeleteProduk(balas http.ResponseWriter, terima *http.Request, id int64) {
	err := h.produkService.DeleteProduk(terima.Context(), id)
	if err != nil {
		responError(balas, http.StatusInternalServerError, err.Error())
		return
	}
	responJSON(balas, http.StatusOK, map[string]string{"message": "Produk berhasil dihapus"})
}

func (h *ProdukHandler) GetProdukById(balas http.ResponseWriter, terima *http.Request, id int64) {
	produk, err := h.produkService.GetProdukById(terima.Context(), id)
	if err != nil {
		responError(balas, http.StatusNotFound, err.Error())
		return
	}
	responJSON(balas, http.StatusOK, produk)
}

func (h *ProdukHandler) UpdateProduk(balas http.ResponseWriter, terima *http.Request, id int64) {
	var input service.InputProduk
	err := json.NewDecoder(terima.Body).Decode(&input)
	if err != nil {
		responError(balas, http.StatusBadRequest, "Format tidak sesuai")
		return
	}

	userClaims := GetUserFromContext(terima.Context())
	if userClaims == nil {
		responError(balas, http.StatusUnauthorized, "Gagal mendapatkan info user dari token")
		return
	}
	produkToUpdate, err := h.produkService.GetProdukById(terima.Context(), id)
	if err != nil {
		responError(balas, http.StatusNotFound, "Produk tidak ditemukan")
		return
	}
	if produkToUpdate.PetaniId != userClaims.UserId {
		responError(balas, http.StatusForbidden, "Anda tidak berhak mengubah produk ini")
		return
	}

	produk, err := h.produkService.UpdateProduk(terima.Context(), id, input)
	if err != nil {
		http.Error(balas, err.Error(), http.StatusInternalServerError)
		return
	}
	responJSON(balas, http.StatusOK, produk)
}
