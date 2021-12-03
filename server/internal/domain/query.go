package domain

// PaginationQuery
type PaginationQuery struct {
	Skip  int64 `form:"skip"`
	Limit int64 `form:"limit"`
}
