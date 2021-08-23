package report

import "rsoi-kp-k-t-l-h/internal/models"

type IRepo interface {
	AddStat(report models.Report) int
	GetFilling() ([]models.ReportFilling, int)
	CheckByUsers() ([]models.ReportUsers, int)
}
