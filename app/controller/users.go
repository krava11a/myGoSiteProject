package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"myGoSiteProject/app/model"
	"net/http"
	"path/filepath"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := model.GetAllUsers()
	if err!=nil{
		http.Error(rw,err.Error(),400)
		return
	}
	main := filepath.Join("public","html","usersDynamicPage.html")
	common :=filepath.Join("public","html","common.html")
	tmpl, err := template.ParseFiles(main,common)
	if err!=nil{
		http.Error(rw,err.Error(),400)
		return
	}
	err = tmpl.ExecuteTemplate(rw,"users",users)
	if err!=nil{
		http.Error(rw,err.Error(),400)
		return
	}

}
