package handler

import (
	"net/http"
	"time"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler{
	return &AuthHandler{}
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//ここで2秒の休憩

	time.Sleep(2 *time.Second)
}