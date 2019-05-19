package resp

import "com.lee/fund/web"

func WritePage(ctx *web.Context, name string, args interface{}) {
	ctx.WriteTemplate("pages/"+name, args)
}
