package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/yaderv/medusario/internal/data"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dns          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  int
	}
}

type application struct {
	config config
	logger *log.Logger
	models *data.Models
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	var cfg config
	err := loadConfig(&cfg)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Println("db connection established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server.
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dns)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(time.Duration(cfg.db.maxIdleTime) * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func loadConfig(cfg *config) error {
	var err error
	cfg.db.dns = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	cfg.db.maxOpenConns, err = strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	if err != nil {
		return err
	}
	cfg.db.maxIdleConns, err = strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		return err
	}
	cfg.db.maxIdleTime, err = strconv.Atoi(os.Getenv("MAX_IDLE_TIME"))
	if err != nil {
		return err
	}
	cfg.port, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return err
	}
	cfg.env = os.Getenv("ENV")
	return nil
}
