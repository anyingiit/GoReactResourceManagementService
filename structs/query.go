package structs

type StanderedQuery struct {
	PaginationPage    int    `form:"pagination_page,default=1"`
	PaginationPerPage int    `form:"pagination_per_page,default=10"`
	SortField         string `form:"sort_field,default=id"`
	SortOrder         string `form:"sort_order,default=asc"`
	FilterJson        string `form:"filter_json,default={}"`
}

type StanderedQueryParsed struct {
	PaginationPage    int
	PaginationPerPage int
	SortField         string
	SortOrder         string
	FilterJson        map[string]interface{}
}
