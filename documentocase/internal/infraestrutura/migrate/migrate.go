package migrate

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(db *sql.DB, path string) error {

	// Cria um driver de migrate
	driver, err := postgres.WithInstance(db,
		&postgres.Config{
			DatabaseName: os.Getenv("DB_NAME"),
		})
	if err != nil {
		log.Fatalf("erro criando driver migrate: %v", err)
	}

	// Inicializa o migrator
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("erro inicializando migrate: %v", err)
	}

	// 🔎 Verifica versão atual
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("erro verificando versão: %v", err)
	}

	if dirty {
		log.Fatalf("Banco está em estado dirty na versão %d. Corrija antes de continuar.", version)
	}

	if err == migrate.ErrNilVersion {
		log.Println("Nenhuma migration aplicada ainda.")
	} else {
		log.Printf("Versão atual do banco: %d\n", version)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Banco já está atualizado.")
			return nil
		}
		log.Fatalf("erro aplicando migrations: %v", err)
	}

	return nil
}
