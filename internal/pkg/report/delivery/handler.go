package delivery

import (
	"net/http"
	"rsoi-kp-k-t-l-h/internal/pkg/middleware"
	"rsoi-kp-k-t-l-h/internal/pkg/report"
)

type ReportHandler struct {
	repo report.IRepo
}

func NewReportHandler(repo report.IRepo) *ReportHandler {
	return &ReportHandler{repo: repo}
}

func (h ReportHandler) CheckFilling(w http.ResponseWriter, r *http.Request) {
	status := h.repo.GetFilling()
	middleware.Response(w, status, nil)
}

func (h ReportHandler) CheckByUsers(w http.ResponseWriter, r *http.Request) {
	status := h.repo.CheckByUsers()
	middleware.Response(w, status, nil)
}
