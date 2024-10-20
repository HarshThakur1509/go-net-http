package api

import (
	"log"
	"net/http"

	"github.com/HarshThakur1509/go-net-http/controllers"
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
	router.HandleFunc("POST /idea/", controllers.PostIdea)
	router.HandleFunc("GET /idea/", controllers.GetIdeas)
	router.HandleFunc("GET /idea/{id}", controllers.GetIdeaIndex)
	router.HandleFunc("PUT /idea/{id}", controllers.UpdateIdea)
	router.HandleFunc("DELETE /idea/{id}", controllers.DeleteIdea)

	server := http.Server{
		Addr:    s.addr,
		Handler: cors.Default().Handler(router),
	}
	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
