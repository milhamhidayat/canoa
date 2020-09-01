package entity

// QueryFilter is used to filter resources
type QueryFilter struct {
	Num     int
	Cursor  string
	Keyword string
	IDs     []string
}
