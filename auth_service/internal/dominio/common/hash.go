package common

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func GerarHashEvento(base EventoBase, payload any) (string, error) {

	baseSemHash := base
	baseSemHash.Hash = ""

	envelope := struct {
		Base    EventoBase `json:"base"`
		Payload any        `json:"payload"`
	}{
		Base:    baseSemHash,
		Payload: payload,
	}

	bytes, err := json.Marshal(envelope)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:]), nil
}
