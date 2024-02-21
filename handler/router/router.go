package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	//health
	healthHandler := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", healthHandler.ServeHTTP)
	//todo
	todoService := service.NewTODOService(todoDB)
	todoHandler := handler.NewTODOHandler(todoService)
	mux.HandleFunc("/todos",todoHandler.ServeHTTP)
	//panic
	panicHandler := handler.NewPanicHandler()
	fixedHandler :=middleware.Recovery(panicHandler)
	mux.HandleFunc("/do-panic",fixedHandler.ServeHTTP)

	//os
	osHandler:=middleware.OSName(fixedHandler)
	mux.HandleFunc("/os",osHandler.ServeHTTP)

	//logger
	//middlewareの呼び出しの順番と変数keyのスコープに注意する
	sleepHandler:=handler.NewSleepHandler()
	accesslogHandler:=middleware.AccessLogger(sleepHandler)
	mux.HandleFunc("/log",accesslogHandler.ServeHTTP)
	//Basic認証
	sleepHandler2:=handler.NewSleepHandler()
	basicauthHandler:=middleware.Basicauth(sleepHandler2,"authTestArea")
	mux.HandleFunc("/auth",basicauthHandler.ServeHTTP)

	return mux
}

