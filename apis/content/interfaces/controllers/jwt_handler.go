package controllers

type JWTHandler interface {
	Verify(string) (int, error)
}
