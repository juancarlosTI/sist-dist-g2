package controllers

// Arquivar documento
// Associar documento
// Atualizar documento
// Criacao de documento
// Invalidar documento

import (
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/interface/http/middlewares"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/commands"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/handlers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
)

type DocumentoController struct {
	storage ports.Storage

	arquivarDocumentoHandler  *handlers.ArquivarDocumentoHandler
	associarDocumentoHandler  *handlers.AssociarDocumentoAoProcessoHandler
	atualizarDocumentoHandler *handlers.AtualizarDocumentoHandler
	criacaoDeDocumentoHandler *handlers.CriarDocumentoHandler
	invalidarDocumentoHandler *handlers.InvalidarDocumentoHandler
}

func NewDocumentoController(
	storage ports.Storage,

	arquivarDocumentoHandler *handlers.ArquivarDocumentoHandler,
	associarDocumentoHandler *handlers.AssociarDocumentoAoProcessoHandler,
	atualizarDocumentoHandler *handlers.AtualizarDocumentoHandler,
	criacaoDeDocumentoHandler *handlers.CriarDocumentoHandler,
	invalidarDocumentoHandler *handlers.InvalidarDocumentoHandler,

) *DocumentoController {
	return &DocumentoController{
		storage: storage,

		arquivarDocumentoHandler:  arquivarDocumentoHandler,
		associarDocumentoHandler:  associarDocumentoHandler,
		atualizarDocumentoHandler: atualizarDocumentoHandler,
		criacaoDeDocumentoHandler: criacaoDeDocumentoHandler,
		invalidarDocumentoHandler: invalidarDocumentoHandler,
	}
}

func (c *DocumentoController) ArquivarDocumento(w http.ResponseWriter, r *http.Request) {
	// auth_ctx := mappers.AuthContextFromRequest(r)
	// ctx := r.Context()

	// processo_id, err := helpers.ProcessoID(r.PathValue("processo_id"))
	// if err != nil {
	// 	return
	// }

	// query := queries.BuscarProcessoPorIDQuery{
	// 	ProcessoID: processo_id,
	// }
	// // Iniciar repo????

	// result, err := c.consultarProcessoPorIDHandler.Handle(
	// 	ctx, query, auth_ctx,
	// )
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusForbidden)
	// 	return
	// }

	// // serialização omitida por brevidade
	// _ = result
	// w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusOK)

}

func (c *DocumentoController) AssociarDocumento(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func (c *DocumentoController) AtualizarDocumento(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

// Criar Documento godoc
// @Summary Criar documento
// @Tags Documento
// @Accept multipart/form-data
// @Produce json
// @Param arquivos formData file true "Arquivos para upload (múltiplos)"
// @Success 202 {string} string "accepted"
// @Failure 400 {string} string "invalid request"
// @Router /documento/add [post]
// @Security BearerAuth
// func (c *DocumentoController) CriarDocumento(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {

// 	ctx := r.Context()
// 	authCtx := middlewares.MustGetAuthContext(r)

// 	err := r.ParseMultipartForm(100 << 20)
// 	if err != nil {

// 		http.Error(
// 			w,
// 			"arquivo obrigatorio",
// 			http.StatusBadRequest,
// 		)

// 		return
// 	}

// 	files := r.MultipartForm.File["arquivos"]

// 	if len(files) == 0 {

// 		http.Error(
// 			w,
// 			"nenhum arquivo enviado",
// 			http.StatusBadRequest,
// 		)

// 		return
// 	}

// 	var uploadedFiles []ports.UploadedFile

// 	for _, header := range files {

// 		file, err := header.Open()

// 		if err != nil {

// 			http.Error(
// 				w,
// 				err.Error(),
// 				http.StatusInternalServerError,
// 			)

// 			return
// 		}

// 		uploadedFile, err :=
// 			c.storage.Upload(
// 				ctx,
// 				file,
// 				header.Filename,
// 				header.Header.Get("Content-Type"),
// 			)

// 		file.Close()

// 		if err != nil {

// 			http.Error(
// 				w,
// 				err.Error(),
// 				http.StatusInternalServerError,
// 			)

// 			return
// 		}

// 		uploadedFiles =
// 			append(
// 				uploadedFiles,
// 				*uploadedFile,
// 			)
// 	}

// 	var (
// 		wg      sync.WaitGroup
// 		errChan = make(chan error, len(uploadedFiles))
// 	)

// 	for _, uploadedFile := range uploadedFiles {

// 		wg.Add(1)

// 		go func(file ports.UploadedFile) {

// 			defer wg.Done()

// 			cmd := commands.CriarDocumentoCommand{
// 				ArquivoID: file.ID,
// 			}

// 			err :=
// 				c.criacaoDeDocumentoHandler.Handle(
// 					cmd,
// 					authCtx,
// 					ctx,
// 				)

// 			if err != nil {
// 				errChan <- err
// 			}

// 		}(uploadedFile)
// 	}

// 	wg.Wait()

// 	close(errChan)

// 	for err := range errChan {

// 		if err != nil {

// 			http.Error(
// 				w,
// 				err.Error(),
// 				http.StatusInternalServerError,
// 			)

// 			return
// 		}
// 	}

// 	w.WriteHeader(http.StatusAccepted)
// }

func (c *DocumentoController) CriarDocumento(
	w http.ResponseWriter,
	r *http.Request,
) {

	ctx := r.Context()
	authCtx := middlewares.MustGetAuthContext(r)

	err := r.ParseMultipartForm(100 << 20)
	if err != nil {

		http.Error(
			w,
			"arquivo obrigatorio",
			http.StatusBadRequest,
		)

		return
	}

	files := r.MultipartForm.File["arquivos"]

	if len(files) == 0 {

		http.Error(
			w,
			"nenhum arquivo enviado",
			http.StatusBadRequest,
		)

		return
	}

	for _, header := range files {

		file, err := header.Open()

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)

			return
		}

		uploadedFile, err :=
			c.storage.Upload(
				ctx,
				file,
				header.Filename,
				header.Header.Get("Content-Type"),
			)

		file.Close()

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)

			return
		}

		cmd := commands.CriarDocumentoCommand{
			ArquivoID: uploadedFile.ID,
		}

		err =
			c.criacaoDeDocumentoHandler.Handle(
				cmd,
				authCtx,
				ctx,
			)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)

			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
}

func (c *DocumentoController) InvalidarDocumento(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}
