package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"rsoi-kp-k-t-l-h/internal/pkg/airport/delivery"
	"rsoi-kp-k-t-l-h/internal/pkg/airport/repo"
	"strings"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func run() error {
	port, ok := os.LookupEnv("AIRPORT_PORT")

	if !ok {
		port = "8020"
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

	repo := repo.NewAirportRepo(*pool)
	handler := delivery.NewAirportHandler(repo)

	r := mux.NewRouter()

	//GET /airports/{airportUid}
	r.HandleFunc("/api/v1/airports/{UUID}",
		handler.GetAirport).Methods("GET")
	//GET /airports
	r.HandleFunc("/api/v1/airports",
		handler.GetAirports).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("AIRPORTS_ORIGINS"), " "),
		AllowCredentials: true,
	})

	corsHandler := c.Handler(r)
	srv := http.Server{Handler: corsHandler, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("airport running on: ", srv.Addr)
	return srv.ListenAndServe()
}
