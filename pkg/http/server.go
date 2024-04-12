package http

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListenAndServe() {

	listenAddr := flag.String("listenAddr", ":8080", "The listen addres of the API server")
	flag.Parse()
	r := chi.NewRouter()

	r.Mount("/users", NewUserController().Routes())

	http.ListenAndServe(*listenAddr, r)
}
