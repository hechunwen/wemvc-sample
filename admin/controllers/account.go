package controllers

import (
	"github.com/Simbory/wemvc"
	"net/http"
)

type Account struct {
	wemvc.Controller
}

func (acc Account) GetLogin() wemvc.ActionResult {
	return acc.View()
}

func (acc Account) PostLogin() wemvc.ActionResult {
	var email = acc.Request.Form.Get("email")
	var pwd = acc.Request.Form.Get("password")
	if email == "simbory@sina.cn" && pwd == "123456" {
		var returnUrl = acc.Request.URL.Query().Get("returnUrl")
		if len(returnUrl) < 1 {
			returnUrl = acc.Namespace().GetName() + "/shell/index"
		}
		var cookie = &http.Cookie{
			Name:     "ADMIN_AUTH",
			Value:    email,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			Domain:   acc.Request.URL.Host,
		}
		http.SetCookie(acc.Response, cookie)
		return acc.Redirect(returnUrl)
	}
	acc.ViewData["email"] = email
	acc.ViewData["error"] = "invalid email or password"
	return acc.View()
}
