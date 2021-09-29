package service

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"github.com/snabb/isoweek"
	"gopkg.in/mgo.v2/bson"
	"ququ.im/jizhangbao-data/model"
	"strconv"
	"time"
)

const COLLECTION_BILL_FLOW = "BillFlow"

var OperateInfo = map[int]string{1: "收入", 2: "支出", 3: "应收", 4: "应付"}

func AddBillFlow(bookId, billTypeId, billAccountId, customerId, remark, photos, accrualId, billTime string, inOut int, amount float64) mgresult.Result {
	billFlow := model.BillFlow{
		Id:                 bson.NewObjectId(),
		BookId:             bookId,
		BillId:             utils.GetUUIDString(),
		BillTypeId:         billTypeId,
		InOut:              inOut,
		BillTime:           billTime,
		UpdateTime:         utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
		Amount:             amount,
		BillAccountId:      billAccountId,
		BusinessCustomerId: customerId,
		Remark:             remark,
		BillAccrualId:      accrualId,
	}
	if billFlow.BillTime == "" {
		billFlow.BillTime = billFlow.UpdateTime
	}
	var photoList []string
	if photos != "" && photos[:1] != "[" {
		utils.FromJSON(photos, &photoList)
	}
	billFlow.Photos = photoList
	billType := getBillType(billTypeId)
	billFlow.BillType = billType.BillType
	billFlow.BillTypeIconId = billType.BillTypeIconId
	billAccount := getBillAccount(bookId, billAccountId)
	billFlow.BillAccount = billAccount.BillAccount
	if customerId != "" {
		businessCustomer := getBusinessCustomer(bookId, customerId)
		billFlow.BusinessCustomer = businessCustomer.BusinessCustomer
	}
	config.Mgo.C(COLLECTION_BILL_FLOW).Insert(billFlow)
	go AddOperateLog(bookId, billFlow.BillId, "新增一笔"+OperateInfo[inOut], inOut, amount)
	return *mgresult.Success(billFlow)
}

func UpdateBillFlow(bookId, billId, billTypeId, billAccountId, customerId, remark, photos, billTime string, amount float64) mgresult.Result {
	var billFlow model.BillFlow
	err := config.Mgo.C(COLLECTION_BILL_FLOW).Find(&bson.M{"bookId": bookId, "billId": billId}).One(&billFlow)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	billFlow.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	if amount != 0.0 {
		billFlow.Amount = amount
	}
	if billTime != "" {
		billFlow.BillTime = billTime
	}
	if billTypeId != "" {
		billType := getBillType(billTypeId)
		billFlow.BillTypeId = billTypeId
		billFlow.BillType = billType.BillType
		billFlow.BillTypeIconId = billType.BillTypeIconId
	}
	if billAccountId != "" {
		billFlow.BillAccountId = billAccountId
		billAccount := getBillAccount(bookId, billAccountId)
		billFlow.BillAccount = billAccount.BillAccount
	}
	if customerId != "" {
		billFlow.BusinessCustomer = customerId
		businessCustomer := getBusinessCustomer(bookId, customerId)
		billFlow.BusinessCustomer = businessCustomer.BusinessCustomer
	}
	if remark != "" {
		billFlow.Remark = remark
	}
	if photos != "" && photos[:1] != "[" {
		var photoList []string
		utils.FromJSON(photos, &photoList)
		billFlow.Photos = photoList
	}
	err = config.Mgo.C(COLLECTION_BILL_FLOW).UpdateId(billFlow.Id, billFlow)
	if err != nil {
		logs.Error("MongoDB更新错误:{}", err.Error())
	}
	go AddOperateLog(bookId, billFlow.BillId, "修改一笔"+OperateInfo[billFlow.InOut], billFlow.InOut, billFlow.Amount)
	return GetBillFlow(billId)
}

func DeleteBillFlow(bookId, billId string) mgresult.Result {
	err := config.Mgo.C(COLLECTION_BILL_FLOW).Remove(&bson.M{"bookId": bookId, "billId": billId})
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(nil)
}

func GetBillFlow(billId string) mgresult.Result {
	var billFlow model.BillFlow
	err := config.Mgo.C(COLLECTION_BILL_FLOW).Find(&bson.M{"billId": billId}).One(&billFlow)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(billFlow)
}

func ListBillFlow(bookId, year, month, week, billTypeId, customerId string, inOut, page, size int) mgresult.Result {
	query := bson.M{"bookId": bookId}
	if billTypeId != "" {
		query["billTypeId"] = billTypeId
	}
	if customerId != "" {
		query["businessCustomerId"] = customerId
	}
	if inOut > 0 && inOut < 3 {
		query["inOut"] = inOut
	}
	if year == "" && month == "" {
		startDate := utils.DateFormat(time.Now(), "yyyy-MM") + "-01"
		query["billTime"] = bson.M{"$gt": startDate}
	} else if month != "" {
		startDate := month + "-01"
		s, _ := time.Parse("2006-01-02", startDate)
		e := s.AddDate(0, 1, 0)
		endDate := utils.DateFormat(e, "yyyy-MM-dd")
		query["billTime"] = bson.M{"$gt": startDate, "$lt": endDate}
	} else if year != "" && week == "" {
		startDate := year + "-01-01"
		endDate := year + "-12-31 23:59:59"
		query["billTime"] = bson.M{"$gt": startDate, "$lt": endDate}
	} else if year != "" && week != "" {
		y, _ := strconv.Atoi(year)
		w, _ := strconv.Atoi(week)
		s := isoweek.StartTime(y, w, time.UTC)
		startDate := utils.DateFormat(s, "yyyy-MM-dd")
		endDate := utils.DateFormat(s.AddDate(0, 0, 7), "yyyy-MM-dd")
		query["billTime"] = bson.M{"$gt": startDate, "$lt": endDate}
	}
	logs.Debug("查询条件:{}", query)
	billFlows := make([]model.BillFlow, 0)
	var err error
	if size > 0 {
		err = config.Mgo.C(COLLECTION_BILL_FLOW).Find(query).Sort("-billTime").Skip((page - 1) * size).Limit(size).All(&billFlows)
	} else {
		err = config.Mgo.C(COLLECTION_BILL_FLOW).Find(query).Sort("-billTime").All(&billFlows)
	}
	if err != nil {
		logs.Error("MongoDB查询错误:{}", err.Error())
	}
	return *mgresult.Success(billFlows)
}
