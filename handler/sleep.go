package handler

import (
	"net/http"
	"time"
)

type SleepHandler struct{}

func NewSleepHandler() *SleepHandler{
	return &SleepHandler{}
}

func (h *SleepHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//ここで2秒の休憩
	time.Sleep(2 *time.Second)
}