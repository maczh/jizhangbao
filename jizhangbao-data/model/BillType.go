package model

import "gopkg.in/mgo.v2/bson"

type BillType struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	BillTypeId     string        `json:"billTypeId" bson:"billTypeId"`         //账单分类id
	BillType       string        `json:"billType" bson:"billType"`             //账单分类名称
	BillTypeIconId string        `json:"billTypeIconId" bson:"billTypeIconId"` //账单分类图标id
	BookId         string        `json:"bookId" bson:"bookId"`                 //账本id
	OperateType    string        `json:"operateType" bson:"operateType"`       //收入/支出/应收/应付
	IsPublic       bool          `json:"isPublic" bson:"isPublic"`             //是否公共分类
	UpdateTime     string        `json:"updateTime" bson:"updateTime"`         //更新时间，yyyy-MM-dd HH:mm:ss
}
