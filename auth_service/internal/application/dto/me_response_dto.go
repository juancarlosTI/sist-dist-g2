package dto

type MeResponseDTO struct {
	ID    string `json:"me_id"`
	Email string `json:"me_email"`
	Nome  string `json:"me_nome"`
	Roles string `json:"me_roles"`
}
