package documento

// func TestDocumentoCriadoComecaEmEstadoValido(t *testing.T) {
// 	doc := CriarDocumento("doc-1")

// 	if doc.Estado() != Ativo {
// 		t.Fatalf("documento deveria iniciar como Ativo")
// 	}
// }

// func TestDocumentoArquivadoNaoPodeSerAtualizado(t *testing.T) {
// 	doc := CriarDocumento("doc-1")
// 	doc.Arquivar()

// 	err := doc.AtualizarConteudo("novo conteudo")

// 	if err == nil {
// 		t.Fatalf("documento arquivado não deveria aceitar atualização")
// 	}
// }

// func TestDocumentoInvalidadoNaoPodeSerAssociadoAoProcesso(t *testing.T) {
// 	doc := CriarDocumento("doc-1")
// 	doc.Invalidar("expirado")

// 	err := doc.AssociarAoProcesso("processo-123")

// 	if err == nil {
// 		t.Fatalf("documento inválido não pode ser associado a processo")
// 	}
// }

// func TestDocumentoArquivadoNaoPodeSerInvalidado(t *testing.T) {
// 	doc := CriarDocumento("doc-1")
// 	doc.Arquivar()

// 	err := doc.Invalidar("qualquer")

// 	if err == nil {
// 		t.Fatalf("documento arquivado não pode ser invalidado")
// 	}
// }
