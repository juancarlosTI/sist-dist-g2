package policies

// import (
// 	"testing"

// 	"github.com/juancarlosTI/documentcase/documento/internal/application/auth"
// 	"github.com/juancarlosTI/documentcase/documento/internal/domain/documento"
// )

// func TestPolicy_Backoffice_PodeArquivarDocumento(t *testing.T) {
// 	policy := DocumentoPolicyImpl{}

// 	auth := auth.AuthContext{
// 		ActorType: "Humano",
// 		Roles:     []string{"Backoffice"},
// 	}

// 	doc := documento.CriarDocumento("doc-1")

// 	if !policy.PodeArquivarDocumento(auth) {
// 		t.Fatalf("backoffice deveria poder arquivar documento")
// 	}
// }

// func TestPolicy_UsuarioComum_NaoPodeArquivarDocumento(t *testing.T) {
// 	policy := DocumentoPolicyImpl{}

// 	auth := auth.AuthContext{
// 		ActorType: "Humano",
// 		Roles:     []string{"Usuario"},
// 	}

// 	doc := documento.CriarDocumento("doc-1")

// 	if policy.PodeArquivarDocumento(auth) {
// 		t.Fatalf("usuario comum não deveria poder arquivar documento")
// 	}
// }

// func TestPolicy_NaoAutenticado_NaoPodeArquivarDocumento(t *testing.T) {
// 	policy := DocumentoPolicyImpl{}

// 	auth := auth.AuthContext{
// 		ActorType: "Humano",
// 		Roles:     []string{"Usuario"},
// 	}

// 	doc := documento.CriarDocumento("doc-1")

// 	if policy.PodeArquivarDocumento(auth) {
// 		t.Fatalf("usuario nao autenticado nao deveria poder arquivar")
// 	}
// }

// func TestPolicy_Sistema_PodeArquivarDocumento(t *testing.T) {
// 	policy := DocumentoPolicyImpl{}

// 	auth := auth.AuthContext{
// 		ActorType: "Sistema",
// 		Roles:     []string{"Backoffice"},
// 	}

// 	doc := documento.CriarDocumento("doc-1")

// 	if !policy.PodeArquivarDocumento(auth) {
// 		t.Fatalf("sistema deveria poder arquivar documento")
// 	}
// }
