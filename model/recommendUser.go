package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

type RecommendUser struct {
	Id      int       `json:"id" xorm:"'id'"`
	UserID  int       `json:"user_id" xorm:"'userid'"`
	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
}
type RecommendUserRepository struct {
	engine *xorm.Engine
}

func NewRecommendUserRepository() RecommendUserRepository {

	return RecommendUserRepository{
		engine: getEngine(),
	}
}
func (self RecommendUserRepository) Create(recommendUser RecommendUser) (int64, error) {
	return self.engine.Insert(recommendUser)

}
