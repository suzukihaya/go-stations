package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "メソットが許可されていません", http.StatusMethodNotAllowed)
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTODORequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "400 BadRequest", http.StatusBadRequest)
		return
	}
	if req.Subject == "" {
		http.Error(w, "400 BadRequest", http.StatusBadRequest)
		return
	}
	todo, err := h.svc.CreateTODO(r.Context(), req.Subject, req.Description)
	if err != nil {
		http.Error(w, "failed to TODO", http.StatusInternalServerError)
		return
	}

	res := model.CreateTODOResponse{TODO: *todo}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "failed to Encode response", http.StatusInternalServerError)
		return
	}
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
