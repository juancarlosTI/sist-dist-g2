package queries

type BuscarDocumentoPorIDHandler struct {
	repo DocumentoReadSQLRepository
}

func NewBuscarDocumentoPorIDHandler(
	repo DocumentoReadSQLRepository,
) *BuscarDocumentoPorIDHandler {
	return &BuscarDocumentoPorIDHandler{
		repo: repo,
	}
}

func (h *BuscarDocumentoPorIDHandler) Handle(q BuscarDocumentoPorIDQuery) (*DocumentoReadDTO, error) {
	return h.repo.PorID(q.DocumentoID)
}

type ListarDocumentosDoUsuarioPorIDHandler struct {
	repo DocumentoReadSQLRepository
}

func NewListarDocumentosDoUsuarioPorIDHandler(
	repo DocumentoReadSQLRepository,
) *ListarDocumentosDoUsuarioPorIDHandler {
	return &ListarDocumentosDoUsuarioPorIDHandler{
		repo: repo,
	}
}

func (h *ListarDocumentosDoUsuarioPorIDHandler) Handle(
	q ListarDocumentosDoUsuarioPorIDQuery,
) (DocumentosDoUsuarioReadDTO, error) {
	return h.repo.ListarPorUsuario(q.AutorID)
}
