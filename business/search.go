package business

import "fmt"

func SearchAllWithNameLikeKeywoard(keyword string) (interface{}, error) {
	statment := "MATCH (n) where n.name =~ $regex return id(n) as id, n.name as name, labels(n)[0] as label"
	params := map[string]interface{}{
		"regex": fmt.Sprintf(".*%s.*", keyword),
	}
	return QueryNeo4j(statment, params, false)
}
