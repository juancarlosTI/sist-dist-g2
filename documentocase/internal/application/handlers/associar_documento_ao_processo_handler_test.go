package handlers

// import (
// 	"errors"
// 	"testing"

// 	"github.com/juancarlosTI/documentcase/documento/internal/application/auth"
// 	"github.com/juancarlosTI/documentcase/documento/internal/application/commands"
// 	"github.com/juancarlosTI/documentcase/documento/internal/domain/documento"
// )

// func TestAssociarDocumentoAoProcessoHandler_DocumentoNaoExiste_RetornaErro(t *testing.T) {
// 	repo := NewFakeDocumentoRepo() // vazio

// 	policy := FakeDocumentoPolicy{Permitir: true}

// 	handler := NewAssociarDocumentoAoProcessoHandler(repo, policy)

// 	cmd := commands.AssociarDocumentoAoProcessoCommand{
// 		DocumentoID: "doc-inexistente",
// 		ProcessoID:  "proc-1",
// 	}

// 	authCtx := auth.AuthContext{
// 		Roles: []string{"Backoffice"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro ao buscar documento inexistente")
// 	}
// }

// func TestAssociarDocumentoAoProcessoHandler_NaoAutorizado_NaoSalva(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()
// 	doc := documento.CriarDocumento("doc-1")
// 	repo.Documentos["doc-1"] = doc

// 	policy := FakeDocumentoPolicy{Permitir: false}

// 	handler := NewAssociarDocumentoAoProcessoHandler(repo, policy)

// 	cmd := commands.AssociarDocumentoAoProcessoCommand{
// 		DocumentoID: "doc-1",
// 		ProcessoID:  "proc-1",
// 	}

// 	authCtx := auth.AuthContext{
// 		Roles: []string{"Usuario"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro de não autorizado")
// 	}

// 	if repo.Salvou {
// 		t.Fatal("não deveria salvar documento")
// 	}
// }

// func TestAssociarDocumentoAoProcessoHandler_Autorizado_AssociaESalva(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()
// 	doc := documento.CriarDocumento("doc-1")
// 	repo.Documentos["doc-1"] = doc

// 	policy := FakeDocumentoPolicy{Permitir: true}

// 	handler := NewAssociarDocumentoAoProcessoHandler(repo, policy)

// 	cmd := commands.AssociarDocumentoAoProcessoCommand{
// 		DocumentoID: "doc-1",
// 		ProcessoID:  "proc-1",
// 	}

// 	authCtx := auth.AuthContext{
// 		UserID: "cliente-1",
// 		Roles:  []string{"Backoffice"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err != nil {
// 		t.Fatalf("erro inesperado: %v", err)
// 	}

// 	if !repo.Salvou {
// 		t.Fatal("documento deveria ter sido salvo")
// 	}

// 	if _, ok := repo.Documentos["doc-1"]; !ok {
// 		t.Fatal("documento deveria existir no repositório")
// 	}

// }

// func TestAssociarDocumentoAoProcessoHandler_ErroAoSalvar_RetornaErro(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()
// 	doc := documento.CriarDocumento("doc-1")
// 	repo.Documentos["doc-1"] = doc
// 	repo.ErrSalvar = errors.New("erro ao salvar")

// 	policy := FakeDocumentoPolicy{Permitir: true}

// 	handler := NewAssociarDocumentoAoProcessoHandler(repo, policy)

// 	cmd := commands.AssociarDocumentoAoProcessoCommand{
// 		DocumentoID: "doc-1",
// 		ProcessoID:  "proc-1",
// 	}

// 	authCtx := auth.AuthContext{
// 		Roles: []string{"Backoffice"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro ao salvar")
// 	}
// }
