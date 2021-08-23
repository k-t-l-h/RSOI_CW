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
	"rsoi-kp-k-t-l-h/internal/pkg/bonus/delivery"
	"rsoi-kp-k-t-l-h/internal/pkg/bonus/repo"
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
	port, ok := os.LookupEnv("BONUS_PORT")

	if !ok {
		port = "8050"
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

	rp := repo.NewBonusRepo(*pool)
	handler := delivery.NewBonusHandler(rp)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/miles", handler.GetBonus).Methods("GET")
	r.HandleFunc("/api/v1/miles", handler.AddBonus).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("BONUS_ORIGINS"), " "),
		AllowCredentials: true,
	})

	corsHandler := c.Handler(r)
	srv := http.Server{Handler: corsHandler, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("bonus running on: ", srv.Addr)
	return srv.ListenAndServe()
}
