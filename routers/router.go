// @APIVersion 1.0.0
// @Title Vtuber wiki api
// @Description Vtuber wiki api后端GO重制版
// @Contact mrs4sxiaoshi@gmail.com
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
