package business

import "fmt"

var searchAllWithNameLikeKeywordStmt = `
	MATCH (n) 
	WHERE n.name =~ $regex
	RETURN n, labels(n)
	SKIP $offset LIMIT $limit
`

var countAllWithNameLikeKeywordStmt = `
	MATCH (n) 
	WHERE n.name =~ $regex
	RETURN count(n)
`

func SearchAllWithNameLikeKeywoard(keyword string, page, limit int64) (interface{}, error) {
	searchStmt := searchAllWithNameLikeKeywordStmt
	searchParamsMap := make(map[string]interface{})
	searchParamsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)
	searchParamsMap["offset"] = (page - 1) * limit
	searchParamsMap["limit"] = limit

	countStmt := countAllWithNameLikeKeywordStmt
	countParamsMap := make(map[string]interface{})
	countParamsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)

	statements := []string{searchStmt, countStmt}
	paramsArr := []map[string]interface{}{searchParamsMap, countParamsMap}
	includeGraphs := []bool{false, false}
	return Neo4jMultiQuery(statements, paramsArr, includeGraphs)
}

var searchInLabelsStmtTmp = `
	MATCH (n)
	WHERE (%s) AND n.name =~ $regex
	RETURN n, labels(n)
	SKIP $offset
	LIMIT $limit
`

func SearchInLabelsWithNameLikeKeyword(keyword string, labels []string, page, limit int64) (interface{}, error) {
	subQuery := fmt.Sprintf("n:%s", labels[0])
	if len(labels) == 1 {
		// pass
	} else {
		for i, label := range labels {
			if i > 0 {
				subQuery += fmt.Sprintf(" OR n:%s", label)
			}
		}
	}

	searchInLabelsStmt := fmt.Sprintf(searchInLabelsStmtTmp, subQuery)

	paramsMap := make(map[string]interface{})
	paramsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)
	paramsMap["offset"] = (page - 1) * limit
	paramsMap["limit"] = limit
	return Neo4jSingleQuery(searchInLabelsStmt, paramsMap, false)
}
