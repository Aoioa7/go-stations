package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/TechBowl-japan/go-stations/db"
	"github.com/TechBowl-japan/go-stations/handler/router"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}

}

func realMain() error {
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(todoDB)

	// TODO: サーバーをlistenする
	//http.ListenAndServe(port, mux)

	//ここからgraceful shutdownの実装。まずはサーバーの準備。
	srv:=&http.Server{
		Addr:port,
		Handler:mux,
	}

	//シグナルを通知する準備。
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	//ゴルーチンが終わるまでmain関数が終わらない。
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		//シャットダウンを試みて、エラーハンドリング。
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Print("graceful shutdown失敗", err)
		} else {
			log.Print("graceful shutdown")
		}
	}()
	//サーバーを立てて、エラーハンドリング。
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	wg.Wait()
	return nil
}
