package main

import (
	"github.com/valyala/fasthttp"
	"fmt"
)

func main() {
	fasthttp.ListenAndServe(":8587", fastHTTPHandler)
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}
