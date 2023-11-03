package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"main/controller"
	"main/dao"
	"main/usecase"
	"net/http"
	"strings"
)

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "http://localhost:3000" || origin == "localhost:3000" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "https://uttc-hackathon-web.vercel.app")
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST") // Specify the allowed methods

	switch r.Method {
	case http.MethodOptions:
		// Handle the preflight request by responding with the allowed headers and methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		controller.Register_controller(w, r)
	case http.MethodGet:
		// ②-1
		controller.Search_controller(w, r)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func knowledgeDeleteHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {

	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/user", handler)
	http.HandleFunc("/knowledge/", knowledgeDeleteHandler)
	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.CloseDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
