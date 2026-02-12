package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewServiceReport(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReportHariIni() (*models.Report, error) {
	// tahap 1. dapatkan ringakasan transaksi hari ini total revenue dan total transaksi
	// Repository GetTodaySummary menghasilkan 3 variabel (revenue, transaksi, error)
	// ini adalah total revenue dan total transaksi, error jika ada
	revenue, transaksi, err := s.repo.GetTodaySummary()

	// cek jika ada errornya
	if err != nil {
		return nil, err
	}

	// tahap 2. dapatkan produk terlaris
	// Repository GetProdukTerlaris mengembalikan 3 variabel (name, qty, err)
	// ini adalah nama produk dan quantity nya, error jika ada
	name, qty, err := s.repo.GetProdukTerlaris()

	// cek jika ada errornya
	if err != nil {
		return nil, err
	}

	// assign nilai tahap 1 dan tahap 2 yang sudah didapatkan tadi ke dalam struct Report yang sudah didefinisikan di models/report.go
	report := &models.Report{
		TotalRevenue:   revenue,
		TotalTransaksi: transaksi,
		// ProdukTerlaris adalah tipenya struct ProdukTerjual
		ProdukTerlaris: models.ProdukTerjual{
			Nama:       name,
			QtyTerjual: qty,
		},
	}

	// kirim hasilnya disini.
	return report, nil
}
