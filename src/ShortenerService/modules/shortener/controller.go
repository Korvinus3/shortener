package shortener

import (
	"html/template"
	"net/http"
	"ShortenerService/modules/templates"
	"ShortenerService/modules/storage"
	"github.com/gorilla/mux"
)

// for tests mocking
var writeData = writeDataToStorage
var getURL = getURLFromStorage

//ProcessLink save link to storage and return its shorten variant
func ProcessLink(w http.ResponseWriter, r *http.Request) {

	var isOk bool

	r.ParseForm()

	url := r.Form["link"]

	if len(url) != 0 && url[0] != "" {

		storageKey, err := writeData(url[0])

		if err == nil && storageKey != "" {

			tpl, _ := template.New("ShortenUrl").Parse(templates.ShortUrlPage)

			tpl.Execute(w, r.Host+"/srt/"+storageKey)

			isOk = true
		}
	}

	if !isOk {

		w.WriteHeader(http.StatusBadRequest)

	}

}

//ProcessRedirect make redirect if link is found, return 404 code otherwise
func ProcessRedirect(w http.ResponseWriter, r *http.Request) {

	var originURL string

	vars := mux.Vars(r)

	if urlKey, ok := vars["key"]; urlKey != "" && ok{

		originURL = getURL(urlKey)

		if originURL == "" {

			w.WriteHeader(http.StatusNotFound)

		} else {

			http.Redirect(w, r, originURL, http.StatusFound)

		}

	} else {

		w.WriteHeader(http.StatusBadRequest)

	}

}

func writeDataToStorage(value string) (string, error) {
	return storage.DBInstance.Set(value)
}

func getURLFromStorage(key string) string {
	return storage.DBInstance.Get(key)
}