package middleware

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Basicauth(h http.Handler,realm string) http.Handler{
	fn:=func(w http.ResponseWriter,r *http.Request){

		err:=godotenv.Load(".env")
		if err!=nil{
			fmt.Println(".env読み込み失敗")
		}
		envID:=os.Getenv("BASIC_AUTH_USER_ID")
		envPASSWORD:=os.Getenv("BASIC_AUTH_USER_PASSWORD")

		id,password,ok:=r.BasicAuth()

		auth1:=true
		if id==""||password==""||!ok{
			auth1=false
		}

		if auth1==false{
			w.Header().Set("WWW-Authenticate",`Basic realm="`+realm+`"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("パスワードは認証されませんでした1"))
			return
		}

		auth2:=true
		if subtle.ConstantTimeCompare([]byte(id),[]byte(envID))!=1||
		subtle.ConstantTimeCompare([]byte(password),[]byte(envPASSWORD))!=1{
			auth2=false
		}

		if auth2==false{
			w.Header().Set("WWW-Authenticate",`Basic realm="`+realm+`"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("パスワードは認証されませんでした2"))
			log.Println(id,envID,password,envPASSWORD)
			return
		}
		
		//ここに辿り着けば認証クリア
		h.ServeHTTP(w,r)
		log.Println("authenticate!!!")
	}
	return http.HandlerFunc(fn)
}
