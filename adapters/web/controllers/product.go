package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nobruin/hexagonal-arch-example/adapters/dto"
	"github.com/nobruin/hexagonal-arch-example/app"
)

func MakeProductController(router *mux.Router, n *negroni.Negroni, service app.ProductServiceInterface) {
	router.Handle(
		"/products/{id}",
		n.With(negroni.Wrap(getProduct(service))),
	).Methods("GET", "OPTIONS")

	router.Handle(
		"/products",
		n.With(negroni.Wrap(createProduct(service))),
	).Methods("POST")

	router.Handle(
		"/products{id}/enable",
		n.With(negroni.Wrap(enableProduct(service))),
	).Methods("GET")

	router.Handle(
		"/products/{id}/disable",
		n.With(negroni.Wrap(disableProduct(service))),
	).Methods("GET")
}

func getProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(messageToJson(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(messageToJson(err.Error()))
			return
		}
	})
}

func createProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productRequest = dto.NewProductRequest()
		err := json.NewDecoder(r.Body).Decode(productRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(messageToJson(err.Error()))
			return
		}
		product, err := service.Create(productRequest.Name, productRequest.Price)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(messageToJson(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(messageToJson(err.Error()))
			return
		}
	})
}

func enableProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(messageToJson(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(messageToJson(err.Error()))
			return
		}

		_, err = service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(messageToJson(err.Error()))
			return
		}
	})
}

func disableProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(messageToJson(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(messageToJson(err.Error()))
			return
		}

		_, err = service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(messageToJson(err.Error()))
			return
		}
	})
}
