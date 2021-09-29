package search

import (
	"github.com/maczh/gintool/mgresult"
	"github.com/maczh/logs"
	"github.com/maczh/mgcall"
	"github.com/maczh/utils"
	"strconv"
	"unicode"
)

const (
	SERVICE_GOSEARCH             = "gosearch"
	URI_ADD_DOCUMENT_BATCH       = "/gosearch/add/batch"
	URI_ADD_DOCUMENT             = "/gosearch/add"
	URI_DEL_DOCUMENT             = "/gosearch/del"
	URI_DEL_DOCUMENT_BY_QUERY    = "/gosearch/del/query"
	URI_UPDATE_DOCUMENT          = "/gosearch/update"
	URI_UPDATE_DOCUMENT_BY_QUERY = "/gosearch/update/query"
	URI_DEL_TABLE                = "/gosearch/del/table"
	URI_DEL_DATABASE             = "/gosearch/del/database"
	URI_SEARCH                   = "/gosearch/query"
	URI_INCR_SUGGEST             = "/gosearch/suggest/incr"
	URI_DEL_SUGGEST              = "/gosearch/suggest/del"
	URI_SUGGEST                  = "/gosearch/suggest"
)

func AddDocument(database, table string, doc interface{}, searchFields []string) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["doc"] = utils.ToJSON(doc)
	if searchFields != nil && len(searchFields) > 0 {
		params["searchFields"] = utils.ToJSON(searchFields)
	}
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_ADD_DOCUMENT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_ADD_DOCUMENT, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func AddDocuments(database, table string, docs interface{}, searchFields []string) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["docs"] = utils.ToJSON(docs)
	if searchFields != nil && len(searchFields) > 0 {
		params["searchFields"] = utils.ToJSON(searchFields)
	}
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_ADD_DOCUMENT_BATCH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_ADD_DOCUMENT_BATCH, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func DeleteDocument(database, table, id string) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["id"] = id
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_DEL_DOCUMENT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_DEL_DOCUMENT, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func DeleteDocumentByQuery(database, table string, query map[string]interface{}) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["query"] = utils.ToJSON(query)
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_DEL_DOCUMENT_BY_QUERY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_DEL_DOCUMENT_BY_QUERY, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func UpdateDocument(database, table, id string, update map[string]interface{}) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["id"] = id
	params["update"] = utils.ToJSON(update)
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_UPDATE_DOCUMENT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_UPDATE_DOCUMENT, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func UpdateDocumentByQuery(database, table string, query map[string]interface{}, update map[string]interface{}) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["query"] = utils.ToJSON(query)
	params["update"] = utils.ToJSON(update)
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_UPDATE_DOCUMENT_BY_QUERY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_UPDATE_DOCUMENT_BY_QUERY, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func DeleteTable(database, table string) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_DEL_TABLE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_DEL_TABLE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func DeleteDatabase(database string) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_DEL_DATABASE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_DEL_DATABASE, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func SearchDocument(database, table string, should, must, not, shouldWildcard, mustWildcard, notWildcard map[string]interface{}, sort []string, offset, size, almost int) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["should"] = utils.ToJSON(should)
	params["must"] = utils.ToJSON(must)
	params["not"] = utils.ToJSON(not)
	params["shouldWildcard"] = utils.ToJSON(shouldWildcard)
	params["mustWildcard"] = utils.ToJSON(mustWildcard)
	params["notWildcard"] = utils.ToJSON(notWildcard)
	params["sort"] = utils.ToJSON(sort)
	params["offset"] = strconv.Itoa(offset)
	params["size"] = strconv.Itoa(size)
	params["almost"] = strconv.Itoa(almost)
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_SEARCH, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_SEARCH, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func IncrSuggest(database, table, keyword, customize string, incr int) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["keyword"] = keyword
	if customize != "" {
		params["customize"] = customize
	}
	if incr > 1 {
		params["incr"] = strconv.Itoa(incr)
	}
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_INCR_SUGGEST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_INCR_SUGGEST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func DeleteSuggest(database, table, customize, keyword string) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["keyword"] = keyword
	if customize != "" {
		params["customize"] = customize
	}
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_DEL_SUGGEST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_DEL_SUGGEST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	return callResult
}

func ListSuggests(database, table, prefix, customize string, size int) mgresult.Result {
	params := make(map[string]string)
	params["database"] = database
	params["table"] = table
	params["prefix"] = prefix
	if customize != "" {
		params["customize"] = customize
		params["size"] = "100"
	} else {
		if size > 0 {
			params["size"] = strconv.Itoa(size)
		} else {
			params["size"] = "10"
		}
	}
	res, err := mgcall.Call(SERVICE_GOSEARCH, URI_SUGGEST, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", SERVICE_GOSEARCH, URI_SUGGEST, err.Error())
		return *mgresult.Error(-1, err.Error())
	}
	var callResult mgresult.Result
	utils.FromJSON(res, &callResult)
	if customize != "" {
		suggests := callResult.Data.([]interface{})
		realSuggests := make([]string, 0)
		if suggests != nil && len(suggests) > 0 {
			for _, suggest := range suggests {
				if isContains(suggest.(string), prefix) {
					realSuggests = append(realSuggests, suggest.(string))
					if len(realSuggests) == size {
						break
					}
				}
			}
		}
		return *mgresult.Success(realSuggests)
	} else {
		return callResult
	}
}

func isContains(suggest, prefix string) bool {
	suggestRune := []rune(suggest)
	for i, r := range []rune(prefix) {
		if unicode.Is(unicode.Han, r) {
			if r != suggestRune[i] {
				return false
			}
		}
	}
	return true
}
