package admin

import (
	adminCtrl "github.com/Simbory/wemvc-sample/admin/controllers"
	"github.com/Simbory/wemvc"
	"net/http"
)

func adminLoginFilter(ctx wemvc.Context) {
	if ctx.Request().URL.Path == "/admin/account/login" {
		return
	}

	loginCookie, err := ctx.Request().Cookie("ADMIN_AUTH")
	if err != nil || len(loginCookie.Value) < 1 {
		http.Redirect(ctx.Response(), ctx.Request(), "/admin/account/login?returnUrl="+ctx.Request().URL.String(), 302)
		ctx.End()
		return
	}
	ctx.SetItem("name", "Admin")
}

func init(){
	wemvc.Namespace("admin").
		Route("/shell/*pathInfo", adminCtrl.Admin{}).
		Route("/account/:action", adminCtrl.Account{})
	wemvc.SetFilter("/admin/shell/", adminLoginFilter)
}