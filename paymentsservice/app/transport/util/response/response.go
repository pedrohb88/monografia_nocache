package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Empty(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func Err(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		log.Default().Println(err.Error())
	}
}

func JSON(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(&object)
	if err != nil {
		log.Default().Println(err.Error())
	}
}
