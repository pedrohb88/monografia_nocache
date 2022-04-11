package util

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ParamAsInt(r *http.Request, key string) int {
	param := chi.URLParam(r, key)
	intParam, _ := strconv.Atoi(param)
	return intParam
}

func DecodeJSON(r *http.Request, object interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(object)
}
