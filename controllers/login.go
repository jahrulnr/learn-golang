package controllers

import (
	"learning/testapp/models"
	"learning/testapp/utils"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "http method not allowed", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || len(password) < 6 {
		http.Error(w, "Username/Password is invalid", http.StatusBadRequest)
		return
	}

	db := utils.DBConn()
	result, err := db.Query("select * from user where username=? and password=? limit 1", username, password)

	if err != nil {
		http.Error(w, "Username/Password is invalid", http.StatusNotAcceptable)
		db.Close()
		return
	}

	var response = ""
	if !result.Next() {
		http.Error(w, "Username/Password is invalid", http.StatusNotAcceptable)
		return
	} else {
		user := models.User{}
		accessToken := models.AccessToken{}
		err = result.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			http.Error(w, "Oops! something wrong", http.StatusInternalServerError)
			utils.Log("Database", err.Error())
			db.Close()
			return
		}

		timeLog := time.Now().Format("2006-01-02 15:04:05")
		accessToken.Token, err = utils.HashPassword(user.Username + "|" + user.Password + "|" + timeLog)
		accessToken.CreatedAt = timeLog
		user.Accesstoken = accessToken

		result, err = db.Query("insert into accessToken (user_id, accessToken) values (?, ?)", user.Id, accessToken.Token)
		if err != nil {
			http.Error(w, "Oops! something wrong", http.StatusInternalServerError)
			utils.Log("Database", err.Error())
			db.Close()
			return
		}

		response = utils.JsonResponse(true, user)
	}
	db.Close()

	w.Write([]byte(response))
}
