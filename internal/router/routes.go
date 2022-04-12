package router

import (
	"log"
	"net/http"

	"github.com/Random-7/ProjectAlpha/internal/handlers"
	"github.com/go-chi/chi"
)

func SetRoutes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)

	log.Println("Set routes called")

	return mux
}
