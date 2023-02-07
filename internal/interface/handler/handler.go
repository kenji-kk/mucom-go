package handler

type RootHandlers interface {
	AuthHandler
	ReviewHandler
}

type rootHandlers struct {
	AuthHandler
	ReviewHandler
}

func NewRootHandlers(authHandler AuthHandler, reviewHandler ReviewHandler) RootHandlers {
	return &rootHandlers{AuthHandler: authHandler, ReviewHandler: reviewHandler}
}
