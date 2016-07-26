package controllers

import (
	"github.com/Simbory/wemvc"
	"net/http"
)

func init() {
	wemvc.Route("/admin/account/:action", Account{})
	wemvc.Route("/admin/shell/*path", Admin{})
	wemvc.SetFilter("/admin", func(ctx wemvc.Context) {
		if (ctx.Request().URL.Path == "/admin/account/login") {
			return
		}

		loginCookie, err := ctx.Request().Cookie("ADMIN_AUTH")
		if err != nil || len(loginCookie.Value) < 1 {
			http.Redirect(ctx.Response(), ctx.Request(), "/admin/account/login?returnUrl=" + ctx.Request().URL.String(), 302)
			ctx.End()
			return
		}
		ctx.SetItem("name", "Admin")
	})
}