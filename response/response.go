package response

import (
	"encoding/json"
	"net/http"
)

type GetResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetDataResponses(w http.ResponseWriter, r *http.Request, data interface{}) {
	response := GetResponse{
		Status:  200,
		Message: "Berhasil",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
