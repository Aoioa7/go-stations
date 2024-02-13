package middleware

import (
	"fmt"
	"net/http"
	"time"
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
		thisPath:= r.URL.Path
		osKey:="os"
		thisOS:="none"

		h.ServeHTTP(w,r)

		defer func(){
			//os
			if osValue,ok :=r.Context().Value(osKey).(string);ok{
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
			fmt.Printf("%+v\n",accessLog)
		}()
	}
	return http.HandlerFunc(fn)
}
