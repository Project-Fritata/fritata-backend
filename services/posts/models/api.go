package models

import (
	usermodel "github.com/Project-Fritata/fritata-backend/services/users/models"
)

// Sorting Order
type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

// Filters
type Filter struct {
	Field    string         `query:"field"`
	Operator FilterOperator `query:"operator"`
	Value    string         `query:"value"`
}

type FilterOperator string

const (
	OperatorEquals      FilterOperator = "eq"
	OperatorNotEquals   FilterOperator = "ne"
	OperatorGreaterThan FilterOperator = "gt"
	OperatorLessThan    FilterOperator = "lt"
	OperatorContains    FilterOperator = "contains"
	OperatorIn          FilterOperator = "in"
)

var AllowedFields = map[string][]FilterOperator{
	"created_at": {
		OperatorEquals,
		OperatorNotEquals,
		OperatorGreaterThan,
		OperatorLessThan,
	},
	"content": {
		OperatorEquals,
		OperatorNotEquals,
		OperatorContains,
	},
	"media": {
		OperatorEquals,
		OperatorNotEquals,
		OperatorContains,
	},
}

type GetPostsReq struct {
	Offset int `query:"offset" validate:"min=0"`
	Limit  int `query:"limit" validate:"required,min=1,max=100"`

	SortOrder *SortOrder `query:"sort" validate:"omitempty,oneof=asc desc"`

	Filters []Filter `query:"filters"`
}

type GetPostsRes struct {
	Post Post           `json:"post"`
	User usermodel.User `json:"user"`
}

type CreatePostReq struct {
	Content string `json:"content"`
	Media   string `json:"media"`
}

type CreatePostRes struct {
	Message string `json:"message"`
}
