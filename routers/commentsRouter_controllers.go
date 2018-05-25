package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["k8s/controllers:AddStorage"] = append(beego.GlobalControllerRouter["k8s/controllers:AddStorage"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:AddStorage"] = append(beego.GlobalControllerRouter["k8s/controllers:AddStorage"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:AutoBuildController"] = append(beego.GlobalControllerRouter["k8s/controllers:AutoBuildController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/up`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:AutoBuildController"] = append(beego.GlobalControllerRouter["k8s/controllers:AutoBuildController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/up`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:AutoDownController"] = append(beego.GlobalControllerRouter["k8s/controllers:AutoDownController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/down`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:AutoDownController"] = append(beego.GlobalControllerRouter["k8s/controllers:AutoDownController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/down`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreateDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:CreateDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreateDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:CreateDeployment"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreatePod"] = append(beego.GlobalControllerRouter["k8s/controllers:CreatePod"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreatePod"] = append(beego.GlobalControllerRouter["k8s/controllers:CreatePod"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreateService"] = append(beego.GlobalControllerRouter["k8s/controllers:CreateService"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreateService"] = append(beego.GlobalControllerRouter["k8s/controllers:CreateService"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreateStatefulset"] = append(beego.GlobalControllerRouter["k8s/controllers:CreateStatefulset"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:CreateStatefulset"] = append(beego.GlobalControllerRouter["k8s/controllers:CreateStatefulset"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:DeleteDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:DeleteDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:DeletePod"] = append(beego.GlobalControllerRouter["k8s/controllers:DeletePod"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:DeleteService"] = append(beego.GlobalControllerRouter["k8s/controllers:DeleteService"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:DeleteStatefulset"] = append(beego.GlobalControllerRouter["k8s/controllers:DeleteStatefulset"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:DeleteStorage"] = append(beego.GlobalControllerRouter["k8s/controllers:DeleteStorage"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:GetDeployment"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:GetDeployment"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetPod"] = append(beego.GlobalControllerRouter["k8s/controllers:GetPod"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetPod"] = append(beego.GlobalControllerRouter["k8s/controllers:GetPod"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetService"] = append(beego.GlobalControllerRouter["k8s/controllers:GetService"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetService"] = append(beego.GlobalControllerRouter["k8s/controllers:GetService"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetStatefulset"] = append(beego.GlobalControllerRouter["k8s/controllers:GetStatefulset"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetStatefulset"] = append(beego.GlobalControllerRouter["k8s/controllers:GetStatefulset"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetStorage"] = append(beego.GlobalControllerRouter["k8s/controllers:GetStorage"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:GetStorage"] = append(beego.GlobalControllerRouter["k8s/controllers:GetStorage"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:IndexController"] = append(beego.GlobalControllerRouter["k8s/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:LogController"] = append(beego.GlobalControllerRouter["k8s/controllers:LogController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:LogController"] = append(beego.GlobalControllerRouter["k8s/controllers:LogController"],
		beego.ControllerComments{
			Method: "Getlog",
			Router: `/log`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:LoginController"] = append(beego.GlobalControllerRouter["k8s/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:LoginController"] = append(beego.GlobalControllerRouter["k8s/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:RollBackDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:RollBackDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/rollback`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:UpdateDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:UpdateDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["k8s/controllers:UpdateDeployment"] = append(beego.GlobalControllerRouter["k8s/controllers:UpdateDeployment"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/update`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
