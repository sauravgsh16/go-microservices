package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sauravgsh16/microservices/mvc/services"
	"github.com/sauravgsh16/microservices/mvc/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userParams := r.URL.Query().Get("user_id")
	id, err := strconv.ParseInt(userParams, 10, 64)
	if err != nil {

		appErr := &utils.AppError{
			ID:      http.StatusBadRequest,
			Message: "user_id must be a number",
		}
		jsonErr, _ := json.Marshal(appErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr)
		return
	}

	user, appErr := services.GetUser(id)
	if appErr != nil {

		jsonErr, _ := json.Marshal(appErr)
		w.WriteHeader(appErr.ID)
		w.Write(jsonErr)
		return
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
}
