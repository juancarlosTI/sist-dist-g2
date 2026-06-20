package sql

import (
	"database/sql"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/queries"
)

type DocumentoReadSQLRepository struct {
	db *sql.DB
}

func NewDocumentoReadSQLRepositoryHandler(
	db *sql.DB,
) *DocumentoReadSQLRepository {
	return &DocumentoReadSQLRepository{
		db: db,
	}
}

func (r *DocumentoReadSQLRepository) PorID(documentoID string) (*queries.DocumentoReadDTO, error) {
	return &queries.DocumentoReadDTO{}, nil
}

func (r *DocumentoReadSQLRepository) ListarPorUsuario(autorID string) (queries.DocumentosDoUsuarioReadDTO, error) {
	return queries.DocumentosDoUsuarioReadDTO{}, nil
}
