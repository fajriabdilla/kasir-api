package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetProducts() ([]models.Product, error) {
	// buat sebuah query untuk mengambil seluruh product
	query := "SELECT id, name, price, stock FROM products"

	// ini adalah proses eksekusi query dan menampung hasilnya dalam variabel row dan juga err
	// karena eksekusi ini mengembalikan dua data, yakni hasil query dan error
	rows, err := repo.db.Query(query)

	// cek jika terjadi error
	// jika terjadi error, maka kembalikan nil dan error, kenapa?
	// karena fungsi ini mengembalikan dua data, yakni hasil query dan error
	if err != nil {
		return nil, err
	}

	// tutup connection pool database
	defer rows.Close()

	// buat sebuah slice kosong yang panjangnya 0
	// ini berfungsi untuk menampung hasil query
	products := make([]models.Product, 0)

	// lakukan loop sebanyak data yang didalam variabel rows
	for rows.Next() {
		// buat sebuah variabel p yang bertipe models.Product
		// untuk menampung setiap baris data rows
		var p models.Product

		// mengambil data yang ada pada variable rows
		// dan isikan kedalam variable p
		// jumlah nilai yang dimasukan harus sama dengan jumlah data yang di dalam variable rows
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	// ingat untuk selalu mengembalikan dua buah nilai
	return products, nil

}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	// buat sebuah query insert
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id"
	// eksekusi query
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID)
	return err
}

// buat sebuah method GetProductById yang menerima parameter id int, mengembalikan pointer models.Product dan error
func (repo *ProductRepository) GetProductById(id int) (*models.Product, error) {
	// buat sebuah query untuk mengambil product berdasarkan id
	query := "SELECT id, name, price, stock FROM products where id = $1"

	// buat sebuah variable product bertipe Product
	var product models.Product

	// eksekusi query
	err := repo.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)

	// cek jika data produk tidak ditemukan
	if err == sql.ErrNoRows {
		return nil, errors.New("Produk tidak ditemukan")
	}

	// cek jika ada error
	if err != nil {
		return nil, err
	}

	// kembalikan data produk
	return &product, nil
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4"

	// eksekusi query
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.ID)

	// cek jika ada error
	if err != nil {
		return err
	}

	// cek row yang terpengaruh
	rows, err := result.RowsAffected()

	// cek jika ada error
	if err != nil {
		return err
	}

	// cek jika tidak ada row yang terpengaruh
	if rows == 0 {
		return errors.New("Produk tidak ditemukan")
	}

	// karena tidak ada error, maka kembalikan nil
	// yang penting harus ada yang dikembalikan
	return nil
}

func (repo *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = $1"

	// eksekusi query dan tangkap sql resultnya
	result, err := repo.db.Exec(query, id)

	if err != nil {
		return err
	}

	// cek row yang terpengaruh
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// cek jika tidak ada row yang terpengaruh
	if rows == 0 {
		return errors.New("Produk tidak ditemukan")
	}

	return err
}
