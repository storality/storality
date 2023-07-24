package middleware

import (
	"fmt"
	"net/http"

	"storality.com/storality/internal/helpers/exceptions"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func(){
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				exceptions.ServerError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}