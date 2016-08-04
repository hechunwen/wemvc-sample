package admin

import (
	adminCtrl "github.com/Simbory/wemvc-sample/admin/controllers"
	"github.com/Simbory/wemvc"
	"net/http"
)

func adminLoginFilter(ctx wemvc.Context) {
	ns := ctx.Namespace()
	var nsName string
	if ns != nil {
		nsName = ns.GetName()
	}
	loginCookie, err := ctx.Request().Cookie("ADMIN_AUTH")
	if err != nil || len(loginCookie.Value) < 1 {
		http.Redirect(ctx.Response(), ctx.Request(), nsName + "/account/login?returnUrl="+ctx.Request().URL.String(), 302)
		ctx.EndContext()
		return
	}
	ctx.SetItem("name", "Admin")
}

func init(){
	wemvc.Namespace("admin").
		Route("/shell/*pathInfo", adminCtrl.Admin{}).
		Route("/account/:action", adminCtrl.Account{}).
		SetFilter("/shell/", adminLoginFilter)
}