package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	config "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/config"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/migrate"
)

func main() {

	cfg := config.Load()

	db, err := sql.Open("pgx", cfg.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := migrate.Run(db, "/app/migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Migration executada com sucesso.")
}
