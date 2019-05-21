package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"Slyph/models"
	"Slyph/models/pod"
)

// CreatePod struct
type CreatePod struct {
	baseController
}

// GetPod struct
type GetPod struct {
	baseController
}

// DeletePod struct
type DeletePod struct {
	baseController
}

// Get to get pods
// @Title Get pods
// @Description Get pods
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
func (d *DeletePod) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "deletepod.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post to delete a specify pod.
// /v1/pod/delete
// @Title Delete a pod
// @Description Delete a pod
// @Param pod body models.Delete true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /delete [post]
func (d *DeletePod) Post() {
	defer d.catchException()
	glog.V(2).Infof("delete a Pod URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
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
	if code, err := pod.PodDelete(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}
	d.success(models.Response{Action: "delete pod successs", Status: "success"})
	return
}

// Get to create a new pod.
// /v1/pod/create
// @router /create [get]
func (d *CreatePod) Get() {
	if d.GetSession("username") == "root" {
		d.TplName = "createpod.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post create a new pod.
// /v1/pod/create
// @Title Create a pod
// @Description Get 1 job's detail info
// @Param pod body models.Pod true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /create [post]
func (d *CreatePod) Post() {
	defer d.catchException()
	glog.V(2).Infof("add a Pod URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Pod{}
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
	if code, err := pod.PodAdd(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}

	d.success(models.Response{Action: "add pod successs", Status: "success"})
	return
}

// Getjson to get all pods's information.
// /v1/pod/getinfo
// @Title Get a podinfo
// @Description Get a pods info
// @Success 200 {object} models.PodList
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /getjson [get]
func (c *GetPod) Getjson() {
	podinfo := pod.Getpodinfo()
	c.Data["json"] = podinfo
	c.ServeJSON()
}

// Get to get all pods's information.
// /v1/pod/getinfo
// @router /getinfo [get]
func (c *GetPod) Get() {
	if c.GetSession("username") == "root" {
		c.TplName = "getpod.html"
	} else {
		c.TplName = "404.html"
	}
}
