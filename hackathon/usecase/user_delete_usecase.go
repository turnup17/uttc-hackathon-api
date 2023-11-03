package usecase

import (
	"log"
	"main/dao"
	"main/model"
	"net/http"
)

func Delete_usecase(w http.ResponseWriter, r *http.Request, knowledgeID string) (model.KnowledgeResForID, error) {
	var response model.KnowledgeResForID

	_, err := dao.Db.Exec("DELETE FROM knowledge WHERE id = ?", knowledgeID)
	if err != nil {
		log.Printf("fail: dao.Db.Exec, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return response, err
	}
	response = model.KnowledgeResForID{Id: knowledgeID}
	log.Printf("success")
	return response, nil
}
