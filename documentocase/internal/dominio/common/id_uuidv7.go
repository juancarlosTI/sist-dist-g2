package common

import "github.com/google/uuid"

type PedidoDocumento struct {
	value string
}

func NovoExternalPedidoDocumentoID(value string) (*PedidoDocumento, error) {
	if value == "" {
		return &PedidoDocumento{}, nil
	}

	return &PedidoDocumento{value: value}, nil
}

func (id PedidoDocumento) String() string {
	return id.value
}

func (e PedidoDocumento) IsEmpty() bool {
	return e.value == ""
}

type DocumentoID struct {
	value string
}

func NovoDocumentoID() (DocumentoID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return DocumentoID{}, err
	}
	return DocumentoID{value: id.String()}, nil
}

func NovoExternalDocumentoID(value string) (DocumentoID, error) {
	if value == "" {
		return DocumentoID{}, nil
	}

	return DocumentoID{value: value}, nil
}

func (id DocumentoID) String() string {
	return id.value
}

func (e DocumentoID) IsEmpty() bool {
	return e.value == ""
}
