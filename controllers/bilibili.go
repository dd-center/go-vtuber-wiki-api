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

// @router /live/:id/history
func (bc *BilibiliController) GetBilibiliLiveHistory() {
	defer bc.ServeJSON()
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	vtuber, vErr := models.GetVtuberById(bc.Ctx.Input.Params()[":id"])
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

// @router /live/:id/comments
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
	var chats []interface{}
	for _, comment := range comments {
		if comment.Type == 1 {
			chats = append(chats, struct {
				AuthorId    int64
				AuthorName  string
				Prefix      string
				PublishTime int64
				Content     string
			}{comment.AuthorId, comment.AuthorName,
				comment.Prefix, comment.PublishTime,
				comment.Content})
		}
	}
	bc.Data["json"] = &struct {
		Success      bool
		HasMoreItems bool
		TotalCount   int
		Comments     []interface{}
	}{true, offset+200 < len(chats), len(chats),
		chats[offset:int(math.Min(float64(offset+200), float64(len(chats)-1)))]}
}

// @router /live/:id/gifts
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
	var gifts []interface{}
	for _, comment := range comments {
		if comment.Type == 0 {
			gifts = append(gifts, struct {
				AuthorId    int64
				AuthorName  string
				PublishTime int64
				GiftName    string
				GiftCount   int
				CostType    string
				CostAmount  int
			}{comment.AuthorId, comment.AuthorName,
				comment.PublishTime, comment.GiftName,
				comment.GiftCount, comment.CostType, comment.CostAmount})
		}
	}
	bc.Data["json"] = &struct {
		Success      bool
		HasMoreItems bool
		TotalCount   int
		Gifts        []interface{}
	}{true, offset+200 < len(gifts), len(gifts),
		gifts[offset:int(math.Min(float64(offset+200), float64(len(gifts)-1)))]}
}
