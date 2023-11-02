package usecase

import (
	"errors"
	"github.com/oklog/ulid/v2"
	"log"
	"main/dao"
	"main/model"
	"net/http"
	"time"
)

func Check_input(knowledge_info model.KnowledgeResForHTTPPost) bool {
	return knowledge_info.Name == "" || len(knowledge_info.Name) > 50
}
func Register_usecase(w http.ResponseWriter, r *http.Request, knowledge_info model.KnowledgeResForHTTPPost) (model.KnowledgeResForID, error) {
	var response model.KnowledgeResForID
	if Check_input(knowledge_info) {
		log.Println("fail: Invalid request data")
		w.WriteHeader(http.StatusBadRequest)
		return response, errors.New("An error ovvured")
	}

	ulid := ulid.MustNew(ulid.Now(), nil)
	currentTime := time.Now()
	log.Println(ulid)
	_, err := dao.Db.Exec("INSERT INTO knowledge (id, name, url, date, category, details, curriculum) VALUES (?, ?, ?, ?, ? ,?, ?)", ulid.String(), knowledge_info.Name, knowledge_info.Url, currentTime, knowledge_info.Category, knowledge_info.Details, knowledge_info.Curriculum)
	if err != nil {
		log.Printf("fail: dao.Db.Exec, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return response, nil
	}
	response = model.KnowledgeResForID{Id: ulid.String()}
	return response, nil
}
