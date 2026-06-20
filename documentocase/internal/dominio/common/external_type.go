package common

type ExternalProcessoID struct {
	value string
}

func NovoExternalProcessoID(value string) (ExternalProcessoID, error) {
	if value == "" {
		return ExternalProcessoID{}, nil
	}

	return ExternalProcessoID{value: value}, nil
}

func (id ExternalProcessoID) String() *string {
	return &id.value
}

func (e ExternalProcessoID) IsEmpty() bool {
	return e.value == ""
}
