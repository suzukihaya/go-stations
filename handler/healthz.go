package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

type Encoder struct {
	// contains filtered or unexported fields
}

// A HealthzHandler implements he	alth check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respons := &model.HealthzResponse{
		Message: "OK",
	}
	err := json.NewEncoder(w).Encode(respons)
	if err != nil {
		log.Println(err)
	}
}
