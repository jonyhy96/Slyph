package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Slyph/controllers:AddStorage"] = append(beego.GlobalControllerRouter["Slyph/controllers:AddStorage"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:AddStorage"] = append(beego.GlobalControllerRouter["Slyph/controllers:AddStorage"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:AutoBuildController"] = append(beego.GlobalControllerRouter["Slyph/controllers:AutoBuildController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/up`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:AutoBuildController"] = append(beego.GlobalControllerRouter["Slyph/controllers:AutoBuildController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/up`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:AutoDownController"] = append(beego.GlobalControllerRouter["Slyph/controllers:AutoDownController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/down`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:AutoDownController"] = append(beego.GlobalControllerRouter["Slyph/controllers:AutoDownController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/down`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreateDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreateDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreateDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreateDeployment"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreatePod"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreatePod"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreatePod"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreatePod"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreateService"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreateService"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreateService"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreateService"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreateStatefulset"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreateStatefulset"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:CreateStatefulset"] = append(beego.GlobalControllerRouter["Slyph/controllers:CreateStatefulset"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:DeleteDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:DeleteDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:DeletePod"] = append(beego.GlobalControllerRouter["Slyph/controllers:DeletePod"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:DeleteService"] = append(beego.GlobalControllerRouter["Slyph/controllers:DeleteService"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:DeleteStatefulset"] = append(beego.GlobalControllerRouter["Slyph/controllers:DeleteStatefulset"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:DeleteStorage"] = append(beego.GlobalControllerRouter["Slyph/controllers:DeleteStorage"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetDeployment"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetDeployment"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetPod"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetPod"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetPod"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetPod"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetService"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetService"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetService"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetService"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetStatefulset"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetStatefulset"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetStatefulset"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetStatefulset"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetStorage"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetStorage"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/getinfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:GetStorage"] = append(beego.GlobalControllerRouter["Slyph/controllers:GetStorage"],
		beego.ControllerComments{
			Method: "Getjson",
			Router: `/getjson`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:IndexController"] = append(beego.GlobalControllerRouter["Slyph/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:LogController"] = append(beego.GlobalControllerRouter["Slyph/controllers:LogController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:LogController"] = append(beego.GlobalControllerRouter["Slyph/controllers:LogController"],
		beego.ControllerComments{
			Method: "Getlog",
			Router: `/log`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:LoginController"] = append(beego.GlobalControllerRouter["Slyph/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:LoginController"] = append(beego.GlobalControllerRouter["Slyph/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:RollBackDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:RollBackDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/rollback`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:UpdateDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:UpdateDeployment"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Slyph/controllers:UpdateDeployment"] = append(beego.GlobalControllerRouter["Slyph/controllers:UpdateDeployment"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/update`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
