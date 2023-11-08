package controllers

import (
	"encoding/json"
	"fmt"
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
			fmt.Printf("err.Error(): %v\n", err.Error())
			return
		}

		timeLog := time.Now().GoString()
		accessToken.Token, err = utils.HashPassword(user.Username + "|" + user.Password + "|" + time.Now().GoString())
		accessToken.CreatedAt = timeLog
		user.Accesstoken = accessToken

		file, _ := json.MarshalIndent(user, "", "")
		response = string(file)
	}
	db.Close()

	w.Write([]byte(response))
}
