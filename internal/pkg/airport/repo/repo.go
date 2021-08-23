package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/pgxpool"
	"log"
	"rsoi-kp-k-t-l-h/internal/models"
)

type AirportRepo struct {
	pool pgxpool.Pool
}

func NewAirportRepo(pool pgxpool.Pool) *AirportRepo {
	return &AirportRepo{pool: pool}
}

func (r *AirportRepo) SelectAirports() ([]models.Airport, int) {
	SelectAll := "SELECT \"AirportUUID\", \"AirportName\" FROM public.airport;"

	row, err := r.pool.Query(context.Background(), SelectAll)
	if err != nil {
		return []models.Airport{}, models.StatusError
	}

	result := []models.Airport{}
	for row.Next() {
		airport := models.Airport{}
		if err := row.Scan(&airport.AirportUUID, &airport.Name); err != nil {
			log.Print(err)
			return []models.Airport{}, models.StatusError
		}
		result = append(result, airport)
	}
	return result, models.StatusOkey
}

func (r *AirportRepo) SelectAirport(uuid uuid.UUID) (models.Airport, int) {
	SelectOne := "SELECT  \"AirportUUID\", \"AirportName\" FROM public.airport" +
		" WHERE \"AirportUUID\" = $1;"

	row := r.pool.QueryRow(context.Background(), SelectOne, uuid)

	airport := models.Airport{}

	if err := row.Scan(&airport.AirportUUID, &airport.Name); err != nil {
		return models.Airport{}, models.StatusError
	}
	return airport, models.StatusOkey
}
