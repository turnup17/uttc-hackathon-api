package controller

import (
	"log"
	"main/usecase"
	"net/http"
)

func Search_controller(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // To be filled
	if name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// ②-2
	bytes, err := usecase.Search_usecase(w, r, name)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
