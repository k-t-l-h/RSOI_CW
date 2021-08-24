package repo

import (
	"RSOI_CW/internal/models"
	"context"
	"github.com/google/uuid"
	_ "github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FlightRepo struct {
	pool pgxpool.Pool
}

func NewFlightRepo(pool pgxpool.Pool) *FlightRepo {
	return &FlightRepo{pool: pool}
}

func (r *FlightRepo) CreateFlight(flight models.Flight) int {

	Create := "INSERT INTO public.flight(\"From\", \"To\", \"Date\", \"FlightID\") " +
		"VALUES ($1, $2, $3, $4);"

	tag, err := r.pool.Exec(context.Background(), Create, flight.From, flight.To, flight.Date, uuid.New())
	if err != nil {
		return models.StatusError
	}
	if tag.RowsAffected() != 1 {
		return models.StatusConflict
	}
	return models.StatusOkey
}

func (r *FlightRepo) ReadFlight(FlightUUID uuid.UUID) (models.Flight, int) {

	SelectOne := "SELECT \"From\", \"To\", \"Date\"  FROM public.flight " +
		" WHERE \"FlightID\" = $1;"

	row := r.pool.QueryRow(context.Background(), SelectOne, FlightUUID)

	flight := models.Flight{}

	if err := row.Scan(&flight.From, &flight.To, &flight.Date); err != nil {
		return models.Flight{}, models.StatusError
	}

	return models.Flight{}, models.StatusOkey
}

func (r *FlightRepo) ReadFlights() ([]models.Flight, int) {

	SelectAll := "SELECT \"From\", \"To\", \"Date\"  FROM public.flight;"

	row, err := r.pool.Query(context.Background(), SelectAll)
	if err != nil {
		return nil, models.StatusError
	}

	result := []models.Flight{}
	for row.Next() {
		flight := models.Flight{}

		if err := row.Scan(&flight.From, &flight.To, &flight.Date); err != nil {
			return nil, models.StatusError
		}
		result = append(result, flight)
	}

	return []models.Flight{}, models.StatusOkey
}

func (r *FlightRepo) UpdateFlight(FlightUUID uuid.UUID, flight models.Flight) int {
	UpdateFL := "UPDATE public.flight " +
		"SET \"Date\"=$1 WHERE \"FlightID\"=$2;"

	tag, err := r.pool.Exec(context.Background(), UpdateFL, flight.Date, FlightUUID)
	if err != nil {
		return models.StatusError
	}

	if tag.RowsAffected() == 0 {
		return models.StatusNotFound
	}

	return models.StatusOkey
}
