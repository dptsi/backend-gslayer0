package usercontroller

import (
	"encoding/json"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(map[string]string{"message": "ok"})
	w.Write(response)
}
