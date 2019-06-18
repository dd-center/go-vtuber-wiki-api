package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BiliLiveDetail struct {
	Id            string    "bson:\"_id\""
	ChannelId     int64     "bson:\"channelId\""
	Title         string    "bson:\"title\""
	BeginTime     time.Time "bson:\"beginTime\""
	EndTime       time.Time "bson:\"endTime\""
	MaxPopularity int       "bson:\"maxPopularity\""
}

type BiliLiveComment struct {
	Id          bson.ObjectId "bson:\"_id\""
	StreamerId  int64         "bson:\"masterId\""
	PublishTime int64         "bson:\"publishTime\""
	Type        int           "bson:\"type\""
	Prefix      string        "bson:\"suffix\"" //这地方字段名写错了。。将就
	Content     string        "bson:\"content\""
	AuthorName  string        "bson:\"fromUsername\""
	AuthorId    int64         "bson:\"fromUserid\""
	GiftName    string        "bson:\"giftName\""
	GiftCount   int           "bson:\"giftCount\""
	CostType    string        "bson:\"costType\""
	CostAmount  int           "bson:\"cost\""
	Popularity  int           "bson:\"popularity\""
}

func GetBiliLiveHistoryByUid(uid uint64) ([]BiliLiveDetail, error) {
	collation := Database.C("bili-live-details")
	var result []BiliLiveDetail
	err := collation.Find(bson.M{"channelId": uid}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetBiliLiveDetailById(id string) (*BiliLiveDetail, error) {
	collation := Database.C("bili-live-details")
	var result BiliLiveDetail
	err := collation.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetBiliLiveCommentsById(id string) ([]BiliLiveComment, error) {
	collation := Database.C("bili-live-comments")
	detail, liveErr := GetBiliLiveDetailById(id)
	if liveErr != nil {
		return nil, liveErr
	}
	var result []BiliLiveComment
	err := collation.Find(bson.M{"masterId": detail.ChannelId, "publishTime": bson.M{"$gte": detail.BeginTime.Unix(), "$lte": detail.EndTime.Unix()}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
