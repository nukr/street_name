package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nukr/street_name/pkg/types"
)

// NewRouter provide http.ListenAndServe handler
func NewRouter(addr *types.Address) http.Handler {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(corsPreflight())
	r.HandleFunc("/", healthCheck())
	r.HandleFunc("/list", getCountry(addr))
	r.HandleFunc("/list/{country}", getCounty(addr))
	r.HandleFunc("/list/{country}/{county}", getDistrict(addr))
	r.HandleFunc("/list/{country}/{county}/{district}", getStreet(addr))
	r.HandleFunc("/healthz", healthCheck())
	return r
}

func corsPreflight() http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		h := w.Header()
		h.Add("Access-Control-Allow-Origin", "*")
		h.Add("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
		h.Add("Access-Control-Allow-Headers", "Accept-Language, Content-Type")
		h.Add("Access-Control-Max-Age", "1728000")
		w.WriteHeader(200)
	}
}

func healthCheck() http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		req *http.Request,
	) {
		h := w.Header()
		h.Add("Cache-Control", "no-cache, no-store, must-revalidate")
		h.Add("Pragma", "no-cache")
		h.Add("Expires", "0")
		fmt.Fprintf(w, "OK")
	}
}

func getCountry(addr *types.Address) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		req *http.Request,
	) {
		h := w.Header()
		h.Add("Content-Type", "application/json; charset=utf-8")
		h.Add("Access-Control-Allow-Origin", "*")
		acceptLang := strings.Split(strings.ToLower(req.Header.Get("Accept-Language")), ",")[0]
		if acceptLang == "" {
			acceptLang = "zh-tw"
		}
		fmt.Fprintf(w, addr.Country[acceptLang].ToJSON())
	}
}

func getCounty(addr *types.Address) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		req *http.Request,
	) {
		h := w.Header()
		h.Add("Content-Type", "application/json; charset=utf-8")
		h.Add("Access-Control-Allow-Origin", "*")
		country := mux.Vars(req)["country"]
		if v, ok := addr.County[country]; ok {
			fmt.Fprintf(w, v.ToJSON())
		} else {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
	}
}

func getDistrict(addr *types.Address) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		req *http.Request,
	) {

		country := mux.Vars(req)["country"]
		county := mux.Vars(req)["county"]

		h := w.Header()
		h.Add("Content-Type", "application/json; charset=utf-8")
		h.Add("Access-Control-Allow-Origin", "*")
		if v, ok := addr.District[country+county]; ok {
			fmt.Fprintf(w, v.ToJSON())
		} else {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
	}
}

func getStreet(addr *types.Address) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		req *http.Request,
	) {
		h := w.Header()
		h.Add("Content-Type", "application/json; charset=utf-8")
		h.Add("Access-Control-Allow-Origin", "*")

		country := mux.Vars(req)["country"]
		county := mux.Vars(req)["county"]
		district := mux.Vars(req)["district"]
		if street, ok := addr.Street[country+county+district]; ok {
			fmt.Fprintf(w, street.ToJSON())
		} else {
			http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		}
	}
}
