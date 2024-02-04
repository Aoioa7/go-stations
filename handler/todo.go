package handler

import (
	"context"
	"encoding/json"
	"log"
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

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	return &model.CreateTODOResponse{TODO: todo}, err
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	todo, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
	return &model.UpdateTODOResponse{TODO: todo}, err
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

//
func(h *TODOHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "POST":
		req:= &model.CreateTODORequest{}
		err1:=json.NewDecoder(r.Body).Decode(&req)
		if err1!=nil{
			log.Println(err1)
			return
		}
		if len(req.Subject)==0{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res,err2:= h.Create(r.Context(), req)
		if err2 !=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		err3:=json.NewEncoder(w).Encode(res)
		if err3!=nil{
			println(err3)
		}

	case "Put":
		req:=&model.UpdateTODORequest{}
		err1:=json.NewDecoder(r.Body).Decode(&req)
		if err1!=nil{
			log.Println(err1)
			return
		}

		if len(req.Subject)==0||req.ID==0{
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res,err2:=h.Update(r.Context(),req)
		if err2!=nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		
		err3:=json.NewEncoder(w).Encode(res)
		if err3!=nil{
			log.Println(err3)
		}

	}
}
