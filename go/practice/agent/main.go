package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	target, err := url.Parse("http://share-arch-cicd-proxy.sit.yumc.local")
	// target, err := url.Parse("http://10.1.2.202:9300")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	director := proxy.Director
	proxy.Director = func(r *http.Request) {
		director(r)
		r.Host = target.Host
	}
	proxy.Transport = &http.Transport{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	engine := gin.New()
	engine.Any("/*any", func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	agent := http.Server{
		Addr:    ":5000",
		Handler: engine,
	}

	agent.ListenAndServe()
}
