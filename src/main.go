package main

import (
	"filebeat-elk/router"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&ecslogrus.Formatter{})

	logFilePath := "/app/logs/out.log"
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	defer file.Close()
	fmt.Println("Start service")
	log.Info("Service started!")

	routes := router.NewRouter()

	//_, err = elasticsearch.NewDefaultClient()
	//if err != nil {
	//	log.Error(err.Error())
	//	os.Exit(1)
	//}

	server := &http.Server{
		Addr:           ":8888",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        routes,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
