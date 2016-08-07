package main

import (
	_ "github.com/Simbory/wemvc-sample/controllers"
	_ "github.com/Simbory/wemvc-sample/admin"
	"github.com/Simbory/wemvc"
)

func sayHello(name interface{}) string {
	n,ok := name.(string)
	if ok {
		return "Hello, " + n
	}
	return ""
}

func main() {
	wemvc.
		AddViewFunc("sayHello", sayHello).
		StaticDir("/css/").
		StaticDir("/js/").
		StaticFile("/favicon.ico").
		Run(8080)
}