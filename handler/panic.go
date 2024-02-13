package handler

import (
	"net/http"
)

//構造体
type PanicHandler struct{}

func NewPanicHandler() *PanicHandler{
	return &PanicHandler{}
}

func (h *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	panic("Panic!!!!!")
}