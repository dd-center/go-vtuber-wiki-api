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

func GetBiliLiveCommentsByTime(channelId int64, begin int64) ([]BiliLiveComment, error) {
	collation := Database.C("bili-live-comments")
	var result []BiliLiveComment
	err := collation.Find(bson.M{"masterId": channelId, "publishTime": bson.M{"$gte": begin}}).Limit(200).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FilterBiliChats(comments []BiliLiveComment) []interface{} {
	var chats []interface{}
	for _, comment := range comments {
		if comment.Type == 1 {
			chats = append(chats, struct {
				AuthorId    int64
				AuthorName  string
				Prefix      string
				PublishTime int64
				Content     string
				Popularity  int
			}{comment.AuthorId, comment.AuthorName,
				comment.Prefix, comment.PublishTime,
				comment.Content, comment.Popularity})
		}
	}
	return chats
}

func FilterBiliGifts(comments []BiliLiveComment) []interface{} {
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
				Popularity  int
			}{comment.AuthorId, comment.AuthorName,
				comment.PublishTime, comment.GiftName,
				comment.GiftCount, comment.CostType, comment.CostAmount, comment.Popularity})
		}
	}
	return gifts
}
