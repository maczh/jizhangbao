package service

import (
	"github.com/maczh/gintool/mgresult"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"gopkg.in/mgo.v2/bson"
	"ququ.im/jizhangbao-data/model"
	"time"
)

const COLLECTION_BILL_TYPE = "BillType"

func getBillType(billTypeId string) model.BillType {
	var billType model.BillType
	config.Mgo.C(COLLECTION_BILL_TYPE).Find(&bson.M{"billTypeId": billTypeId}).One(&billType)
	return billType
}

func GetBillType(billTypeId string) mgresult.Result {
	return *mgresult.Success(getBillType(billTypeId))
}

func listBillTypes(bookId, operateType string) []model.BillType {
	billTypes := make([]model.BillType, 0)
	config.Mgo.C(COLLECTION_BILL_TYPE).Find(&bson.M{"$or": []bson.M{{"isPublic": true, "operateType": operateType}, {"bookId": bookId, "operateType": operateType}}}).All(&billTypes)
	return billTypes
}

func AddPublicBillType(billType, billTypeIconId, operateType string) mgresult.Result {
	billtype := model.BillType{
		Id:             bson.NewObjectId(),
		BillTypeId:     utils.GetUUIDString(),
		BillType:       billType,
		BillTypeIconId: billTypeIconId,
		OperateType:    operateType,
		IsPublic:       true,
		UpdateTime:     utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
	}
	config.Mgo.C(COLLECTION_BILL_TYPE).Insert(billtype)
	return *mgresult.Success(billtype)
}

func AddBillType(bookId, billType, billTypeIconId, operateType string) mgresult.Result {
	billtype := model.BillType{
		Id:             bson.NewObjectId(),
		BookId:         bookId,
		BillTypeId:     utils.GetUUIDString(),
		BillType:       billType,
		BillTypeIconId: billTypeIconId,
		OperateType:    operateType,
		IsPublic:       false,
		UpdateTime:     utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
	}
	config.Mgo.C(COLLECTION_BILL_TYPE).Insert(billtype)
	return *mgresult.Success(billtype)
}

func UpdateBillType(bookId, billTypeId, billType, billTypeIconId string) mgresult.Result {
	billtype := getBillType(billTypeId)
	if billtype.BillTypeId == "" {
		return *mgresult.Error(-1, "无此账单类型")
	}
	if billtype.IsPublic {
		return *mgresult.Error(-1, "不可更改系统账单类型")
	}
	if billType != "" {
		billtype.BillType = billType
	}
	if billTypeIconId != "" {
		billtype.BillTypeIconId = billTypeIconId
	}
	billtype.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	config.Mgo.C(COLLECTION_BILL_TYPE).UpdateId(billtype.Id, billtype)
	return *mgresult.Success(billtype)
}

func DeleteBillType(bookId, billTypeId string) mgresult.Result {
	config.Mgo.C(COLLECTION_BILL_TYPE).Remove(&bson.M{"bookId": bookId, "billTypeId": billTypeId})
	return *mgresult.Success(nil)
}

func ListBillTypes(bookId, operateType string) mgresult.Result {
	return *mgresult.Success(listBillTypes(bookId, operateType))
}
