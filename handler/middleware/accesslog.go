package middleware

import (
	"context"
	"encoding/json"
	_ "fmt"
	"net/http"
	"time"

	"github.com/mileusna/useragent"
)

type AccessLog struct{
	Timestamp time.Time `json:"timestamp"`
	Latency int64 `json:"latency"`
	Path string `json:"path"`
	OS string `json:"os"`
}

func AccessLogger(h http.Handler) http.Handler{

	fn:=func(w http.ResponseWriter,r *http.Request){

		accessTime:=time.Now()
		thisPath:=r.URL.Path
		thisOS:="none"
		osKey:="os"

		//OSのkeyとvalueの準備
		ua:=r.UserAgent()
		os:=useragent.Parse(ua).OS
		ctx:=context.WithValue(r.Context(),osKey,os)
		
		//ここのハンドラーでsleepする、でないと時間差が丸められて0になる
		h.ServeHTTP(w,r)

		defer func(){
			//os
			if osValue,ok :=r.WithContext(ctx).Context().Value(osKey).(string);ok{
				thisOS=osValue
			}
			
			//latency
			thisLatency:=time.Since(accessTime).Milliseconds()
			accessLog:=AccessLog{
				Timestamp:accessTime,
				Latency:thisLatency,
				Path:thisPath,
				OS:thisOS,
			}
			//あとはここをjsonで出力するだけ
			//fmt.Printf("%+v\n",accessLog)
			json.NewEncoder(w).Encode(accessLog)

		}()
	}
	return http.HandlerFunc(fn)
}
