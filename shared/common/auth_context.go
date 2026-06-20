package common

type AuthContext struct {
	UserID    string
	ActorType string
	Roles     []string
}
