package common

import "encoding/json"

func (d DocumentoID) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

func (d *DocumentoID) UnmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.value = value

	return nil
}

func (d PedidoDocumento) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

func (d *PedidoDocumento) UnmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.value = value

	return nil
}
