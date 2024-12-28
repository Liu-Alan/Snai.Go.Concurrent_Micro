package middleware

import (
	"book/service/user/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

// 调用其它服务的中间件
func AnotherMiddleware(s *svc.AnotherService) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Middleware", s.GetToken())
			logx.Info("another-middleware")
			next(w, r)
		}
	}
}
