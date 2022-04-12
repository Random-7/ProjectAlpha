package handlers

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	log.Println("handler called")

}
