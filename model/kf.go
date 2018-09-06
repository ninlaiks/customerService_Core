package model

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

// 客服模型
type Kf struct {
	Id         string    `json:"id" bson:"id"`
	TokenId    string    `json:"token_id" bson:"token_id"`
	NickName   string    `json:"nick_name" bson:"nick_name"`
	Type       int       `json:"type" bson:"type"`
	HeadImgUrl string    `json:"head_img_url" bson:"head_img_url"`
	IsOnline   bool      `json:"is_online" bson:"is_online"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}

// 指定在线客服是否存在
func (k Kf) OnlineExist() bool {
	var kefuC = Db.C("kefu")

	if count, err := kefuC.Find(bson.M{"id": k.Id, "is_online": true}).Count(); err != nil {
		log.Printf("model.Kf.Exist() is err :%s", err.Error())
		return false
	} else {
		if count > 0 {
			return true
		} else {
			return false
		}
	}
}

// 获取所有在线客服
func (k Kf) QueryOnlines() ([]*Kf, error) {
	var (
		err     error
		onlines []*Kf
		kefuC   = Db.C("kefu")
	)
	if err = kefuC.Find(bson.M{"is_online": true}).All(&onlines); err != nil {
		log.Printf("model.QueryOnlines is err: %s", err.Error())
	}

	return onlines, err
}

// 修改客服在线状态
func (k Kf) ChangeStatus() (err error) {
	kefuC := Db.C("kefu")

	if err = kefuC.Update(bson.M{"id": k.Id}, bson.M{"$set": bson.M{"is_online": k.IsOnline, "update_time": time.Now()}}); err != nil {
		log.Printf("model.ChangeStatus is err: %s", err.Error())
	}
	return
}
