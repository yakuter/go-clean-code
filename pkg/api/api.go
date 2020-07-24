package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	DB *sql.DB
}

func (a *API) GetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get post")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
