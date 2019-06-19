package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"],
		beego.ControllerComments{
			Method:           "GetBilibiliLiveDanmakuByTime",
			Router:           `/:id/danmaku`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"],
		beego.ControllerComments{
			Method:           "GetBilibiliLiveCommentById",
			Router:           `/live/:id/comments`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"],
		beego.ControllerComments{
			Method:           "GetBilibiliLiveGiftById",
			Router:           `/live/:id/gifts`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:BilibiliController"],
		beego.ControllerComments{
			Method:           "GetBilibiliLiveHistory",
			Router:           `/live/:id/history`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:VtuberController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:VtuberController"],
		beego.ControllerComments{
			Method:           "GetAllVtubers",
			Router:           `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:VtuberController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:VtuberController"],
		beego.ControllerComments{
			Method:           "SearchVtuber",
			Router:           `/search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"],
		beego.ControllerComments{
			Method:           "GetYoutubeLiveChats",
			Router:           `/live/:id/chats`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"],
		beego.ControllerComments{
			Method:           "GetYoutubeLiveDetail",
			Router:           `/live/:id/details`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"],
		beego.ControllerComments{
			Method:           "GetYoutubeLiveViewersDetail",
			Router:           `/live/:id/details/viewers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"],
		beego.ControllerComments{
			Method:           "GetYoutubeLiveHistory",
			Router:           `/live/:id/history`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"] = append(beego.GlobalControllerRouter["go-vtuber-wiki-api/controllers:YoutubeController"],
		beego.ControllerComments{
			Method:           "GetYoutubeSuperchats",
			Router:           `/live/:id/superchats`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
