package router

import (
	"github.com/gorilla/mux"
	"ShortenerService/modules/mainpage"
	"ShortenerService/modules/shortener"
)

//NewRouter create router
func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", mainpage.GetForm).Name("GetMainPageForm").Methods("GET")
	router.HandleFunc("/short-link/", shortener.ProcessLink).Name("ProcessLink").Methods("POST")
	router.HandleFunc("/srt/{key}", shortener.ProcessRedirect).Name("ProcessRedirect").Methods("GET")

	return router

}
