package documento

// import (
// 	shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/eventos"
// )

// const DocumentoCriadoProcessorVersao = 1

// type DocumentoCriadoProcessorPayload struct {
// 	DocumentoID string `json:"documento_id"`
// 	ArquivoID   string `json:"arquivo_id"`
// }

// type DocumentoCriadoProcessor struct {
// 	BaseEvento    shared_context.EventoIntegracaoBase
// 	PayloadEvento DocumentoCriadoProcessorPayload
// }

// func (e DocumentoCriadoProcessor) Nome() string {
// 	return "ProcessoCriadoIndex"
// }

// func (e DocumentoCriadoProcessor) Base() shared_context.EventoIntegracaoBase {
// 	return e.BaseEvento
// }

// func (e DocumentoCriadoProcessor) Payload() DocumentoCriadoProcessorPayload {
// 	return e.PayloadEvento
// }
