package service

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/utils"
	"ququ.im/jizhangbao-data/model"
	"ququ.im/jizhangbao-data/search"
	"ququ.im/searchx"
	"time"
)

const (
	GOSEARCH_DATABASE = "jizhangbao"
	GOSEARCH_TABLE    = "business_customer"
)

func getBusinessCustomer(bookId, businessCustomerId string) model.BusinessCustomer {
	var businessCustomer model.BusinessCustomer
	query := map[string]interface{}{"bookId": bookId, "businessCustomerId": businessCustomerId}
	resp := search.SearchDocument(GOSEARCH_DATABASE, GOSEARCH_TABLE, nil, query, nil, nil, nil, nil, nil, 0, 1, 1)
	if resp.Status == 1 && resp.Data != nil {
		var payload searchx.Payload
		utils.FromJSON(utils.ToJSON(resp.Data), &payload)
		if payload.Totals >= 1 {
			utils.FromJSON(utils.ToJSON(payload.Docs[0]), &businessCustomer)
		}
	}
	return businessCustomer
}

func GetBusinessCustomer(bookId, businessCustomerId string) mgresult.Result {
	return *mgresult.Success(getBusinessCustomer(bookId, businessCustomerId))
}

func ListBusinessCustomer(bookId, customerName string) mgresult.Result {
	must := map[string]interface{}{"bookId": bookId}
	var r mgresult.Result
	if customerName != "" {
		wildcard := map[string]interface{}{"businessCustomer": customerName}
		r = search.SearchDocument(GOSEARCH_DATABASE, GOSEARCH_TABLE, nil, must, nil, nil, wildcard, nil, nil, 0, 0, 0)
	} else {
		r = search.SearchDocument(GOSEARCH_DATABASE, GOSEARCH_TABLE, nil, must, nil, nil, nil, nil, nil, 0, 0, 0)
	}
	if r.Status == 1 && r.Data != nil {
		var payload searchx.Payload
		utils.FromJSON(utils.ToJSON(r.Data), &payload)
		if payload.Totals >= 1 {
			return *mgresult.Success(payload.Docs)
		} else {
			return *mgresult.Error(0, "无此往来客户")
		}
	} else {
		return r
	}
}

func AddBusinessCustomer(bookId, customerName string) mgresult.Result {
	businessCustomer := model.BusinessCustomer{
		Id:                 utils.GetRandomHexString(16),
		BusinessCustomerId: utils.GetUUIDString(),
		BusinessCustomer:   customerName,
		BookId:             bookId,
		UpdateTime:         utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss"),
	}
	return search.AddDocument(GOSEARCH_DATABASE, GOSEARCH_TABLE, businessCustomer, []string{"businessCustomer"})
}

func UpdateBusinessCustomer(bookId, customerId, customerName string) mgresult.Result {
	businessCustomer := getBusinessCustomer(bookId, customerId)
	if businessCustomer.BusinessCustomerId == "" {
		return *mgresult.Error(-1, "无此往来客户")
	}
	return search.UpdateDocument(GOSEARCH_DATABASE, GOSEARCH_TABLE, businessCustomer.Id, map[string]interface{}{"businessCustomer": customerName, "updateTime": utils.DateFormat(time.Now(), "yyyy-MM-dd hh:mm:ss")})
}

func DeleteBusinessCustomer(bookId, customerId string) mgresult.Result {
	businessCustomer := getBusinessCustomer(bookId, customerId)
	if businessCustomer.BusinessCustomerId == "" {
		return *mgresult.Error(-1, "无此往来客户")
	}
	return search.DeleteDocument(GOSEARCH_DATABASE, GOSEARCH_TABLE, businessCustomer.Id)
}
