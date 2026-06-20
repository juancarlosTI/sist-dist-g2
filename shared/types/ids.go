package types

type EventoID string

func (id EventoID) String() string {
	return string(id)
}
