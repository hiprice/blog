package model

import (
	db "github.com/hiprice/blog/helper/db"
	"gopkg.in/mgo.v2/bson"
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

// NewUser ...
func NewUserInfo() UserInfo {
	return UserInfo{}
}

//SaveToDB s
func (self *UserInfo) SaveUserInfo() error {
	err := db.UserInfoCollection.Insert(&self)

	if err != nil {
		return err
	}
	return nil
}

//ReadFromDB r
func (u *UserInfo) ReadFromDB() ([]UserInfo, error) {
	result := []UserInfo{}
	err := db.UserInfoCollection.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//ReadByID only one
func (u *UserInfo) ReadByID() (*UserInfo, error) {
	err := db.UserInfoCollection.Find(bson.M{"userId": u.UserId}).One(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
