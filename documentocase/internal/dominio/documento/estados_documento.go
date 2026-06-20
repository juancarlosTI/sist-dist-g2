package documento

type EstadoDocumento int

const (
	Ativo EstadoDocumento = iota
	Invalidado
	Arquivado
)
