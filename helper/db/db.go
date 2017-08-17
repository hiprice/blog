package db

import (
	"gopkg.in/mgo.v2"
)

var (
	MongoSession       *mgo.Session
	UserInfoCollection *mgo.Collection
)
