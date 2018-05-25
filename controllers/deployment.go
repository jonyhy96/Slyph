package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"k8s/models"
	"k8s/models/deployment"
)

// CreateDeployment struct
type CreateDeployment struct {
	baseController
}

// GetDeployment struct
type GetDeployment struct {
	baseController
}

// DeleteDeployment struct
type DeleteDeployment struct {
	baseController
}

// UpdateDeployment struct
type UpdateDeployment struct {
	baseController
}

// RollBackDeployment struct
type RollBackDeployment struct {
	baseController
}

var clientset = models.GetClientset()

// Post to delete the specify deployment.
// /v1/deployment/delete
// @Title Delete a deployment
// @Description Delete a deployment
// @Param deployment body models.Delete true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /delete [post]
func (d *DeleteDeployment) Post() {
	defer d.catchException()
	glog.V(2).Infof("delete a deployment URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Delete{}
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &req); err != nil {
		d.badRequest(models.Error{Error: err.Error()})
		return
	}
	if len(req.Name) == 0 {
		d.badRequest(models.Error{Error: "名称为空"})
		return
	}
	if code, err := deployment.Deletedeployment(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}
	d.success(models.Response{Action: "delete deployment successs", Status: "success"})
	return

}

// Post create a new deployment.
// /v1/deployment/create
// @Title Create a deployment
// @Description Create a deployment
// @Param deployment body models.Deployment true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /create [post]
func (d *CreateDeployment) Post() {
	defer d.catchException()
	glog.V(2).Infof("add a deployment URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Deployment{}
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

	if code, err := deployment.AddDeployment(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}

	d.success(models.Response{Action: "add deployment successs", Status: "success"})
	return
}

// Get info
// @router /create [get]
func (d *CreateDeployment) Get() {
	var name = d.GetSession("username")
	if name == "root" {
		d.TplName = "createdeployment.html"
	} else {
		d.TplName = "404.html"
	}
}

// Getjson to get all deployment's information.
// /v1/deployment/getinfo
// @Title get deployments information
// @Description get deployments information
// @Success 200 {object} models.DeploymentList
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /getjson [get]
func (c *GetDeployment) Getjson() {
	deploymentinfo := deployment.Getdeploymentinfo()
	c.Data["json"] = deploymentinfo
	c.ServeJSON()
}

// Get to get the get deeployment html
// @router /getinfo [get]
func (c *GetDeployment) Get() {
	var name = c.GetSession("username")
	if name == "root" {
		c.TplName = "getdeployment.html"
	} else {
		c.TplName = "404.html"
	}
}

// Post to update the specify deployment.
// /v1/deployment/update
// @Title update deployments
// @Description update deployments
// @Param deployment body models.Deployment true
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /update [post]
func (d *UpdateDeployment) Post() {
	defer d.catchException()
	glog.V(2).Infof("add a deployment URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Deployment{}
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &req); err != nil {
		d.badRequest(models.Error{Error: err.Error()})
		return
	}
	if len(req.Name) == 0 {
		d.badRequest(models.Error{Error: "名称为空"})
		return
	}
	if code, err := deployment.Updatedeployment(&req); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}

	d.success(models.Response{Action: "update deployment successs", Status: "success"})
	return
}

// @router /update [get]
func (d *UpdateDeployment) Get() {
	var name = d.GetSession("username")
	if name == "root" {
		d.TplName = "updatedeployment.html"
	} else {
		d.TplName = "404.html"
	}
}

// Post to rollback the specify deployment.
// /v1/deployment/rollback
// @Title rollback deployments
// @Description rollback deployments
// @Success 200 {object} models.Response
// @Failure 400  "连接超时"
// @Failure 500  "内部服务器错误"
// @router /rollback [post]
func (d *RollBackDeployment) Post() {
	defer d.catchException()
	glog.V(2).Infof("add a deployment URL[%s] body[%s]\n", d.Ctx.Input.URI(), string(d.Ctx.Input.RequestBody))
	req := models.Deployment{}
	if err := json.Unmarshal(d.Ctx.Input.RequestBody, &req); err != nil {
		d.badRequest(models.Error{Error: err.Error()})
		return
	}
	if len(req.Name) == 0 {
		d.badRequest(models.Error{Error: "名称为空"})
		return
	}
	if code, err := deployment.RollBack(req.Name); err != nil {
		d.customError(models.Error{Error: err.Error()}, code)
		return
	}

	d.success(models.Response{Action: "rollback deployment successs", Status: "success"})
	return
}
