package main

import (
	"database/sql"
	"log"
	stdhttp "net/http"
	"os"
	"os/signal"
	"personia/infra/http"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

func migrateDB(db *sql.DB) {
	q := `CREATE TABLE employee (name text, supervisor text)`
	if _, err := db.Exec(q); err != nil {
		log.Fatalf("migrate db: %v", err)
	}
}

func getenv(env, fallback string) string {
	v := os.Getenv(env)
	if v == "" {
		return fallback
	}
	return v
}

func main() {
	addr := getenv("PERSONIA_ADDR", ":3000")
	apiSecret := getenv("PERSONIA_API_SECRET", "secret")

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	migrateDB(db)

	server := http.NewServer(http.ServerOpts{
		DB:        db,
		APISecret: apiSecret,
	})

	go func() {
		err := server.Start(addr)
		if err != nil && err != stdhttp.ErrServerClosed {
			log.Fatalf("start server failed: %v", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
}
