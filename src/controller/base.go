package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Base(w http.ResponseWriter, r *http.Request) {
	log.Info("This Log base")
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode("Welcome base API"); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
