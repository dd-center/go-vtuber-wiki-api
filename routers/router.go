// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-vtuber-wiki-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/vtuber",
			beego.NSInclude(
				&controllers.VtuberController{},
			),
		),
		beego.NSNamespace("/youtube",
			beego.NSInclude(
				&controllers.YoutubeController{},
			),
		),
		beego.NSNamespace("/bilibili",
			beego.NSInclude(
				&controllers.BilibiliController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
