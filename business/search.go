package business

import "fmt"

func SearchAllWithNameLikeKeywoard(keyword string) (interface{}, error) {
	statment := "MATCH (n) where n.name =~ $regex return n, labels(n)"
	params := map[string]interface{}{
		"regex": fmt.Sprintf(".*%s.*", keyword),
	}
	return QueryNeo4j(statment, params, false)
}
