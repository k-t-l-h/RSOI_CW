package main

import (
	"RSOI_CW/internal/pkg/flights/delivery"
	"RSOI_CW/internal/pkg/flights/repo"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func run() error {
	port, ok := os.LookupEnv("FLIGHTS_PORT")

	if !ok {
		port = "8030"
	}

	conn, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		return errors.New("no database url")
	}

	pool, err := pgxpool.Connect(context.Background(),
		conn)

	if err != nil {
		return err
	}

	rp := repo.NewFlightRepo(*pool)
	handler := delivery.NewFlightHandler(rp)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/flights", handler.AllFlights).Methods("GET")
	r.HandleFunc("/api/v1/flights", handler.AddFlight).Methods("POST")
	r.HandleFunc("/api/v1/flights/{UUID}", handler.UpdateFlight).Methods("PATCH")

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("FLIGHTS_ORIGINS"), " "),
		AllowCredentials: true,
	})

	corsHandler := c.Handler(r)
	srv := http.Server{Handler: corsHandler, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("flights running on: ", srv.Addr)
	return srv.ListenAndServe()
}
