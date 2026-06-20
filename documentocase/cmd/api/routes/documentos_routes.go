package routes

import (
	"github.com/go-chi/chi/v5"
	profissional "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/interface/http/controllers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/interface/http/middlewares"
)

func BuildDocumentoControllers(deps Dependencies) *profissional.DocumentoController {

	return profissional.NewDocumentoController(
		deps.Storage,
		deps.ArquivarDocumentoHandler,
		deps.AssociarDocumentoHandler,
		deps.AtualizarDocumentoHandler,
		deps.CriacaoDeDocumentoHandler,
		deps.InvalidarDocumentoHandler,
	)
}

func RegisterDocumentoRoutes(r chi.Router, c APIControllers) {
	controller := c.DocumentoController

	r.Route("/documento", func(r chi.Router) {
		r.Use(
			middlewares.AuthMiddleware,
		)

		// Queries
		// r.Get("/search_process/{processo_id}", controller.ConsultarProcesso)
		// r.Get("/all_clients/{cliente_id}", controller.ListarProcessosPorCliente)
		// r.Get("/profissional/{autor_id}", controller.ListarPorResponsavel)

		// Commands
		r.Post("/add", controller.CriarDocumento)

		// r.Post("/{processo_id}/conclusao",
		// 	controller.ConcluirProcesso,
		// )

		// r.Post("/{processo_id}/conclusao-tecnica",
		// 	controller.ConcluirProcessoTecnico,
		// )

		// // Pendências
		// r.Post("/{processo_id}/pendencias",
		// 	controller.RegistrarPendencia,
		// )

		// r.Post("/{processo_id}/pendencias/{pendencia_id}/resolver",
		// 	controller.ResolverPendencia,
		// )

		// // Associações
		// r.Post("/{processo_id}/clientes/{cliente_id}",
		// 	controller.AssociarCliente,
		// )

		// r.Post("/{processo_id}/advogados/{advogado_id}",
		// 	controller.AssociarAdvogado,
		// )

		// r.Post("/{processo_id}/documentos/{documento_id}",
		// 	controller.AssociarDocumento,
		// )

	})
}
