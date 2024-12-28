package main

import (
	"flag"
	"fmt"
	"net/http"

	"book/service/user/api/internal/config"
	"book/service/user/api/internal/handler"
	"book/service/user/api/internal/middleware"
	"book/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("global middleware")
			next(w, r)
		}
	})
	// 常规中间件
	server.Use(middleware.Staticmiddleware)
	// 调用其它服务的中间件
	aos := svc.NewAnotherService()
	server.Use(middleware.AnotherMiddleware(aos))

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
