package usecase

import (
	"encoding/json"
	"log"
	"main/dao"
	"main/model"
	"net/http"
	"time"
)

func Search_usecase(w http.ResponseWriter, r *http.Request) ([]byte, error) {
    var bytes []byte

    rows, err := dao.Db.Query("SELECT id, name, url, date, category, details, curriculum FROM knowledge")
    if err != nil {
        log.Printf("fail: dao.Db.Query, %v\n", err)
        w.WriteHeader(http.StatusInternalServerError)
        return bytes, err
    }

    var knowledges []model.KnowledgeResForHTTPGET

    for rows.Next() {
        var k model.KnowledgeResForHTTPGET
        var dateValue string // Assuming the date value is a string in the database

        if err := rows.Scan(&k.Id, &k.Name, &k.Url, &dateValue, &k.Category, &k.Details, &k.Curriculum); err != nil {
            log.Printf("fail: rows.Scan, %v\n", err)

            if err := rows.Close(); err != nil {
                log.Printf("fail: rows.Close(), %v\n", err)
            }
            w.WriteHeader(http.StatusInternalServerError)
            return bytes, err
        }

        // Convert the string dateValue into a time.Time
        k.Date, err = time.Parse("2006-01-02", dateValue)
		if err != nil {
			log.Printf("fail: time.Parse, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)  // Corrected 'httpInternalServerError' to 'http.StatusInternalServerError'
			return bytes, err
		}

        knowledges = append(knowledges, k)
    }

    bytes, err = json.Marshal(knowledges)
    if err != nil {
        log.Printf("fail: json.Marshal, %v\n", err)
        w.WriteHeader(http.StatusInternalServerError)
        return bytes, err
    }

    w.Header().Set("Content-Type", "application/json")

    return bytes, nil
}
