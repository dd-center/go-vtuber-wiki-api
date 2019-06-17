package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"go-vtuber-wiki-api/models"
	"math"
	"strconv"
)

type YoutubeController struct {
	beego.Controller
}

// @router /live/:id/history
func (yc *YoutubeController) GetYoutubeLiveHistory() {
	type LiveDetailResult struct {
		Id        string
		Title     string
		BeginTime int64
		EndTime   int64
	}
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	vtuberId := yc.Ctx.Input.Params()[":id"]
	vtuberInfo, err := models.GetVtuberById(vtuberId)
	if err != nil || vtuberInfo.YoutubeChannelId == "" {
		errorTemplate.Message = "Vtuber not found."
		yc.Data["json"] = errorTemplate
		yc.ServeJSON()
		return
	}
	lives, e := models.GetLiveHistoryByChannelId(vtuberInfo.YoutubeChannelId)
	if e != nil {
		errorTemplate.Message = e.Error()
		yc.Data["json"] = errorTemplate
		yc.ServeJSON()
		return
	}
	var result []LiveDetailResult
	for _, liveDetail := range lives {
		result = append(result, LiveDetailResult{
			Id:        liveDetail.Id,
			Title:     liveDetail.Title,
			BeginTime: liveDetail.BeginTime.Unix(),
			EndTime:   liveDetail.EndTime.Unix(),
		})
	}
	yc.Data["json"] = &struct {
		Success bool
		Lives   []LiveDetailResult
	}{true, result}
	yc.ServeJSON()
}

// @router /live/:id/details
func (yc *YoutubeController) GetYoutubeLiveDetail() {
	type LiveDetailResult struct {
		Id             string
		Title          string
		BeginTime      int64
		EndTime        int64
		Amount         float32
		LiveChatCount  int
		SuperchatCount int
		SuperchatInfo  map[string]float32
		ExchangeRate   map[string]float32
	}
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	videoId := yc.Ctx.Input.Params()[":id"]
	liveDetail, detailErr := models.GetLiveDetailByVideoId(videoId)
	liveChats, chatErr := models.GetLiveChatsByVideoId(videoId)
	if detailErr != nil || chatErr != nil {
		errorTemplate.Message = "cannot get live detail."
		yc.Data["json"] = errorTemplate
		yc.ServeJSON()
		return
	}
	result := LiveDetailResult{
		Id:             liveDetail.Id,
		Title:          liveDetail.Title,
		BeginTime:      liveDetail.BeginTime.Unix(),
		EndTime:        liveDetail.EndTime.Unix(),
		Amount:         0,
		LiveChatCount:  len(liveChats),
		SuperchatCount: 0,
		SuperchatInfo:  liveDetail.SuperchatInfo,
		ExchangeRate:   liveDetail.ExchangeRate,
	}
	for k, v := range liveDetail.SuperchatInfo {
		rate := liveDetail.ExchangeRate[k]
		result.Amount += rate * v
	}
	for _, chat := range liveChats {
		if chat.SuperChatDetails != "" {
			result.SuperchatCount++
		}
	}
	yc.Data["json"] = &struct {
		Success bool
		Detail  LiveDetailResult
	}{true, result}
	yc.ServeJSON()
}

// @router /live/:id/details/viewers
func (yc *YoutubeController) GetYoutubeLiveViewersDetail() {
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	liveChats, chatErr := models.GetLiveChatsByVideoId(yc.Ctx.Input.Params()[":id"])
	liveDetails, detailErr := models.GetLiveDetailByVideoId(yc.Ctx.Input.Params()[":id"])
	if chatErr != nil || detailErr != nil {
		errorTemplate.Message = "cannot get live chats info."
		yc.Data["json"] = errorTemplate
		yc.ServeJSON()
		return
	}
	// 统计观众曲线
	var viewersTrend []interface{}
	sum := 0
	for i, chat := range liveChats {
		sum += chat.Viewers
		avg := sum / (i + 1)
		if (i >= 1 && liveChats[i-1].Viewers != chat.Viewers && chat.Viewers-avg < 2000) || i == 0 {
			viewersTrend = append(viewersTrend, struct {
				Time    int64
				Viewers int
			}{chat.PublishTime.Unix(), chat.Viewers})
		}
	}
	// 统计初见观众
	var firstRate float32
	liveHistory, _ := models.GetLiveHistoryByChannelId(liveDetails.ChannelId)
	if len(liveHistory) > 6 {
		var publishedViewers []string
		var unPublishedViewers []string
		count := 0
		for i := len(liveHistory) - 1; count < 3; i-- {
			if liveHistory[i].BeginTime.Unix() >= liveDetails.BeginTime.Unix() {
				continue
			}
			count++
			hisChats, _ := models.GetLiveChatsByVideoId(liveHistory[i].Id)
			for _, hisChat := range hisChats {
				if !models.Contains(hisChat.AuthorChannelId, publishedViewers) {
					publishedViewers = append(publishedViewers, hisChat.AuthorChannelId)
				}
			}
		}
		for _, chat := range liveChats {
			if !models.Contains(chat.AuthorChannelId, publishedViewers) {
				unPublishedViewers = append(unPublishedViewers, chat.AuthorChannelId)
			}
		}
		firstRate = float32(len(unPublishedViewers)) / float32(len(liveChats))
	} else {
		firstRate = 0
	}
	yc.Data["json"] = &struct {
		Success      bool
		ViewersTrend []interface{}
		FirstRate    float32
	}{true, viewersTrend, firstRate * 100}
	yc.ServeJSON()
}

// @router /live/:id/chats
func (yc *YoutubeController) GetYoutubeLiveChats() {
	type LiveChatDetail struct {
		AuthorChannelId string
		Message         string
		PublishTime     int64
	}
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	liveChats, chatErr := models.GetLiveChatsByVideoId(yc.Ctx.Input.Params()[":id"])
	if chatErr != nil {
		errorTemplate.Message = "cannot get live chats."
		yc.Data["json"] = errorTemplate
		yc.ServeJSON()
		return
	}
	var commonChats []LiveChatDetail
	for _, chat := range liveChats {
		if chat.SuperChatDetails == "" {
			commonChats = append(commonChats, LiveChatDetail{
				AuthorChannelId: chat.AuthorChannelId,
				Message:         chat.DisplayMessage,
				PublishTime:     chat.PublishTime.Unix(),
			})
		}
	}
	offset, parseErr := strconv.Atoi(yc.Input().Get("offset"))
	if parseErr != nil || offset > len(commonChats) {
		errorTemplate.Message = "offset error."
		yc.Data["json"] = errorTemplate
		yc.ServeJSON()
		return
	}
	yc.Data["json"] = &struct {
		Success     bool
		HasMoreItem bool
		LiveChats   []LiveChatDetail
	}{true, offset+200 < len(commonChats), commonChats[offset:int(math.Min(float64(offset+200), float64(len(commonChats)-1)))]}
	yc.ServeJSON()
}

// @router /live/:id/superchats
func (yc *YoutubeController) GetYoutubeSuperchats() {
	type SuperchatDetail struct {
		AuthorChannelId string
		Message         string
		PublishTime     int64
		CostName        string
		Amount          int32
	}
	type SuperchatPayDetail struct {
		AmountMicros string `json:"amountMicros"`
		Currency     string `json:"currency"`
		Comment      string `json:"userComment"`
	}
	errorTemplate := &struct {
		Success bool
		Message string
	}{false, ""}
	liveChats, chatErr := models.GetLiveChatsByVideoId(yc.Ctx.Input.Params()[":id"])
	if chatErr != nil {
		errorTemplate.Message = "cannot get live chats."
		yc.Data["json"] = errorTemplate
		yc.ServeJSON()
		return
	}
	var superchats []SuperchatDetail
	for _, chat := range liveChats {
		if chat.SuperChatDetails != "" {
			var payDetail SuperchatPayDetail
			_ = json.Unmarshal([]byte(chat.SuperChatDetails), &payDetail)
			amount, _ := strconv.ParseFloat(payDetail.AmountMicros, 32)
			superchats = append(superchats, SuperchatDetail{
				AuthorChannelId: chat.AuthorChannelId,
				Message:         payDetail.Comment,
				PublishTime:     chat.PublishTime.Unix(),
				CostName:        payDetail.Currency,
				Amount:          int32(amount / 1000000),
			})
		}
	}
	yc.Data["json"] = &struct {
		Success    bool
		Superchats []SuperchatDetail
	}{true, superchats}
	yc.ServeJSON()
}
