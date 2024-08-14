package router

import (
	"filebeat-elk/controller"
	"fmt"
	//"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//func NewRouter() *httprouter.Router {
//	router := httprouter.New()
//	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//		log.Info("Welcome home log")
//		fmt.Fprint(w, "Welcome home")
//		router.GET("/api/base", controller.Base)
//	})
//	return router
//}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Welcome home log")
		fmt.Fprint(w, "Welcome home")
	})
	mux.HandleFunc("/api/base", controller.Base)
	return mux
}
