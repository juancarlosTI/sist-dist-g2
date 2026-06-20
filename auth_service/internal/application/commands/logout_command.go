package commands

type LogoutCommand struct {
	RefreshToken string `json:"refresh_token"`
	Password     string
}
