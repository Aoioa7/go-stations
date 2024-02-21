package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mileusna/useragent"
)


type key string
var osKey key

func OSName(h http.Handler) http.Handler{
	osKey="os"
	fn:=func(w http.ResponseWriter,r *http.Request){
		ua:=r.UserAgent()
		os:=useragent.Parse(ua).OS
		ctx:=context.WithValue(r.Context(),osKey,os)
		//ターミナルからcurlコマンドでAPIを叩いた時と、リンクを開いてwabブラウザからAPIを叩くときではHTTPリクエストヘッダに含まれるUserAgentが異なる
		fmt.Println(ua,os)
		
		h.ServeHTTP(w,r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}