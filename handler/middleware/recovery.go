package middleware

import (
	"fmt"
	"net/http"
)

func Recovery(h http.Handler) http.Handler{
	fn:=func(w http.ResponseWriter, r *http.Request){
		//遅れて実行
		defer func(){
			err:=recover()
			if err!=nil{
				fmt.Println("recover",err)
			}
		}()
		//panicが起こる
		h.ServeHTTP(w,r)
	}
	return http.HandlerFunc(fn)
}