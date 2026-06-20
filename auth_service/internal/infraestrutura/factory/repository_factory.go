package factory

import (
	"database/sql"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/persistencia/postgres"
)

type RepositoryFactory struct {
	db *sql.DB
}

func NewRepositoryFactory(db *sql.DB) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

// SEM TRANSAÇÃO
func (f *RepositoryFactory) NewUserRepository() identidade.UserRepository {
	return postgres.NewUserRepositorySQL(f.db)
}

// COM TRANSAÇÃO
func (f *RepositoryFactory) NewUserRepositoryTx(tx *sql.Tx) identidade.UserRepository {
	return postgres.NewUserRepositorySQL(tx)
}

func (f *RepositoryFactory) NewRefreshTokenRepositoryTx(tx *sql.Tx) service.RefreshTokenRepository {
	return postgres.NewRefreshTokenRepository(tx)
}
