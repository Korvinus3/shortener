package main

import (
	"ShortenerService/config"
	"net/http"
	"ShortenerService/router"
	"strconv"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/context"
	"ShortenerService/modules/storage"
)

func main() {

	storageErr := storage.InitDB()

	if storageErr != nil {

		panic(storageErr)

	}

	defer storage.CloseDBConn()

	routerHandlers := router.NewRouter()
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(config.Config.HTTPPort), context.ClearHandler(routerHandlers)))

}