package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("h", "https://code.hwwt2.com", "host")
	flag.Parse()

	target, err := url.Parse(*host)
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
