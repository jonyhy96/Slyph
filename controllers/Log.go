package controllers

import (
	"k8s/models"
)

type LogController struct {
	baseController
}

// Get to enter index
// @router / [get]
func (m *LogController) Get() {
	var name = m.GetSession("username")
	if name == "root" {
		m.TplName = "admin-log.html"
	} else {
		m.TplName = "404.html"
	}
}

// Getjson to get all pods's information.
// /v1/pod/getinfo
// @Title Get a podinfo
// @Description Get a pods info
// @Success 200 {object} models.Log
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /log [get]
func (m *LogController) Getlog() {
	var logs []models.Log
	logs = models.Readlog()
	m.Data["json"] = logs
	m.ServeJSON()
	m.Finish()
}
