package service

import (
	"github.com/maczh/gintool/mgresult"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"gopkg.in/mgo.v2/bson"
	"ququ.im/jizhangbao-data/model"
	"time"
)

const COLLECTION_BILL_ACCOUNT = "BillAccount"

func getBillAccount(bookId, billAccountId string) model.BillAccount {
	var billAccount model.BillAccount
	config.Mgo.C(COLLECTION_BILL_ACCOUNT).Find(&bson.M{"bookId": bookId, "billAccountId": billAccountId}).One(&billAccount)
	return billAccount
}

func listBillAccount(bookId string) []model.BillAccount {
	billAccounts := make([]model.BillAccount, 0)
	config.Mgo.C(COLLECTION_BILL_ACCOUNT).Find(&bson.M{"bookId": bookId}).All(&billAccounts)
	return billAccounts
}

func NewBillAccount(bookId string) mgresult.Result {
	billAccounts := listBillAccount(bookId)
	if billAccounts != nil && len(billAccounts) > 0 {
		return *mgresult.Success(billAccounts)
	}
	defaultAccounts := config.Redis.SMembers("jizhangbao:bill:account").Val()
	for _, accountName := range defaultAccounts {
		billAccount := model.BillAccount{
			Id:            bson.NewObjectId(),
			BookId:        bookId,
			BillAccountId: utils.GetUUIDString(),
			BillAccount:   accountName,
			UpdateTime:    utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
		}
		config.Mgo.C(COLLECTION_BILL_ACCOUNT).Insert(billAccount)
	}
	return *mgresult.Success(listBillAccount(bookId))
}

func AddBillAccount(bookId, billAccountName string) mgresult.Result {
	billAccount := model.BillAccount{
		Id:            bson.NewObjectId(),
		BookId:        bookId,
		BillAccountId: utils.GetUUIDString(),
		BillAccount:   billAccountName,
		UpdateTime:    utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
	}
	config.Mgo.C(COLLECTION_BILL_ACCOUNT).Insert(billAccount)
	return *mgresult.Success(billAccount)
}

func UpdateBillAccount(bookId, billAccountId, billAccountName string) mgresult.Result {
	billAccount := getBillAccount(bookId, billAccountId)
	if billAccount.BookId == "" {
		return *mgresult.Error(-1, "无此资金账户")
	}
	billAccount.BillAccount = billAccountName
	billAccount.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	config.Mgo.C(COLLECTION_BILL_ACCOUNT).UpdateId(billAccount.Id, billAccount)
	return *mgresult.Success(billAccount)
}

func DeleteBillAccount(bookId, billAccountId string) mgresult.Result {
	config.Mgo.C(COLLECTION_BILL_ACCOUNT).Remove(&bson.M{"bookId": bookId, "billAccountId": billAccountId})
	return *mgresult.Success(nil)
}

func ListBillAccount(bookId string) mgresult.Result {
	return *mgresult.Success(listBillAccount(bookId))
}

func GetBillAccount(bookId, billAccountId string) mgresult.Result {
	billAccount := getBillAccount(bookId, billAccountId)
	if billAccount.BookId == "" {
		return *mgresult.Error(-1, "无此资金账户")
	}
	return *mgresult.Success(billAccount)
}
