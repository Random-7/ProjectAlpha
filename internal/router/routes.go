package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func SetRoutes() http.Handler {

	mux := chi.NewRouter()

	//mux.Get("/", handlers.Repo.Home)

	log.Println("Set routes called")

	return mux
}
