package business

import "fmt"

var searchAllWithNameLikeKeywordStmt = `
	MATCH (n) 
	WHERE n.name =~ $regex
	RETURN n, labels(n)
	SKIP $offset LIMIT $limit
`

var countSearchAllWithNameLikeKeywordStmt = `
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

	countStmt := countSearchAllWithNameLikeKeywordStmt
	countParamsMap := make(map[string]interface{})
	countParamsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)

	statements := []string{searchStmt, countStmt}
	paramsMapArr := []map[string]interface{}{searchParamsMap, countParamsMap}
	includeGraphs := []bool{false, false}
	return Neo4jMultiQuery(statements, paramsMapArr, includeGraphs)
}

var searchInLabelsStmtTempl = `
	MATCH (n)
	WHERE (%s) AND n.name =~ $regex
	RETURN n, labels(n)
	SKIP $offset
	LIMIT $limit
`

var countSearchInLabelsStmtTempl = `
	MATCH (n)
	WHERE (%s) AND n.name =~ $regex
	RETURN count(n)
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

	searchInLabelsStmt := fmt.Sprintf(searchInLabelsStmtTempl, subQuery)
	searchParamsMap := make(map[string]interface{})
	searchParamsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)
	searchParamsMap["offset"] = (page - 1) * limit
	searchParamsMap["limit"] = limit

	countInLabelsStmt := fmt.Sprintf(countSearchInLabelsStmtTempl, subQuery)
	countParamsMap := make(map[string]interface{})
	countParamsMap["regex"] = fmt.Sprintf(".*%s.*", keyword)

	statements := []string{searchInLabelsStmt, countInLabelsStmt}
	paramsMapArr := []map[string]interface{}{searchParamsMap, countParamsMap}
	includeGraphs := []bool{false, false}

	return Neo4jMultiQuery(statements, paramsMapArr, includeGraphs)
}
