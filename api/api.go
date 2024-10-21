package api

import (
	"log"
	"net/http"

	"github.com/HarshThakur1509/go-net-http/controllers"
	"github.com/HarshThakur1509/go-net-http/middleware"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /idea", controllers.GetIdeas)
	router.HandleFunc("GET /idea/{id}", controllers.GetIdeaIndex)

	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("POST /idea", controllers.PostIdea)
	adminRouter.HandleFunc("PUT /idea/{id}", controllers.UpdateIdea)
	adminRouter.HandleFunc("DELETE /idea/{id}", controllers.DeleteIdea)

	router.Handle("/", middleware.RequireAuth(adminRouter))

	stack := middleware.MiddlewareChain(middleware.Logger)

	server := http.Server{
		Addr:    s.addr,
		Handler: cors.Default().Handler(stack(router)),
	}
	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
