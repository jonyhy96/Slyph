package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"Slyph/models"
	"Slyph/models/service"
)

// CreateService struct
type CreateService struct {
	baseController
}

// GetService struct
type GetService struct {
	baseController
}

// DeleteService struct
type DeleteService struct {
	baseController
}

// Post to delete the specify service.
// /v1/service/delete
// @Title Delete a Service
// @Description Delete a Service
// @Param service body models.Delete true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /delete [post]
func (d *DeleteService) Post() {
	defer d.catchException()
	glog.V(2).Infof("delete a Service URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Delete{}
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &req); err != nil {
		d.badRequest(models.Error{Error: err.Error()})
		return
	}
	if len(req.Name) == 0 {
		d.badRequest(models.Error{Error: "名称为空"})
		return
	}
	if code, err := service.DeleteService(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}
	d.success(models.Response{Action: "delete service successs", Status: "success"})
	return

}

// Post create a new Service.
// /v1/service/create
// @Title Create a Service
// @Description Create a Service
// @Param service body models.Service true
// @Success 200 {object} models.Response
// @Failure 400  连接超时
// @Failure 500  内部服务器错误
// @router /create [post]
func (d *CreateService) Post() {
	defer d.catchException()

	glog.V(2).Infof("add a service URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Service{}
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &req); err != nil {
		d.badRequest(models.Error{Error: err.Error()})
		return
	}
	if len(req.Name) == 0 {
		d.badRequest(models.Error{Error: "名称为空"})
		return
	}
	if len(req.Namespace) == 0 {
		d.badRequest(models.Error{Error: "命名空间为空"})
		return
	}
	if code, err := service.AddService(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}

	d.success(models.Response{Action: "add service successs", Status: "success"})
	return
}

// @router /create [get]
func (c *CreateService) Get() {
	if c.GetSession("username") == "root" {
		c.TplName = "createservice.html"
	} else {
		c.TplName = "404.html"
	}
}

// Get to get all service's information.
// /v1/service/getinfo
// @Title Get Services information
// @Description Get Service information
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /getjson [get]
func (c *GetService) Getjson() {
	serviceinfo := service.GetServiceinfo()
	c.Data["json"] = serviceinfo
	c.ServeJSON()
}

// @router /getinfo [get]
func (c *GetService) Get() {
	if c.GetSession("username") == "root" {
		c.TplName = "getservice.html"
	} else {
		c.TplName = "404.html"
	}
}
