package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

// 常规中间件
func Staticmiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "static-middleware")
		logx.Info("static-middleware")
		next(w, r)
	}
}
