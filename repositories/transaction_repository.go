package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	// mulai database transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}

	// lakukan rollback jika terjadi error
	defer tx.Rollback()

	// definisikan nilai awal
	totalAmount := 0

	// buat sebuah slice kosong
	details := make([]models.TransactionDetail, 0)

	// lakukan iterasi pada items yang dikirim
	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductId).Scan(&productName, &productPrice, &stock)

		if err == sql.ErrNoRows {
			// return ini menggunakan placeholder di golang, %d adalah integer.
			return nil, fmt.Errorf("produk dengan id %d not found", item.ProductId)
		}

		if err != nil {
			return nil, err
		}

		// hitung harga produk dikali quantity
		subtotal := productPrice * item.Quantity

		// total amount = total amount + subtotal
		totalAmount += subtotal

		// update jumlah stock product yang di checkout
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductId)

		if err != nil {
			return nil, err
		}

		// tambahkan detail transaksi ke dalam struct
		details = append(details, models.TransactionDetail{
			ProductId:   item.ProductId,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})

	}

	// buat sebuah variabel untuk transaction id
	var transactionId int
	var transactionDetailId int

	// masukan ke database total amountnya dan tangkap id nya
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionId)

	if err != nil {
		return nil, err
	}

	// masukan semua item ke transaction detail
	for index := range details {
		details[index].TransactionId = transactionId
		err = tx.QueryRow("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id", transactionId, details[index].ProductId, details[index].Quantity, details[index].Subtotal).Scan(&transactionDetailId)

		if err != nil {
			return nil, err
		}

		details[index].Id = transactionDetailId
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		Id:          transactionId,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
