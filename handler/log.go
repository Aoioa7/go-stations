package handler

import (
	"net/http"
	"time"
)

type LogHandler struct{}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

func (h *LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//ここで2秒の休憩
	time.Sleep(4 * time.Second)
}
