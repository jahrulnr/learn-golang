package router

import (
	"learning/testapp/controllers"
	"net/http"
)

func SetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/register", controllers.Register)

	return mux
}
