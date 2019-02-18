package mainpage

import (
	"html/template"
	"net/http"
	"ShortenerService/modules/templates"
)

//GetForm return the form for main page
func GetForm (w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		tpl, _ := template.New("MainForm").Parse(templates.MainForm)
		tpl.Execute(w, "")
	}

}