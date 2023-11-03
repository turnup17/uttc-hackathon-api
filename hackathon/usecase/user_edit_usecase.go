package usecase

import (
	"log"
	"main/dao"
	"main/model"
	"net/http"
	"time"
)

func Edit_usecase(w http.ResponseWriter, r *http.Request, knowledge_info model.KnowledgeReqForHTTPPUT) (model.KnowledgeResForID, error) {
	var response model.KnowledgeResForID

	currentTime := time.Now()
	_, err := dao.Db.Exec("UPDATE knowledge SET name = ?, url = ?, date = ?, category = ?, details = ?, curriculum = ?  WHERE id = ?", knowledge_info.Name, knowledge_info.Url, currentTime, knowledge_info.Category, knowledge_info.Details, knowledge_info.Curriculum, knowledge_info.Id)
	if err != nil {
		log.Printf("fail: dao.Db.Exec, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return response, nil
	}
	response = model.KnowledgeResForID{Id: knowledge_info.Id}
	return response, nil
}
