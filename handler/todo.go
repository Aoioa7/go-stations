package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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
	todos,err := h.svc.ReadTODO(ctx, req.PrevID, req.Size)
	return &model.ReadTODOResponse{TODOs:todos}, err
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	todo, err := h.svc.UpdateTODO(ctx, req.ID, req.Subject, req.Description)
	return &model.UpdateTODOResponse{TODO: todo}, err
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	err := h.svc.DeleteTODO(ctx,req.IDs)
	return &model.DeleteTODOResponse{}, err
}

//
func(h *TODOHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "POST":
		req:= &model.CreateTODORequest{}
		if err:=json.NewDecoder(r.Body).Decode(req);
		err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("NG1",err)
			return
		}
		if len(req.Subject)==0{
			w.WriteHeader(http.StatusBadRequest)
			log.Println("NG2")
			return
		}
		res,err:= h.Create(r.Context(), req)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("NG3",err)
			return
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		log.Println("OK")
		if err:=json.NewEncoder(w).Encode(res);
		err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("NG4",err)
			return
		}

	case "PUT":
		req:=&model.UpdateTODORequest{}
		err:=json.NewDecoder(r.Body).Decode(req)
		if err!=nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(req.Subject)==0||req.ID==0{
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res,err:=h.Update(r.Context(),req)
		if err!=nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type","application/json")
		if err:=json.NewEncoder(w).Encode(res);
		err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "GET":
		req:=&model.ReadTODORequest{PrevID:0,Size:5}
		prevID:=r.URL.Query().Get("prev_id")
		size:=r.URL.Query().Get("size")
		var err error
		if prevID!=""{
			req.PrevID,err=strconv.ParseInt(prevID,10,64)
		}
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if size!=""{
			req.Size,err=strconv.ParseInt(size,10,64)
		}
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res,err:=h.Read(r.Context(),req)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type","application/json")
		if err:=json.NewEncoder(w).Encode(res);err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "DELETE":
		req:=&model.DeleteTODORequest{}	
		err:=json.NewDecoder(r.Body).Decode(req)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(req.IDs)==0{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res,err:=h.Delete(r.Context(),req)

		var errnotfound *model.ErrNotFound
		if errors.As(err,&errnotfound){
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type","application/json")
		if err:=json.NewEncoder(w).Encode(res);err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
