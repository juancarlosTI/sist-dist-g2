package handlers

// import (
// 	"errors"
// 	"testing"

// 	"github.com/juancarlosTI/documentcase/documento/internal/application/auth"
// 	"github.com/juancarlosTI/documentcase/documento/internal/application/commands"
// 	"github.com/juancarlosTI/documentcase/documento/internal/domain/documento"
// )

// func TestInvalidarDocumentoHandler_DocumentoNaoExiste_RetornaErro(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()

// 	policy := FakeDocumentoPolicy{Permitir: true}

// 	handler := NewInvalidarDocumentoHandler(repo, policy)

// 	cmd := commands.InvalidarDocumentoCommand{
// 		DocumentoID: "doc-inexistente",
// 		Motivo:      "Documento inválido",
// 	}

// 	authCtx := auth.AuthContext{
// 		Roles: []string{"Backoffice"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro ao buscar documento inexistente")
// 	}
// }

// func TestInvalidarDocumentoHandler_NaoAutorizado_NaoSalva(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()
// 	doc := documento.CriarDocumento("doc-1")
// 	repo.Documentos["doc-1"] = doc

// 	policy := FakeDocumentoPolicy{Permitir: false}

// 	handler := NewInvalidarDocumentoHandler(repo, policy)

// 	cmd := commands.InvalidarDocumentoCommand{
// 		DocumentoID: "doc-1",
// 		Motivo:      "Documento inválido",
// 	}

// 	authCtx := auth.AuthContext{
// 		Roles: []string{"Usuario"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro de não autorizado")
// 	}

// 	if repo.Salvou {
// 		t.Fatal("documento não deveria ter sido salvo")
// 	}
// }

// func TestInvalidarDocumentoHandler_Autorizado_InvalidaESalva(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()
// 	doc := documento.CriarDocumento("doc-1")
// 	repo.Documentos["doc-1"] = doc

// 	policy := FakeDocumentoPolicy{Permitir: true}

// 	handler := NewInvalidarDocumentoHandler(repo, policy)

// 	cmd := commands.InvalidarDocumentoCommand{
// 		DocumentoID: "doc-1",
// 		Motivo:      "Documento expirado",
// 	}

// 	authCtx := auth.AuthContext{
// 		Roles: []string{"Backoffice"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err != nil {
// 		t.Fatalf("erro inesperado: %v", err)
// 	}

// 	if !repo.Salvou {
// 		t.Fatal("documento deveria ter sido salvo")
// 	}

// 	if !doc.EstaInvalidado() {
// 		t.Fatal("documento deveria estar invalidado")
// 	}
// }

// func TestInvalidarDocumentoHandler_ErroAoSalvar_RetornaErro(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()
// 	doc := documento.CriarDocumento("doc-1")
// 	repo.Documentos["doc-1"] = doc
// 	repo.ErrSalvar = errors.New("erro ao salvar")

// 	policy := FakeDocumentoPolicy{Permitir: true}

// 	handler := NewInvalidarDocumentoHandler(repo, policy)

// 	cmd := commands.InvalidarDocumentoCommand{
// 		DocumentoID: "doc-1",
// 		Motivo:      "Documento inválido",
// 	}

// 	authCtx := auth.AuthContext{
// 		Roles: []string{"Backoffice"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro ao salvar documento")
// 	}
// }
