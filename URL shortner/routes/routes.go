package routes

import (
	"github.com/gorilla/mux"
	"github.com/igauravrana/URL-Shortner/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/shorten", controllers.ShortenURLHandler).Methods("POST")
	r.HandleFunc("/redirect/{shortUrl}", controllers.RedirectToOriginalUrl).Methods("GET")
	r.HandleFunc("/delete/{shortUrl}", controllers.DeleteURLHandler).Methods("DELETE")
	r.HandleFunc("/url/{shortUrl}", controllers.GetURLHandler).Methods("GET")

	return r
}
