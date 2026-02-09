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
	revenue, transaksi, err := s.repo.GetTodaySummary()

	if err != nil {
		return nil, err
	}

	name, qty, err := s.repo.GetProdukTerlaris()

	if err != nil {
		return nil, err
	}

	report := &models.Report{
		TotalRevenue:   revenue,
		TotalTransaksi: transaksi,
		ProdukTerlaris: models.ProdukTerjual{
			Nama:       name,
			QtyTerjual: qty,
		},
	}

	return report, nil
}
