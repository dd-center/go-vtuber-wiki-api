package models

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type VtuberEntity struct {
	Id               bson.ObjectId "bson:\"_id\",json:\"id\""
	OriginalName     string        "bson:\"originalName\""
	ChineseName      string        "bson:\"chineseName\""
	YoutubeChannelId string        "bson:\"youtubeChannelId\""
	TwitterProfileId string        "bson:\"twitterProfileId\""
	BiliUid          int64         "bson:\"bilibiliUid\""
	GroupName        string        "bson:\"groupName\""
	NickNames        []string      "bson:\"nickNameList\""
}

func GetAllVtubers() ([]VtuberEntity, error) {
	collation := Database.C("vtubers")
	var result []VtuberEntity
	err := collation.Find(bson.M{}).All(&result)
	if err != nil {
		fmt.Println("cannot find all vtubers: " + err.Error())
		return nil, err
	}
	return result, nil
}

func GetVtuberById(id string) (*VtuberEntity, error) {
	collation := Database.C("vtubers")
	var result VtuberEntity
	err := collation.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		fmt.Println("cannot find  vtubers " + id + " :" + err.Error())
		return nil, err
	}
	return &result, nil
}

func SearchVtuber(keyword string) (*VtuberEntity, error) {
	vtubers, err := GetAllVtubers()
	if err != nil {
		return nil, err
	}
	for _, vtuber := range vtubers {
		if strings.Contains(vtuber.OriginalName, keyword) || Contains(keyword, vtuber.NickNames) {
			return &vtuber, nil
		}
	}
	return nil, errors.New("Vtuber not found.")
}


