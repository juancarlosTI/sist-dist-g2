package queries

type DocumentoReadDTO struct {
	ID         string
	Estado     string
	ProcessoID string
	Valido     bool
}

type DocumentoResumoReadDTO struct {
	ID     string
	Estado string
}

type DocumentosDoUsuarioReadDTO struct {
	AutorID    string
	Documentos []DocumentoResumoReadDTO
}

type DetalheDocumentoReadDTO struct {
	ID         string
	Estado     string
	ProcessoID string
	Historico  []string
}
