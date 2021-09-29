package service

import (
	"github.com/maczh/gintool/mgresult"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"gopkg.in/mgo.v2/bson"
	"ququ.im/jizhangbao-data/model"
	"time"
)

const COLLECTION_BILL_OPERATE_LOG = "OperateLog"

func AddOperateLog(bookId, billId, operate string, inOut int, amount float64) mgresult.Result {
	operateLog := model.OperateLog{
		Id:      bson.NewObjectId(),
		Date:    utils.DateFormat(time.Now(), "yyyy-MM-dd"),
		Time:    utils.DateFormat(time.Now(), "hh:mm:ss"),
		BookId:  bookId,
		BillId:  billId,
		Operate: operate,
		InOut:   inOut,
		Amount:  amount,
	}
	config.Mgo.C(COLLECTION_BILL_OPERATE_LOG).Insert(operateLog)
	return *mgresult.Success(operateLog)
}

func ListOperateLog(bookId, startDate, endDate string) mgresult.Result {
	query := bson.M{"bookId": bookId}
	query["date"] = bson.M{"$gte": startDate, "$lte": endDate}
	sortFields := []string{"-date", "-time"}
	operateLogs := make([]model.OperateLog, 0)
	config.Mgo.C(COLLECTION_BILL_OPERATE_LOG).Find(query).Sort(sortFields...).All(&operateLogs)
	return *mgresult.Success(operateLogs)
}
