package handler

type RootHandlers interface {
	AuthHandler
}

type rootHandlers struct {
	AuthHandler
}

func NewRootHandlers(authHandler AuthHandler) RootHandlers {
	return &rootHandlers{AuthHandler: authHandler}
}
