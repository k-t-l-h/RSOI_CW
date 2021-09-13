package delivery

import (
	"RSOI_CW/internal/pkg/middleware"
	"RSOI_CW/internal/pkg/report"
	"net/http"
)

type ReportHandler struct {
	repo report.IRepo
}

func NewReportHandler(repo report.IRepo) *ReportHandler {
	return &ReportHandler{repo: repo}
}

func (h ReportHandler) CheckFilling(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8887")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")

	if r.Method == http.MethodOptions {
		return
	}
	body, status := h.repo.GetFilling()
	middleware.Response(w, status, body)
}

func (h ReportHandler) CheckByUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8887")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")

	if r.Method == http.MethodOptions {
		return
	}
	body, status := h.repo.CheckByUsers()
	middleware.Response(w, status, body)
}
