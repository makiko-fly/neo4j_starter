package types

type Neo4jQueryResponse struct {
	Results []Neo4jQueryResult `json:"results"`
	Errors  []Neo4jQueryErr    `json:"errors"`
}

type Neo4jQueryResult struct {
	Columns []string               `json:"columns"`
	DataArr []Neo4jQueryResultData `json:"data"`
}

type Neo4jQueryErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (self *Neo4jQueryErr) String() string {
	return self.Code + "  " + self.Message
}

type Neo4jQueryResultData struct {
	Rows []interface{} `json:"row"`
}
