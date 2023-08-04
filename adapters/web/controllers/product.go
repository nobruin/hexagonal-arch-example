package controllers

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nobruin/hexagonal-arch-example/app"
	"net/http"
)

func MakeProductController(router *mux.Router, n *negroni.Negroni, service app.ProductServiceInterface) {
	router.Handle(
		"/products/{id}",
		n.With(negroni.Wrap(getProduct(service))),
	).Methods("GET", "OPTIONS")
}

func getProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
