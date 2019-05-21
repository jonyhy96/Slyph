package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"Slyph/models"
)

// AddStorage struct
type AddStorage struct {
	baseController
}

// GetStorage struct
type GetStorage struct {
	baseController
}

// DeleteStorage struct
type DeleteStorage struct {
	baseController
}

// Get to get PV
// @Title Get PV
// @Description Get PV
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
func (d *DeleteStorage) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "deletePV.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post to delete a specify PV.
// /v1/sto/delete
// @Title Delete a PV
// @Description Delete a PV
// @Param PV body models.Delete true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /delete [post]
func (d *DeleteStorage) Post() {
	defer d.catchException()
	glog.V(2).Infof("delete a PV URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Delete{}
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &req); err != nil {
		d.badRequest(models.Error{Error: err.Error()})
		return
	}
	if len(req.Name) == 0 {
		d.badRequest(models.Error{Error: "名称为空"})
		return
	}
	if code, err := models.DeletePV(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}
	d.success(models.Response{Action: "delete PV successs", Status: "success"})
	return
}

// Get to create a new PV.
// /v1/sto/create
// @router /create [get]
func (d *AddStorage) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "createsto.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post create a new PV.
// /v1/sto/create
// @Title Create a PV
// @Description Get 1 job's detail info
// @Param PV body models.PV true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /create [post]
func (d *AddStorage) Post() {
	defer d.catchException()
	glog.V(2).Infof("add a PV URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.PV{}
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &req); err != nil {
		d.badRequest(models.Error{Error: err.Error()})
		return
	}
	if len(req.Name) == 0 {
		d.badRequest(models.Error{Error: "名称为空"})
		return
	}
	if code, err := models.CreatePV(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}

	d.success(models.Response{Action: "add PV successs", Status: "success"})
	return
}

// Getjson to get all PV's information.
// /v1/sto/getinfo
// @Title Get a PVinfo
// @Description Get a PV info
// @Success 200 {object} models.PV{}
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /getjson [get]
func (c *GetStorage) Getjson() {
	PVinfo := models.GetPV()
	c.Data["json"] = PVinfo
	c.ServeJSON()
}

// Get to get all PV's information.
// /v1/sto/getinfo
// @router /getinfo [get]
func (c *GetStorage) Get() {
	if c.GetSession("username") == "root" {
		c.TplName = "getsto.html"
	} else {
		c.TplName = "404.html"
	}
}
