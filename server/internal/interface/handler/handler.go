package handler

type Handlers interface {

}

type handlers struct {
	authHandler AuthHandler
}

func NewHandlers (authHandler AuthHandler) Handlers{
	return &handlers{authHandler}
}

