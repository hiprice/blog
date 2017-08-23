package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionUserInfo = "userInfos"
)

type UserInfo struct {
	ID         bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UserId     string        `json:"userId,omitempty" bson:"userId,omitempty"`
	Username   string        `json:"username,omitempty" bson:"username,omitempty"`
	Sex        string        `json:"sex,omitempty" bson:"sex,omitempty"`
	Age        string        `json:"age,omitempty" bson:"age,omitempty"`
	Occupation string        `json:"occupation,omitempty" bson:"occupation,omitempty"`
	Birthday   string        `json:"birthday,omitempty" bson:"image,omitempty"`
	Detail     string        `json:"detail,omitempty" bson:"detail,omitempty"`
}
