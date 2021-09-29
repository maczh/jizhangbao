package service

import (
	"github.com/maczh/gintool/mgresult"
	config "github.com/maczh/mgconfig"
	"github.com/maczh/utils"
	"gopkg.in/mgo.v2/bson"
	"ququ.im/jizhangbao-data/model"
	"time"
)

const COLLECTION_BILL_BOOK = "BillBook"

func NewBillBook(shopId, userId, bookName string) mgresult.Result {
	billBook := model.BillBook{
		Id:         bson.NewObjectId(),
		ShopId:     shopId,
		UserId:     userId,
		BookName:   bookName,
		BookId:     utils.GetUUIDString(),
		UpdateTime: utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
	}
	config.Mgo.C(COLLECTION_BILL_BOOK).Insert(billBook)
	NewBillAccount(billBook.BookId)
	return *mgresult.Success(billBook)
}

func UpdateBillBook(bookId, bookName string) mgresult.Result {
	var billBook model.BillBook
	config.Mgo.C(COLLECTION_BILL_BOOK).Find(&bson.M{"bookId": bookId}).One(&billBook)
	if billBook.BookName == "" {
		return *mgresult.Error(-1, "无此账本")
	}
	billBook.BookName = bookName
	billBook.UpdateTime = utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")
	config.Mgo.C(COLLECTION_BILL_BOOK).UpdateId(billBook.Id, billBook)
	return *mgresult.Success(billBook)
}

func ListBillBooks(shopId, userId string) mgresult.Result {
	billBooks := make([]model.BillBook, 0)
	config.Mgo.C(COLLECTION_BILL_BOOK).Find(&bson.M{"shopId": shopId, "userId": userId}).All(&billBooks)
	if len(billBooks) == 0 {
		return *mgresult.Error(-1, "此用户无账本")
	}
	return *mgresult.Success(billBooks)
}

func GetBillBook(bookId string) mgresult.Result {
	var billBook model.BillBook
	config.Mgo.C(COLLECTION_BILL_BOOK).Find(&bson.M{"bookId": bookId}).One(&billBook)
	if billBook.BookId == "" {
		return *mgresult.Error(-1, "此用户无账本")
	}
	return *mgresult.Success(billBook)
}
