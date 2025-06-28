package main

import (
	"log"

	"github.com/sikozonpc/social/internal/db"
	"github.com/sikozonpc/social/internal/env"
	"github.com/sikozonpc/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgresql://admin:adminpassword@localhost:54321/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
