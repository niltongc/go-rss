package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 400, "Something went wrong")
}

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, struct{}{})
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}
