package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) GetCategories() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"

	// fungsi Query() menerima 2 argument, yaitu query string dan variadic arguments
	// karena argument keduanya adalah variadic, maka argument kedua menjadi optional
	// kapan diisi? ketika ada placeholder di dalam query
	rows, err := repo.db.Query(query)
	if err != nil {
		// ingat untuk selalu mengembalikan 2 nilai, jika fungsi yang kita buat mengembalikan 2 nilai
		return nil, err
	}

	defer rows.Close()

	// buat slice kosong
	categories := make([]models.Category, 0)
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (repo *CategoryRepository) CreateCategory(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"

	// fungsi QueryRow() mengembalikan satu baris data
	// dan Scan() mengembalikan error
	// ambil id yang baru saja diinsert menggunakan fungsi Scan()
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.Id)
	return err
}

func (repo *CategoryRepository) GetCategoryById(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	var category models.Category

	// database mencari satu baris berdasarkan id
	err := repo.db.QueryRow(query, id).
		// ambil hasil dari query database, lalu "MASUKAN KE ALAMAT" field field category.
		Scan(&category.Id, &category.Name, &category.Description)

	if err == sql.ErrNoRows {
		return nil, errors.New("Kategori tidak ditemukan")
	}

	if err != nil {
		return nil, err
	}

	// kirim ALAMAT CATEGORY (&category) ke pemanggil
	// INGAT ! ini hanya mengirim alamat di memory nya saja.
	return &category, nil
}

func (repo *CategoryRepository) UpdateCategory(category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"

	result, err := repo.db.Exec(query, category.Name, category.Description, category.Id)
	if err != nil {
		return err
	}

	// cek row yang terpengaruh, ini mengembalikan nilai integer64 dan error
	rows, err := result.RowsAffected()

	// cek jika terjadi error
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Gagal melakukan update, kategori tidak ditemukan")
	}

	// walaupun tidak ada error, tetap harus mengembalikan nilai.
	return nil

}

func (repo *CategoryRepository) DeleteCategory(id int) error {
	query := "DELETE FROM categories WHERE id = $1"

	result, err := repo.db.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Gagal melakukan delete, kategori tidak ditemukan")
	}

	return nil
}
