package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Goreng", Harga: 10000, Stok: 10},
	{ID: 2, Nama: "Indomie Rebus", Harga: 10500, Stok: 10},
}

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{Id: 1, Name: "Kebutuhan rumah tangga", Description: "pembersih, deterjen, tisu, dan lainnya"},
	{Id: 2, Name: "Makanan dan minuman", Description: "bahan pokok, instan, camilan dan minuman"},
}

func getCategoryById(w http.ResponseWriter, r *http.Request) {
	// ambil id category dari url yang dikirimkan
	idCategory := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// konversi id ke int
	id, err := strconv.Atoi(idCategory)

	// cek jika ada error
	if err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	// ambil semua category
	for _, category := range categories {
		// ambil hanya category yang memiliki id yang sama
		if category.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}

	}

	// jika tidak ada category id yang dicari
	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	// ganti ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	var updateCategory Category

	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].Id == id {
			updateCategory.Id = id
			categories[i] = updateCategory
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategory)
			return
		}
	}

}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idCategoryStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	// ganti ke int
	id, err := strconv.Atoi(idCategoryStr)
	if err != nil {
		http.Error(w, "Invalid category id", http.StatusBadRequest)
		return
	}

	// loop category untuk mencari ID yang akan dihapus
	for index, category := range categories {
		if category.Id == id {
			// cara menghapusnya adalah dengan mengurangi isi didalam slice
			// membuat slice baru dengan data sebelum dan sesudah index yang ingin dihilangkan
			categories = append(categories[:index], categories[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "berhasil menghapus category",
			})
			return
		}
	}
	http.Error(w, "Category tidak ditemukan", http.StatusNotFound)
}

func getProdukByID(w http.ResponseWriter, r *http.Request) {
	// menghapus prefix, menyisakan hanya id produk
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// mengkonversi dari id yang bertipe string ke int
	id, err := strconv.Atoi(idStr)

	// cek jika ada error
	if err != nil {
		http.Error(w, "Invalid produk id", http.StatusBadRequest)
		return
	}

	// kalau tidak ada error, ambil produk ID berdasarkan id yang dikirimkan
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid produk id", http.StatusBadRequest)
		return
	}

	var updateProduk Produk

	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid produk id", http.StatusBadRequest)
		return
	}

	// loop produk untuk mencari ID yang akan dihapus
	for i, p := range produk {

		if p.ID == id {
			// cara menghapusnya adalah dengan mengurangi isi didalam slice
			// membuat slice baru dengan data sebelum dan sesudah index yang ingin dihilangkan
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "berhasil menghapus produk",
			})
			return
		}
	}
	http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)

}

func main() {
	// GET http://localhost:8080/api/categories
	// POST http://localhost:8080/api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		} else if r.Method == "POST" {
			var category Category
			err := json.NewDecoder(r.Body).Decode(&category)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			category.Id = len(categories) + 1
			categories = append(categories, category)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(category)
		}

	})

	// GET http://localhost:8080/api/categories/{id}
	// PUT http://localhost:8080/api/categories/{id}
	// DELETE http://localhost:8080/api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryById(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	// PUT http://localhost:8080/api/produk{id}
	// GET http://localhost:8080/api/produk/{id}
	// DELETE http://localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}

	})

	// GET http://localhost:8080/api/produk
	// POST http://localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		// cek method yang digunakan
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}
			// masukan data ke dalam variable produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			// set header
			w.Header().Set("Content-Type", "application/json")
			// kirim http status
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}

	})

	// http://localhost/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})
	fmt.Println("Server berjalan di port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
