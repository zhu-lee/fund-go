package resp

import (
	"com.lee/fund/web"
	"fmt"
)

func WritePage(ctx *web.Context, name string, args interface{}) {
	ctx.WriteTemplate("pages/"+name, args)
}

func WriteString(ctx *web.Context, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(ctx.Response, format, args...)
}
