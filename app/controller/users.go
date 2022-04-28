package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"myGoSiteProject/app/model"
	"net/http"

)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := model.GetAllUsers()
	if err!=nil{
		http.Error(rw,err.Error(),400)
		return
	}

	err = json.NewEncoder(rw).Encode(users)
	if err!=nil{
		http.Error(rw,err.Error(),400)
		return
	}
}
