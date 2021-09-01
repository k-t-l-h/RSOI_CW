package report

import "RSOI_CW/internal/models"

type IRepo interface {
	AddStat(report models.Report) int
	GetFilling() ([]models.ReportFilling, int)
	CheckByUsers() ([]models.ReportUsers, int)
}
