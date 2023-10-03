package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carvalhocaio/golang-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
