package rest

import "github.com/labstack/echo/v4"

// import "github.com/mohammadne/bookman/library/internal/books/usecase"

// type server struct {
// 	usecase usecase.Usecase
// }

// func New(usecase usecase.Usecase) *server {
// 	return &server{usecase: usecase}
// }

// func (rest *restServer) setupRoutes() {
// 	rest.instance.POST("/auth/metrics", echo.WrapHandler(promhttp.Handler()))
// 	rest.instance.POST("/auth/sign_up", rest.signUp)
// 	rest.instance.POST("/auth/sign_in", rest.signIn)
// 	rest.instance.POST("/auth/sign_out", rest.signOut)
// 	rest.instance.POST("/auth/refresh_token", rest.refreshToken)
// }

func Route(g *echo.Group, h Handler) {
	g.POST("", h.getBook)
}

// type router struct {
// 	group *echo.Group
// }

// func New(group *echo.Group) *router {
// 	s := &router{group: group}
// 	s.setupRoutes()
// 	return s
// }

// func (router *router) setupRoutes() {}
