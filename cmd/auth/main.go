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
	"rsoi-kp-k-t-l-h/internal/pkg/auth/delivery"
	"rsoi-kp-k-t-l-h/internal/pkg/auth/repo"
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
	port, ok := os.LookupEnv("AUTH_PORT")

	if !ok {
		port = "8000"
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

	repo := repo.NewAuthRepo(*pool)
	handler := delivery.NewAuthHandler(repo)

	r := mux.NewRouter()

	//Авторизация.
	//header: Authorization: basic(<login>:<password>)
	//POST /auth -> JWT token
	r.HandleFunc("/api/v1/auth",
		handler.GetToken).Methods("POST")

	//Проверка токена пользователя.
	//header: Authorization: bearer <jwt>
	//POST /verify
	r.HandleFunc("/api/v1/verify",
		handler.CheckToken).Methods("POST")

	r.HandleFunc("/api/v1/admin",
		handler.CheckAdminToken).Methods("POST")

	//Список всех пользователей. [A][G]
	//GET /users
	r.HandleFunc("/api/v1/users",
		handler.GetAllUsers).Methods("GET")

	//Добавление нового пользователя. [A][G]
	//POST /users
	r.HandleFunc("/api/v1/users",
		handler.AddUser).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(os.Getenv("AUTH_ORIGINS"), " "),
		AllowCredentials: true,
	})

	corsHandler := c.Handler(r)
	srv := http.Server{Handler: corsHandler, Addr: fmt.Sprintf(":%s", port)}
	http.Handle("/", r)

	log.Print("auth running on ", srv.Addr)
	return srv.ListenAndServe()
}
