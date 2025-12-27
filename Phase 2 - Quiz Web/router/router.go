package router

import (
	"net/http"
	"quiz/controller"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Home)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
