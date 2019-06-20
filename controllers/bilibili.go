package controllers

import (
	"github.com/astaxie/beego"
	"go-vtuber-wiki-api/models"
	"math"
	"strconv"
)

type BilibiliController struct {
	beego.Controller
}

// @Title 获取bilibili直播历史
// @Description 通过VtuberId获取目标的直播历史
// @Param id path string true "VtuberId"
// @router /live/:id/history [get]
func (bc *BilibiliController) GetBilibiliLiveHistory() {
	defer bc.ServeJSON()
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	vtuber, vErr := models.GetVtuberByVdbId(bc.Ctx.Input.Params()[":id"])
	if vErr != nil {
		errorTemplate.Message = "vtuber id error."
		bc.Data["json"] = errorTemplate
		return
	}
	history, dbErr := models.GetBiliLiveHistoryByUid(vtuber.BiliUid)
	if dbErr != nil {
		errorTemplate.Message = "cannot find history."
		bc.Data["json"] = errorTemplate
		return
	}
	var result []interface{}
	var liveTime int64
	for _, live := range history {
		liveTime += live.EndTime.Unix() - live.BeginTime.Unix()
		result = append(result, struct {
			Id            string
			Title         string
			BeginTime     int64
			EndTime       int64
			MaxPopularity int
		}{live.Id, live.Title, live.BeginTime.Unix(), live.EndTime.Unix(), live.MaxPopularity})
	}
	bc.Data["json"] = &struct {
		Success  bool
		LiveTime int64
		Lives    []interface{}
	}{true, liveTime, result}
}

// @Title 获取bilibili直播中的聊天
// @Description 通过LiveId获取该场直播的聊天弹幕
// @Param id path string true "Bilibili live id"
// @Param offset query int32 true "返回偏移"
// @router /live/:id/comments [get]
func (bc *BilibiliController) GetBilibiliLiveCommentById() {
	defer bc.ServeJSON()
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	comments, err := models.GetBiliLiveCommentsById(bc.Ctx.Input.Params()[":id"])
	offset, parseErr := strconv.Atoi(bc.Input().Get("offset"))
	if err != nil {
		errorTemplate.Message = err.Error()
		bc.Data["json"] = errorTemplate
		return
	}
	if parseErr != nil {
		errorTemplate.Message = "offset error."
		bc.Data["json"] = errorTemplate
		return
	}
	chats := models.FilterBiliChats(comments)
	bc.Data["json"] = &struct {
		Success      bool
		HasMoreItems bool
		TotalCount   int
		Comments     []interface{}
	}{true, offset+200 < len(chats), len(chats),
		chats[offset:int(math.Min(float64(offset+200), float64(len(chats)-1)))]}
}

// @Title 获取bilibili直播中的礼物
// @Description 通过LiveId获取该场直播的礼物弹幕
// @Param id path string true "Bilibili live id"
// @Param offset query int32 true "返回偏移"
// @router /live/:id/gifts [get]
func (bc *BilibiliController) GetBilibiliLiveGiftById() {
	defer bc.ServeJSON()
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	comments, err := models.GetBiliLiveCommentsById(bc.Ctx.Input.Params()[":id"])
	offset, parseErr := strconv.Atoi(bc.Input().Get("offset"))
	if err != nil {
		errorTemplate.Message = err.Error()
		bc.Data["json"] = errorTemplate
		return
	}
	if parseErr != nil {
		errorTemplate.Message = "offset error."
		bc.Data["json"] = errorTemplate
		return
	}
	gifts := models.FilterBiliGifts(comments)
	bc.Data["json"] = &struct {
		Success      bool
		HasMoreItems bool
		TotalCount   int
		Gifts        []interface{}
	}{true, offset+200 < len(gifts), len(gifts),
		gifts[offset:int(math.Min(float64(offset+200), float64(len(gifts)-1)))]}
}

// @Title 获取bilibili弹幕
// @Description 通过时间获取该Vtuber的200条直播弹幕
// @Param id path string true "VtuberId"
// @Param time query int32 true "Unix时间戳"
// @router /:id/danmaku [get]
func (bc *BilibiliController) GetBilibiliLiveDanmakuByTime() {
	defer bc.ServeJSON()
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	vtuber, vErr := models.GetVtuberByVdbId(bc.Ctx.Input.Params()[":id"])
	time, timeErr := strconv.ParseInt(bc.Input().Get("time"), 0, 64)
	if vErr != nil || timeErr != nil {
		errorTemplate.Message = "input error."
		bc.Data["json"] = errorTemplate
		return
	}
	comments, err := models.GetBiliLiveCommentsByTime(int64(vtuber.BiliUid), time)
	if err != nil {
		errorTemplate.Message = err.Error()
		bc.Data["json"] = errorTemplate
		return
	}
	bc.Data["json"] = &struct {
		Success  bool
		Comments []interface{}
		Gifts    []interface{}
	}{true, models.FilterBiliChats(comments), models.FilterBiliGifts(comments)}

}
