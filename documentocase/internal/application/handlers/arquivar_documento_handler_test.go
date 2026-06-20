package handlers

// import (
// 	"testing"

// 	"github.com/juancarlosTI/documentcase/documento/internal/application/auth"
// 	"github.com/juancarlosTI/documentcase/documento/internal/application/commands"
// 	"github.com/juancarlosTI/documentcase/documento/internal/domain/documento"
// )

// func TestArquivarDocumentoHandler_Autorizado_ArquivaESalva(t *testing.T) {
// 	doc := documento.CriarDocumento("doc-1")

// 	repo := NewFakeDocumentoRepo()
// 	repo.Documentos["doc-1"] = doc

// 	policy := FakeDocumentoPolicy{Permitir: true}

// 	handler := NewArquivarDocumentoHandler(repo, policy)

// 	cmd := commands.ArquivarDocumentoCommand{
// 		DocumentoID: "doc-1",
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

// 	if !doc.EstaArquivado() {
// 		t.Fatal("documento deveria estar arquivado")
// 	}
// }
