package business

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitlab.wallstcn.com/matrix/xgbkb/g"
	"gitlab.wallstcn.com/matrix/xgbkb/std"
	"gitlab.wallstcn.com/matrix/xgbkb/std/logger"
	"gitlab.wallstcn.com/matrix/xgbkb/std/redislogger"
	"gitlab.wallstcn.com/matrix/xgbkb/types"
)

func InitNeo4j() {
	// add property exist constraint to Product's `name` property
	AssertQueryNeo4j("CREATE CONSTRAINT ON (p:Product) ASSERT exists(p.name)", nil)
	// add unique constraint to Product's `name` property
	AssertQueryNeo4j("CREATE CONSTRAINT ON (p:Product) ASSERT p.name IS UNIQUE", nil)

	// add property exist constraint to Company's properties
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Company) ASSERT exists(c.name)", nil)
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Company) ASSERT exists(c.nameAbbr)", nil)
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Company) ASSERT exists(c.code)", nil)
	// add unique constraint to Company's properties
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Company) ASSERT c.name IS UNIQUE", nil)
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Company) ASSERT c.nameAbbr IS UNIQUE", nil)
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Company) ASSERT c.code IS UNIQUE", nil)

	// add property exist constraint to Stock's properties
	AssertQueryNeo4j("CREATE CONSTRAINT ON (s:Stock) ASSERT exists(s.name)", nil)
	AssertQueryNeo4j("CREATE CONSTRAINT ON (s:Stock) ASSERT exists(s.symbol)", nil)
	// add unique constraint to Stock's properties
	AssertQueryNeo4j("CREATE CONSTRAINT ON (s:Stock) ASSERT s.symbol IS UNIQUE", nil)

	// add property exist constraint to Chain's properties
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Chain) ASSERT exists(c.name)", nil)
	// add unique constraint to Chain's properties
	AssertQueryNeo4j("CREATE CONSTRAINT ON (c:Chain) ASSERT c.name IS UNIQUE", nil)
}

// assert the statement can be executed successfully, if any err occurs, panic
func AssertQueryNeo4j(stmt string, params map[string]interface{}) {
	data, err := QueryNeo4j(stmt, params, false)
	if err != nil {
		redislogger.Errf("InitNeo4j, stmt: %s, fails to execute, err: %v", stmt, err)
		panic(err)
	}
	resp, err := parseNeo4jJsonResp(data)
	if err != nil {
		redislogger.Errf("InitNeo4j, stmt: %s, fails to parse json resp: %s, err: %v", stmt, string(data), err)
		panic(err)
	}
	if len(resp.Errors) > 0 {
		redislogger.Errf("InitNeo4j, stmt: %s, neo4j indicates error: %v", stmt, resp.Errors[0])
		panic(resp.Errors[0].String())
	}
}

// single query
func QueryNeo4j(statement string, params map[string]interface{}, includeGraphData bool) ([]byte, error) {
	parametersStr := "{}"
	if len(params) > 0 {
		bytes, err := json.Marshal(params)
		if err != nil {
			redislogger.Errf("QueryNeo4j, fails to marshal params: %v, err: %v", params, err)
			return nil, err
		}
		parametersStr = string(bytes)
	}
	graphData := ""
	if includeGraphData {
		graphData = `,"graph"`
	}
	reqBodyStr := fmt.Sprintf(singleStatementTempl, EscapeStmt(statement), parametersStr, graphData)
	if byteArr, err := callNeo4jHttpApi("/db/data/transaction/commit", reqBodyStr); err != nil {
		return nil, err
	} else {
		if resp, err := parseNeo4jJsonResp(byteArr); err != nil {
			return nil, err
		} else {
			if len(resp.Errors) > 0 {
				return nil, std.NewNeo4jQueryErr(resp.Errors[0].String())
			} else {
				return byteArr, nil
			}
		}
	}
}

var singleStatementTempl = `
{
	"statements" : [ {
	  "statement" : "%s",
	  "parameters" : %s,
	  "resultDataContents" : [ "row" %s ]
	} ]
}
`

func callNeo4jHttpApi(path, bodyStr string) ([]byte, error) {
	urlStr := fmt.Sprintf("http://%s:%d%s", g.SysConf.Neo4jDb.Addr, g.SysConf.Neo4jDb.HttpPort, path)
	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer([]byte(bodyStr)))
	if err != nil {
		redislogger.Errf("callNeo4jHttpApi, create request fails, err: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	encodedAuthStr := encodeNeo4jUserNameAndPassword(g.SysConf.Neo4jDb.UserName, g.SysConf.Neo4jDb.Password)
	// logger.Infof("===> encodedAuthStr: %s", encodedAuthStr)
	req.Header.Set("Authorization", "Basic "+encodedAuthStr)
	logger.Infof("=== calling neo4j HTTP API with statements: %s", bodyStr)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		redislogger.Errf("callNeo4jHttpApi, http POST fails, err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		errStr := fmt.Sprintf("callNeo4jHttpApi, http POST returns code other than 200, code: %d", resp.StatusCode)
		redislogger.Errf(errStr)
		return nil, errors.New(errStr)
	}
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errStr := fmt.Sprintf("callNeo4jHttpApi, ioutil.ReadAll fails, err: %v", err)
		redislogger.Errf(errStr)
		return nil, errors.New(errStr)
	}
	return respBodyBytes, nil
}

func parseNeo4jJsonResp(data []byte) (*types.Neo4jQueryResponse, error) {
	var resp types.Neo4jQueryResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	} else {
		// TODO... make sure the results are in the correct format
		return &resp, nil
	}
}

func encodeNeo4jUserNameAndPassword(userName, password string) string {
	strToEncode := userName + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(strToEncode))
}

func EscapeStmt(stmt string) string {
	stmt = strings.Replace(stmt, "\n", " ", -1)
	stmt = strings.Replace(stmt, "\t", " ", -1)
	return stmt
}
