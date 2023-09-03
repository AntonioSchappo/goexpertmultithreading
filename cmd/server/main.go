package main

import (
	"net/http"

	handlers "github.com/AntonioSchappo/desafiomultithreading/internal/webserver"
	"github.com/AntonioSchappo/desafiomultithreading/pkg/client"
	"github.com/AntonioSchappo/desafiomultithreading/pkg/requester"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	apiClient := client.NewClient()
	requester := requester.NewRequester(apiClient)
	cepHandler := handlers.NewCepHandler(requester)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/busca", func(r chi.Router) {
		r.Get("/{cep}", cepHandler.GetCep)
	})

	http.ListenAndServe(":8000", r)
}
