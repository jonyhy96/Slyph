package controllers

import (
	"Slyph/models"
)

// IndexController struct
type IndexController struct {
	baseController
}

// Get to enter index
// @router / [get]
func (m *IndexController) Get() {
	var name = m.GetSession("username")
	if name == "root" {
		models.Truncate() // This will clean up the table you have to ensure no deployment were running before the slyph start
		//TODO redesign sql formate or code to surpport rollback
		m.TplName = "admin-index.html"
	} else {
		m.TplName = "404.html"
	}
}
