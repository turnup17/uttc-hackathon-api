package usecase

import (
	"encoding/json"
	"log"
	"main/dao"
	"main/model"
	"net/http"
)

func Search_usecase(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	var bytes []byte

	rows, err := dao.Db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: dao.Db.Query, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return bytes, err
	}

	// ②-3
	users := make([]model.UserResForHTTPGet, 0)
	for rows.Next() {
		var u model.UserResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
			return bytes, err
		}
		users = append(users, u)
	}

	// ②-4
	bytes, err = json.Marshal(users)
	return bytes, err
}
