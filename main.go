package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// membaca file .env
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// mengambil nilai PORT pada file .env
	config := Config{
		// ambil port di file .env
		Port: viper.GetString("PORT"),
		// ambil database connection di file .env
		DBConn: viper.GetString("DB_CONN"),
	}

	// mendefinisikan address dan port untuk server
	addr := "0.0.0.0:" + config.Port

	// setup koneksi database
	// gunakan package database yang sudah dibuat tadi, kemudian panggil fungsi InitDB nya.
	// kirim config.DBConn sebagai argument ke InitDB
	db, err := database.InitDB(config.DBConn)
	// cek jika ada error
	if err != nil {
		log.Fatal("Failed to initialize database :", err)
	}
	defer db.Close()

	// setup handler
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductById)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryById)

	// http://localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	//  cetak ke terminal
	fmt.Println("Server berjalan di", addr)
	// menjalankan server
	// Listen harus diakhir, handler harus diawal
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}

}
