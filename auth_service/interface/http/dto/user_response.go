package dto

type UserResponseDTO struct {
	ID    string  `json:"id" example:"user_123"`
	Email string  `json:"email" example:"user@email.com"`
	Nome  *string `json:"nome,omitempty" example:"João Silva"`
	Roles string  `json:"roles" example:"admin,user"`
}
