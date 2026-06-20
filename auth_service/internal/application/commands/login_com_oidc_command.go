package commands

type LoginComOIDCCommand struct {
	Provider    string
	Code        string
	RedirectURI string
}
