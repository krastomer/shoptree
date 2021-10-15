package entities

type AuthService interface {
	Login(string, string) (string, error)
}

type ProfileService interface {
}
