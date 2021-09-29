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

const COLLECTION_BILL_ACCRUAL = "BillAccrual"

func AddBillAccrual(bookId, billTypeId, customerId, remark, photos, billTime, excuteDate string, inOut int, amount float64) mgresult.Result {
	billAccrual := model.BillAccrual{
		Id:                 bson.NewObjectId(),
		BillId:             utils.GetUUIDString(),
		BookId:             bookId,
		BillTypeId:         billTypeId,
		InOut:              inOut,
		BillTime:           billTime,
		UpdateTime:         utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
		Amount:             amount,
		RemainAmount:       amount,
		BusinessCustomerId: customerId,
		Remark:             remark,
		ExcuteDate:         excuteDate,
	}
	if billTime == "" {
		billAccrual.BillTime = billAccrual.UpdateTime
	}
	var photoList []string
	if photos != "" && photos[:1] != "[" {
		utils.FromJSON(photos, &photoList)
	}
	billAccrual.Photos = photoList
	billType := getBillType(billTypeId)
	billAccrual.BillType = billType.BillType
	billAccrual.BillTypeIconId = billType.BillTypeIconId
	if customerId != "" {
		businessCustomer := getBusinessCustomer(bookId, customerId)
		billAccrual.BusinessCustomer = businessCustomer.BusinessCustomer
	}
	config.Mgo.C(COLLECTION_BILL_ACCRUAL).Insert(billAccrual)
	go AddOperateLog(bookId, billAccrual.BillId, "新增一笔"+OperateInfo[inOut], inOut, amount)
	return *mgresult.Success(billAccrual)
}

func UpdateBillAccrual(bookId, billId, billTypeId, customerId, remark, photos, billTime, excuteDate string, amount float64) mgresult.Result {
	var billAccrual model.BillAccrual
	err := config.Mgo.C(COLLECTION_BILL_ACCRUAL).Find(&bson.M{"bookId": bookId, "billId": billId}).One(&billAccrual)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	if billTime != "" {
		billAccrual.BillTime = billTime
	}
	billAccrual.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	if amount != 0 {
		if billAccrual.RemainAmount == 0 {
			return *mgresult.Error(-1, "账单已完成，不可改金额")
		} else {
			if amount < billAccrual.Amount-billAccrual.RemainAmount {
				return *mgresult.Error(-1, "账单金额不小于已完成金额")
			}
			diff := billAccrual.Amount - amount
			billAccrual.Amount = amount
			billAccrual.RemainAmount -= diff
		}
	}
	if billTypeId != "" && billAccrual.BillTypeId != billTypeId {
		billType := getBillType(billTypeId)
		billAccrual.BillType = billType.BillType
		billAccrual.BillTypeIconId = billType.BillTypeIconId
	}
	if customerId != "" && billAccrual.BusinessCustomerId != customerId {
		businessCustomer := getBusinessCustomer(bookId, customerId)
		billAccrual.BusinessCustomer = businessCustomer.BusinessCustomer
	}
	if remark != "" {
		billAccrual.Remark = remark
	}
	if excuteDate != "" {
		billAccrual.ExcuteDate = excuteDate
	}
	if photos != "" && photos[:1] != "[" {
		var photoList []string
		utils.FromJSON(photos, &photoList)
		billAccrual.Photos = photoList
	}
	config.Mgo.C(COLLECTION_BILL_ACCRUAL).UpdateId(billAccrual.Id, billAccrual)
	go AddOperateLog(bookId, billAccrual.BillId, "修改一笔"+OperateInfo[billAccrual.InOut], billAccrual.InOut, billAccrual.Amount)
	return GetBillAccrual(billId)
}

//从应收、应付中收款/付款操作
func AddBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, sync bool) mgresult.Result {
	var billAccrual model.BillAccrual
	err := config.Mgo.C(COLLECTION_BILL_ACCRUAL).Find(&bson.M{"bookId": bookId, "billId": billId}).One(&billAccrual)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	if billAccrual.RemainAmount == 0 {
		return *mgresult.Error(-1, "应收/应付账单状态已经完成")
	}
	if amount > billAccrual.RemainAmount {
		return *mgresult.Error(-1, "本次收款/付款金额大于账单未收/未付金额")
	}
	billRecord := model.BillRecord{
		Time:          utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
		Amount:        amount,
		BillAccountId: billAccountId,
		Remark:        remark,
	}
	billAccount := getBillAccount(bookId, billAccountId)
	billRecord.BillAccount = billAccount.BillAccount
	if photos != "" && photos[:1] != "[" {
		var photoList []string
		utils.FromJSON(photos, &photoList)
		billRecord.Photos = photoList
	}
	billAccrual.RemainAmount -= amount
	billAccrual.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	if billAccrual.BillRecords == nil || len(billAccrual.BillRecords) == 0 {
		billAccrual.BillRecords = []model.BillRecord{billRecord}
	} else {
		billAccrual.BillRecords = append(billAccrual.BillRecords, billRecord)
	}
	config.Mgo.C(COLLECTION_BILL_ACCRUAL).UpdateId(billAccrual.Id, billAccrual)
	if sync {
		AddBillFlow(bookId, billAccrual.BillTypeId, billAccountId, billAccrual.BusinessCustomerId, remark, photos, billAccrual.BillId, "", billAccrual.InOut-2, amount)
	}
	billOp := map[int]string{3: "的收款记录", 4: "的付款记录"}
	go AddOperateLog(bookId, billAccrual.BillId, "新增一笔"+OperateInfo[billAccrual.InOut]+billOp[billAccrual.InOut], billAccrual.InOut, amount)
	return GetBillAccrual(billId)
}

func DeleteBillRecord(bookId, billId string, billRecordIndex int) mgresult.Result {
	var billAccrual model.BillAccrual
	err := config.Mgo.C(COLLECTION_BILL_ACCRUAL).Find(&bson.M{"bookId": bookId, "billId": billId}).One(&billAccrual)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	if billAccrual.BillRecords == nil || len(billAccrual.BillRecords) == 0 {
		return *mgresult.Error(-1, "此账单无收付款记录")
	}
	if billRecordIndex >= len(billAccrual.BillRecords) {
		return *mgresult.Error(-1, "收付款记录号必须是0-"+strconv.Itoa(len(billAccrual.BillRecords)-1)+"之间")
	}
	billAccrual.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	billAccrual.RemainAmount += billAccrual.BillRecords[billRecordIndex].Amount
	billAccrual.BillRecords = append(billAccrual.BillRecords[:billRecordIndex], billAccrual.BillRecords[billRecordIndex+1:]...)
	config.Mgo.C(COLLECTION_BILL_ACCRUAL).UpdateId(billAccrual.Id, billAccrual)
	return GetBillAccrual(billId)
}

func UpdateBillRecord(bookId, billId, billAccountId, remark, photos string, amount float64, billRecordIndex int) mgresult.Result {
	var billAccrual model.BillAccrual
	err := config.Mgo.C(COLLECTION_BILL_ACCRUAL).Find(&bson.M{"bookId": bookId, "billId": billId}).One(&billAccrual)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	if billAccrual.BillRecords == nil || len(billAccrual.BillRecords) == 0 {
		return *mgresult.Error(-1, "此账单无收付款记录")
	}
	if billRecordIndex >= len(billAccrual.BillRecords) {
		return *mgresult.Error(-1, "收付款记录号必须是0-"+strconv.Itoa(len(billAccrual.BillRecords)-1)+"之间")
	}
	billAccrual.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	if billAccountId != "" && billAccountId != billAccrual.BillRecords[billRecordIndex].BillAccountId {
		billAccrual.BillRecords[billRecordIndex].BillAccountId = billAccountId
		billAccount := getBillAccount(bookId, billAccountId)
		billAccrual.BillRecords[billRecordIndex].BillAccount = billAccount.BillAccount
	}
	if remark != "" {
		billAccrual.BillRecords[billRecordIndex].Remark = remark
	}
	if photos != "" && photos[:1] != "[" {
		var photoList []string
		utils.FromJSON(photos, &photoList)
		billAccrual.BillRecords[billRecordIndex].Photos = photoList
	}
	if amount > 0.0 && amount != billAccrual.BillRecords[billRecordIndex].Amount {
		diff := billAccrual.BillRecords[billRecordIndex].Amount - amount
		billAccrual.BillRecords[billRecordIndex].Amount = amount
		billAccrual.RemainAmount += diff
	}
	billAccrual.BillRecords[billRecordIndex].Time = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	config.Mgo.C(COLLECTION_BILL_ACCRUAL).UpdateId(billAccrual.Id, billAccrual)
	billOp := map[int]string{3: "的收款记录", 4: "的付款记录"}
	go AddOperateLog(bookId, billAccrual.BillId, "修改一笔"+OperateInfo[billAccrual.InOut]+billOp[billAccrual.InOut], billAccrual.InOut, amount)
	return GetBillAccrual(billId)
}

func DeleteBillAccrual(bookId, billId string) mgresult.Result {
	err := config.Mgo.C(COLLECTION_BILL_ACCRUAL).Remove(&bson.M{"bookId": bookId, "billId": billId})
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(nil)
}

func GetBillAccrual(billId string) mgresult.Result {
	var billAccrual model.BillAccrual
	err := config.Mgo.C(COLLECTION_BILL_ACCRUAL).Find(&bson.M{"billId": billId}).One(&billAccrual)
	if err != nil {
		return *mgresult.Error(-1, err.Error())
	}
	return *mgresult.Success(billAccrual)
}

func ListBillAccrual(bookId, year, month, week, billTypeId, customerId string, excuteDate, inOut, hideFinished, page, size int) mgresult.Result {
	query := bson.M{"bookId": bookId}
	if billTypeId != "" {
		query["billTypeId"] = billTypeId
	}
	if customerId != "" {
		query["businessCustomerId"] = customerId
	}
	if inOut > 2 && inOut < 5 {
		query["inOut"] = inOut
	}
	if hideFinished == 1 {
		query["remainAmount"] = bson.M{"$gt": 0}
	}
	queryTime := "billTime"
	if excuteDate == 1 {
		queryTime = "excuteDate"
	}
	if year == "" && month == "" {
		startDate := utils.DateFormat(time.Now(), "yyyy-MM") + "-01"
		query[queryTime] = bson.M{"$gt": startDate}
	} else if month != "" {
		startDate := month + "-01"
		s, _ := time.Parse("2006-01-02", startDate)

		e := s.AddDate(0, 1, 0)
		endDate := utils.DateFormat(e, "yyyy-MM-dd")
		query[queryTime] = bson.M{"$gt": startDate, "$lt": endDate}
	} else if year != "" && week == "" {
		startDate := year + "-01-01"
		endDate := year + "-12-31 23:59:59"
		query[queryTime] = bson.M{"$gt": startDate, "$lt": endDate}
	} else if year != "" && week != "" {
		y, _ := strconv.Atoi(year)
		w, _ := strconv.Atoi(week)
		s := isoweek.StartTime(y, w, time.UTC)
		startDate := utils.DateFormat(s, "yyyy-MM-dd")
		endDate := utils.DateFormat(s.AddDate(0, 0, 7), "yyyy-MM-dd")
		query[queryTime] = bson.M{"$gt": startDate, "$lt": endDate}
	}
	logs.Debug("查询语句:{}", query)
	billAccruals := make([]model.BillAccrual, 0)
	if size > 0 {
		config.Mgo.C(COLLECTION_BILL_ACCRUAL).Find(query).Sort("-" + queryTime).Skip((page - 1) * size).Limit(size).All(&billAccruals)
	} else {
		config.Mgo.C(COLLECTION_BILL_ACCRUAL).Find(query).Sort("-" + queryTime).All(&billAccruals)
	}
	return *mgresult.Success(billAccruals)
}
