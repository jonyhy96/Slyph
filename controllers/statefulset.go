package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"k8s/models"
	"k8s/models/statefulset"
)

// CreateStatefulset struct
type CreateStatefulset struct {
	baseController
}

// GetStatefulset struct
type GetStatefulset struct {
	baseController
}

// DeleteStatefulset struct
type DeleteStatefulset struct {
	baseController
}

// Get to get statefulsets
// @Title Get statefulsets
// @Description Get statefulsets
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
func (d *DeleteStatefulset) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "deleteStatefulset.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post to delete a specify statefulset.
// /v1/st/delete
// @Title Delete a statefulset
// @Description Delete a statefulset
// @Param statefulset body models.Delete true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /delete [post]
func (d *DeleteStatefulset) Post() {
	defer d.catchException()
	glog.V(2).Infof("delete a Statefulset URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Delete{}
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
	if code, err := statefulset.DeleteStatefulset(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}
	d.success(models.Response{Action: "delete Statefulset successs", Status: "success"})
	return
}

// Get to create a new statefulset.
// /v1/pod/create
// @router /create [get]
func (d *CreateStatefulset) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "createst.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post create a new statefulset.
// /v1/st/create
// @Title Create a statefulset
// @Description Get 1 statefulset's detail info
// @Param statefulset body models.Statefulset true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /create [post]
func (d *CreateStatefulset) Post() {
	defer d.catchException()
	glog.V(2).Infof("add a Statefulset URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Statefulset{}
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
	if code, err := statefulset.AddStatefulset(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}

	d.success(models.Response{Action: "add Statefulset successs", Status: "success"})
	return
}

// Getjson to get all pods's information.
// /v1/st/getinfo
// @Title Get a statefulset info
// @Description Get a statefulsets info
// @Success 200 {object} models.StatefulSetList{}
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /getjson [get]
func (c *GetStatefulset) Getjson() {
	Statefulsetinfo := statefulset.GetStatefulset()
	c.Data["json"] = Statefulsetinfo
	c.ServeJSON()
}

// Get to get all statefulsets's information.
// /v1/st/getinfo
// @router /getinfo [get]
func (c *GetStatefulset) Get() {
	if c.GetSession("username") == "root" {
		c.TplName = "getst.html"
	} else {
		c.TplName = "404.html"
	}
}
