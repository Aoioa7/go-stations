package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var p *model.HealthzResponse
	p = &model.HealthzResponse{}
	p.Message = "OK"

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	err :=json.NewEncoder(w).Encode(p)

	if err != nil {
		log.Println(err)
	}else{
		log.Println("OK")
	}

}
