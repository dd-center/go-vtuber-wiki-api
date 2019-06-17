package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type YoutubeLiveDetail struct {
	Id            string             "bson:\"_id\""
	Title         string             "bson:\"title\""
	ChannelId     string             "bson:\"channelId\""
	BeginTime     time.Time          "bson:\"beginTime\""
	EndTime       time.Time          "bson:\"endTime\""
	SuperchatInfo map[string]float32 "bson:\"superchatInfo\""
	ExchangeRate  map[string]float32 "bson:\"exchangeRate\""
}

type YoutubeLiveChat struct {
	Id               string    "bson:\"_id\""
	VideoId          string    "bson:\"videoId\""
	AuthorChannelId  string    "bson:\"authorChannelId\""
	PublishTime      time.Time "bson:\"publishedAt\""
	DisplayMessage   string    "bson:\"displayMessage\""
	Viewers          int       "bson:\"viewerCount\""
	SuperChatDetails string    "bson:\"superChatDetails\""
}

func GetLiveHistoryByChannelId(channelId string) ([]YoutubeLiveDetail, error) {
	collation := Database.C("youtube-live-details")
	var result []YoutubeLiveDetail
	err := collation.Find(bson.M{"channelId": channelId}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetLiveDetailByVideoId(videoId string) (*YoutubeLiveDetail, error) {
	collation := Database.C("youtube-live-details")
	var result YoutubeLiveDetail
	err := collation.Find(bson.M{"_id": videoId}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetLiveChatsByVideoId(videoId string) ([]YoutubeLiveChat, error) {
	collation := Database.C("youtube-live-chats")
	var result []YoutubeLiveChat
	err := collation.Find(bson.M{"videoId": videoId}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
