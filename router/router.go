package router

import (
	"learning/testapp/controllers"
	"net/http"
)

func SetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", controllers.Login)

	return mux
}
