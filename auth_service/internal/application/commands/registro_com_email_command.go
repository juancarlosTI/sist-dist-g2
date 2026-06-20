package commands

type RegistroComEmailCommand struct {
	Nome        string `json:"nome" validate:"required,min=3"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	AccountType string `json:"account_type" example:"PROFISSIONAL"`
}
