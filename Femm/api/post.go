package api

import (
	"encoding/json"
	"fem/go/museum/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exhibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exhibition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exhibition)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
	// w.Header().Set("Content-Type", "application/json")
}