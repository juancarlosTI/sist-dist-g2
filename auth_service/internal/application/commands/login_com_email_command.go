package commands

type LoginComEmailCommand struct {
	Email    string `json:"email" example:"user@email.com"`
	Password string `json:"password" example:"123456"`
}
