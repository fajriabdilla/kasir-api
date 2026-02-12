package repositories

import "database/sql"

type ReportRepository struct {
	db *sql.DB
}

func NewRepositoryReport(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetTodaySummary() (int, int, error) {
	// query untuk mendapatkan total revenue dan total transaksi
	// fungsi COALESCE, jika ekspresi pertama mengembalikan null,
	// 					maka mengembalikan nilai kedua, jika tidak null maka mengembalikan nilai pertama
	query := `
	SELECT
		COALESCE(SUM(total_amount), 0) AS total_revenue,
		COUNT(*) AS total_transaksi
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`

	// buat variable baru
	var revenue, transaksi int

	// assign hasil query ke variabel revenue dan transaksi
	// disini kita sudah mendapatkan total revenue dan total transaksi berdasarkan tanggal saat ini
	err := r.db.QueryRow(query).Scan(&revenue, &transaksi)
	if err != nil {
		return 0, 0, err
	}
	// kirim hasilnua disini
	return revenue, transaksi, err
}

func (r *ReportRepository) GetProdukTerlaris() (string, int, error) {
	// query untuk mendapatkan produk terlaris berdasarkan tanggal saat ini
	// ada 3 tabel yang saling berhubungan, yaitu transaction_details, product dan transaction
	// menghasilkan nama produk (name) dan quantity (qty)
	// kemudian menggunakan agregat tanggal hari ini. DATE(created_at) = CURRENT_DATE
	// diberikan limitasi hanya satu data saja.
	query := `
		SELECT p.name, SUM(td.quantity) as qty
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`

	// buat variable baru
	var name string
	var qty int

	// eksekusi query dan masukan ke variable name dan qty
	// disini kita sudah mendapatkan produk nama terlaris dan jumlahnya
	err := r.db.QueryRow(query).Scan(&name, &qty)
	return name, qty, err
}
