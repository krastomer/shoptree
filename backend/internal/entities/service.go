package entities

type AuthService interface {
	LoginCustomer(string, string) (string, error)
	// LogoutCustomer(string, string)
}
