package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HariIni(w http.ResponseWriter, r *http.Request) {
	// tangkap hasil dari Service GetReportHariIni
	// ini menghasilkan 2 variable, report dan error
	report, err := h.service.GetReportHariIni()

	// cek errornya disini
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// kirim response nya
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)

}
