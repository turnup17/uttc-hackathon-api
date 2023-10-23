package usecase

import (
	"errors"
	"github.com/oklog/ulid/v2"
	"log"
	"main/dao"
	"main/model"
	"net/http"
)

func Register_usecase(w http.ResponseWriter, r *http.Request, user_info model.UserResForHTTPPost) (model.UserResForID, error) {
	var response model.UserResForID
	if Check_input(user_info) {
		log.Println("fail: Invalid request data")
		w.WriteHeader(http.StatusBadRequest)
		return response, errors.New("An error ovvured")
	}

	ulid := ulid.MustNew(ulid.Now(), nil)

	_, err := dao.Db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", ulid.String(), user_info.Name, user_info.Age)
	if err != nil {
		log.Printf("fail: dao.Db.Exec, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return response, nil
	}
	response = model.UserResForID{Id: ulid.String()}
	return response, nil
}
func Check_input(user_info model.UserResForHTTPPost) bool {
	return user_info.Name == "" || len(user_info.Name) > 50 || user_info.Age < 20 || user_info.Age > 80
}
