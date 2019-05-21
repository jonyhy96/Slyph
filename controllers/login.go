package controllers

import (
	"github.com/astaxie/beego"
	"Slyph/models"
	"net"
)

// LoginController struct
type LoginController struct {
	beego.Controller
}

// SignupController struct
type SignupController struct {
	beego.Controller
}

// Post to handle login information
// @Title Post to login system
// @Description Post to login system
// @Param username query string true
// @Param password query string true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /login [post]
func (main *LoginController) Post() {
	log := models.Log{}
	name := main.GetString("username")
	password := main.GetString("password")
	var pass string
	db, err := models.LinkTomysql()
	defer db.Close()
	if err != nil {
		log.Time = models.Gettime()
		log.Etype = "error"
		log.Event = "Failed to link mysql:" + err.Error()
		go models.Insertlog(&log)
	} else {
		rows, err := db.Query("select password from user where name=? ", name)
		for rows.Next() {
			rows.Scan(&pass)
		}
		defer rows.Close()
		if err != nil {
			log.Time = models.Gettime()
			log.Etype = "error"
			log.Event = name + " failed to get password ip:" + main.Ctx.Request.RemoteAddr
			go models.Insertlog(&log)
			main.Redirect("/v1/user/login", 302)
		} else if pass == password {
			log.Time = models.Gettime()
			log.Etype = "info"
			log.Event = name + " Login from  " + main.Ctx.Request.RemoteAddr
			go models.Insertlog(&log)
			main.SetSession("username", name)
			main.Redirect("/v1/index/", 302)
		} else {
			log.Time = models.Gettime()
			log.Etype = "error"
			log.Event = name + " login with wrong password ip:" + main.Ctx.Request.RemoteAddr
			go models.Insertlog(&log)
			main.Redirect("/v1/user/login", 302)
		}
	}
}

// Get to enter login
// @router /login [get]
func (main *LoginController) Get() {
	main.TplName = "login.html"
}

func getLocalIP() string {
	addrSlice, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}
	for _, addr := range addrSlice {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}
