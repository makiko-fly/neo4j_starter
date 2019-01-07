package business

func SearchAllWithNameLikeKeywoard(keyword string) (interface{}, error) {
	statment := "MATCH (n) where n.name =~ '.*.铁*' return n"
	return QueryNeo4j(statment, nil)
}
