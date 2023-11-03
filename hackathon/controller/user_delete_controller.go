package controller

import (
	"encoding/json"
	"log"
	"main/usecase"
	"net/http"
	"strings"
)

func Delete_controller(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract the knowledge ID from the request URL
	// e.g., /knowledge/01HE79R4VV0000000000000000
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	knowledgeID := parts[2]

	// Call the usecase function to delete knowledge
	response, err := usecase.Delete_usecase(w, r, knowledgeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("fail: json.Encode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
