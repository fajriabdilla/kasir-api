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
	query := `
	SELECT
		COALESCE(SUM(total_amount), 0) AS total_revenue,
		COUNT(*) AS total_transaksi
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
`
	var revenue, transaksi int

	// assign hasil query ke variabel revenue dan transaksi
	err := r.db.QueryRow(query).Scan(&revenue, &transaksi)
	if err != nil {
		return 0, 0, err
	}

	return revenue, transaksi, err
}

func (r *ReportRepository) GetProdukTerlaris() (string, int, error) {
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
	var name string
	var qty int

	err := r.db.QueryRow(query).Scan(&name, &qty)
	return name, qty, err
}
