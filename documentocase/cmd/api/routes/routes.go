package routes

import (
	"github.com/go-chi/chi/v5"
	documento "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/interface/http/controllers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/handlers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
)

type APIControllers struct {
	DocumentoController *documento.DocumentoController
}

type Dependencies struct {
	Storage ports.Storage
	// DocumentoCriacaoSolicitada *handlers.DocumentoCriacaoSolicitadaConsumer
	// BuscarProcessoHandler     *queries.BuscarDocumentoPorIDHandler
	// ListarDocumentosDoUsuario *queries.ListarDocumentosDoUsuarioPorIDHandler
	ArquivarDocumentoHandler  *handlers.ArquivarDocumentoHandler
	AssociarDocumentoHandler  *handlers.AssociarDocumentoAoProcessoHandler
	AtualizarDocumentoHandler *handlers.AtualizarDocumentoHandler
	CriacaoDeDocumentoHandler *handlers.CriarDocumentoHandler
	InvalidarDocumentoHandler *handlers.InvalidarDocumentoHandler
}

func RegisterRoutes(r chi.Router, c APIControllers) {
	RegisterDocumentoRoutes(r, c)
}

func BuildControllers(deps Dependencies) APIControllers {
	return APIControllers{
		DocumentoController: BuildDocumentoControllers(deps),
	}
}
