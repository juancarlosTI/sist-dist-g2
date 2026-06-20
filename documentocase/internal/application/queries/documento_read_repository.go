package queries

type DocumentoReadSQLRepository interface {
	PorID(id string) (*DocumentoReadDTO, error)
	ListarPorUsuario(autorID string) (DocumentosDoUsuarioReadDTO, error)
}
