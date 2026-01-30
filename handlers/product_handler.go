package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

// buat sebuah struct yang isinya adalah pointer ke ProductService
// karena sebuah pointer, maka seperti membuat sebuah link ke ProductService
type ProductHandler struct {
	service *services.ProductService
}

// kalau dicari penjelasannya ini adalah sebuah contructor
// yang tujuannya membuat object ProductHandler.
func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// ini digunakan untuk menghandel request masuk dari client
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	// cek http methodnya
	switch r.Method {
	case http.MethodGet:
		// jika methodnya GET maka jalankan ini
		h.GetProducts(w, r)
	case http.MethodPost:
		// jika methodnya POST maka jalankan ini
		h.CreateProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// jika methodnya GET, maka kesini
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// eksekusi fungsi GetProducts pada ProductService
	// services/product_service.go
	products, err := h.service.GetProducts()
	// cek jika ada error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set header content type menjadi json
	w.Header().Set("Content-Type", "application/json")
	// kirim response
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// buat variable kosong yang bertipe Product
	var product models.Product

	// baca JSON yang dikirimkan oleh klien
	// lalu masukan isinya ke variable product atau isi field struct product
	err := json.NewDecoder(r.Body).Decode(&product)

	// cek apakah ada error
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// eksekusi fungsi CreateProduct pada ProductService
	// dan kririm product nya
	err = h.service.CreateProduct(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)

}

// handler utama untuk semua endpoint yang menggunakan ID
func (h *ProductHandler) HandleProductById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetProductById(w, r)
	case http.MethodPut:
		h.UpdateProduct(w, r)
	case http.MethodDelete:
		h.DeleteProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// ambil id dari url
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")

	// konversi ke int
	id, err := strconv.Atoi(idStr)

	// cek jika ada error
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// buat sebuah variabel bertipe Product struct
	var product models.Product

	// baca JSON yang dikirimkan oleh klien dan lakukan decode
	// lalu masukan isinya ke variable product
	err = json.NewDecoder(r.Body).Decode(&product)

	// cek jika ada error
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// set ID, untuk memastikan ID yang akan diganti sama dengan id yang dikirim
	product.ID = id

	// eksekusi fungsi UpdateProduct pada ProductService / sercies/product_service.go
	err = h.service.UpdateProduct(&product)

	// cek jika ada error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// kirim response header dan body
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Berhasil menghapus produk",
	})
}
