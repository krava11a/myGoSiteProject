package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"myGoSiteProject/app/model"
	"net/http"
	"path/filepath"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	main := filepath.Join("public", "html", "usersDynamicPage.html")
	common := filepath.Join("public", "html", "common.html")
	tmpl, err := template.ParseFiles(main, common)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

}

func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	if name == "" || surname == "" {
		http.Error(rw, "Name and surname don't be a pure", 400)
		return
	}
	user := model.User{
		Name:    name,
		Surname: surname,
	}
	err := user.AddUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = json.NewEncoder(rw).Encode("User has been add successfully!!!")

}

func DeleteUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")
	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = user.DeleteUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = json.NewEncoder(rw).Encode("User has been DELETE successfully!!!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userId := p.ByName("userId")
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	user, err := model.GetUserById(userId)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	user.Name = name
	user.Surname = surname
	err = user.UpdateUser()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = json.NewEncoder(rw).Encode("User has been UPDATE successfully!!!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
