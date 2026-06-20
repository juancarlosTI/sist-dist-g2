package documento

type DocumentoSnapshot struct {
	ID            string
	Estado        int
	OrigemCanal   string
	OrigemSistema string
	AutorTipo     string
	AutorID       string
	Versao        int
	ProcessosIDs  []string
}
