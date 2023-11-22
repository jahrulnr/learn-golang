package controllers

import (
	"learning/testapp/models"
	"learning/testapp/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "http method not allowed", http.StatusBadRequest)
		return
	}

	user := models.User{}

	user.Name = r.FormValue("name")
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	if user.Name == "" || len(user.Username) < 5 || len(user.Password) < 6 {
		http.Error(w, "Name/Username/Password is invalid", http.StatusBadRequest)
		return
	}

	db := utils.DBConn()
	var userExists bool
	err := db.QueryRow("select IF(COUNT(*),'true','false') from `user` where username=? limit 1", user.Username).Scan(&userExists)

	if err != nil {
		http.Error(w, "Oops! something wrong", http.StatusInternalServerError)
		utils.Log("Database", err.Error())
		db.Close()
		return
	}

	if userExists {
		http.Error(w, "Username already exists", http.StatusNotAcceptable)
		db.Close()
		return
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Oops! something wrong", http.StatusInternalServerError)
		utils.Log("Hash", err.Error())
		db.Close()
		return
	}

	_, errInsert := db.Query("insert into `user` (`name`, `username`, `password`) values (?,?,?)", user.Name, user.Username, user.Password)
	if errInsert != nil {
		http.Error(w, "Oops! something wrong", http.StatusInternalServerError)
		utils.Log("Database", errInsert.Error())
		db.Close()
		return
	}

	user.Password = ""
	response := utils.JsonResponse(true, user)
	db.Close()

	w.Write([]byte(response))
}
