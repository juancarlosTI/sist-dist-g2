package handlers

// import (
// 	"errors"
// 	"testing"

// 	"github.com/juancarlosTI/documentcase/documento/internal/application/auth"
// 	"github.com/juancarlosTI/documentcase/documento/internal/application/commands"
// )

// func TestCriarDocumentoHandler_Autorizado_CriaESalva(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()

// 	policy := FakeDocumentoPolicy{
// 		Permitir: true,
// 	}

// 	handler := NewCriarDocumentoHandler(repo, policy)

// 	cmd := commands.CriarDocumentoCommand{
// 		DocumentoID: "doc-1",
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

// func TestCriarDocumentoHandler_NaoAutorizado_NaoCriaDocumento(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()

// 	policy := FakeDocumentoPolicy{
// 		Permitir: false,
// 	}

// 	handler := NewCriarDocumentoHandler(repo, policy)

// 	cmd := commands.CriarDocumentoCommand{
// 		DocumentoID: "doc-1",
// 	}

// 	authCtx := auth.AuthContext{
// 		UserID: "cliente-1",
// 		Roles:  []string{"Usuario"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro de não autorizado")
// 	}

// 	if repo.Salvou {
// 		t.Fatal("documento não deveria ter sido salvo")
// 	}
// }

// func TestCriarDocumentoHandler_ErroAoSalvar_RetornaErro(t *testing.T) {
// 	repo := NewFakeDocumentoRepo()
// 	repo.ErrSalvar = errors.New("erro ao salvar")

// 	policy := FakeDocumentoPolicy{
// 		Permitir: true,
// 	}

// 	handler := NewCriarDocumentoHandler(repo, policy)

// 	cmd := commands.CriarDocumentoCommand{
// 		DocumentoID: "doc-1",
// 	}

// 	authCtx := auth.AuthContext{
// 		UserID: "cliente-1",
// 		Roles:  []string{"Backoffice"},
// 	}

// 	err := handler.Handle(cmd, authCtx)

// 	if err == nil {
// 		t.Fatal("esperava erro ao salvar documento")
// 	}
// }
