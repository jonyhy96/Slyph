package main

import (
	_ "Slyph/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("username").(string)
		if !ok && ctx.Request.RequestURI != "/v1/user/login" {
			ctx.Redirect(302, "/v1/user/login")
		} else {
			ctx.Redirect(302, "/v1/index")
		}
	}
	beego.InsertFilter("/", beego.BeforeRouter, FilterUser)
	beego.SetStaticPath("/v1/pod/assets", "static/assets")
	beego.SetStaticPath("/v1/sto/assets", "static/assets")
	beego.SetStaticPath("/v1/st/assets", "static/assets")
	beego.SetStaticPath("/v1/index/assets", "static/assets")
	beego.SetStaticPath("/v1/user/assets", "static/assets")
	beego.SetStaticPath("/v1/auto/assets", "static/assets")
	beego.SetStaticPath("/v1/assets", "static/assets")
	beego.SetStaticPath("/v1/deployment/assets", "static/assets")
	beego.SetStaticPath("/v1/service/assets", "static/assets")
	beego.SetStaticPath("/v1/index/static/docker.png", "static/img/docker.png")
	beego.SetStaticPath("/v1/deployment/static/deployment.jpg", "static/img/deployment.jpg")
	beego.SetStaticPath("/v1/service/static/service.png", "static/img/service.png")
	beego.SetStaticPath("/v1/user/static/assets/i/favicon.png", "static/assets/i/favicon.png")
	beego.SetStaticPath("/v1/pod/static/docker.png", "static/img/docker.png")
	beego.SetStaticPath("/v1/sto/static/sto.png", "static/img/sto.png")
	beego.SetStaticPath("/v1/st/static/st.png", "static/img/st.png")
	beego.SetStaticPath("/v1/auto/static/docker.png", "static/img/docker.png")
	beego.SetStaticPath("/v1/static/docker.png", "static/img/docker.png")
	beego.SetStaticPath("/v1/pod/static", "static")
	beego.SetStaticPath("/v1/sto/static", "static")
	beego.SetStaticPath("/v1/st/static", "static")
	beego.SetStaticPath("/v1/service/static", "static")
	beego.SetStaticPath("/v1/deployment/static", "static")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
