package main

import (
	"RSOI_CW/internal/pkg/report/delivery"
	"RSOI_CW/internal/pkg/report/repo"
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
	port, ok := os.LookupEnv("REPORTS_PORT")

	if !ok {
		port = "8060"
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

	rp := repo.NewReportRepo(*pool)
	handler := delivery.NewReportHandler(rp)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/reports/flights", handler.CheckByUsers).Methods("GET")
	r.HandleFunc("/api/v1/reports/flights-filling", handler.CheckFilling).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("REPORT_ORIGINS"), " "),
		AllowCredentials: true,
	})

	corsHandler := c.Handler(r)
	srv := http.Server{Handler: corsHandler, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("report running on: ", srv.Addr)
	return srv.ListenAndServe()
}
