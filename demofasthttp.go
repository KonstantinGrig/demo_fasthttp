package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8587"
	}
	fasthttp.ListenAndServe(":"+port, fastHTTPHandler)
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}
