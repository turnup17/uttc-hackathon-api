package controller

import (
	"encoding/json"
	"log"
	"main/model"
	"main/usecase"
	"net/http"
)

func Register_controller(w http.ResponseWriter, r *http.Request) {
	var user_info model.UserResForHTTPPost
	decoded := json.NewDecoder(r.Body)
	if err := decoded.Decode(&user_info); err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := usecase.Register_usecase(w, r, user_info)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("fail: json.Encode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
