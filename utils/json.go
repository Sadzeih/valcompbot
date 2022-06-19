package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data any) {
	j, _ := json.Marshal(data)
	w.Write(j)
}
