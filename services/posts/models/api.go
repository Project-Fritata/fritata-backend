package models

import (
	userModels "github.com/Project-Fritata/fritata-backend/services/users/models"
)

// Sorting Order
type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

// Filters
const (
	OperatorEquals      string = "eq"
	OperatorNotEquals   string = "ne"
	OperatorGreaterThan string = "gt"
	OperatorLessThan    string = "lt"
	OperatorContains    string = "contains"
	OperatorIn          string = "in"
)

var AllowedFields = map[string][]string{
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
	Offset int `query:"offset"`
	Limit  int `query:"limit" default:"10"`

	SortOrder *SortOrder `query:"sort"`
	Filters   []string   `query:"filters" description:"<field>:<operator>:<value> - allowed fields: {created_at, content, media}, allowed operators: {eq, ne, gt, lt, contains, in}"`
}

type GetPostsRes struct {
	Post Post            `json:"post"`
	User userModels.User `json:"user"`
}

type CreatePostReq struct {
	Content string `json:"content"`
	Media   string `json:"media"`
}

type CreatePostRes struct {
	Message string `json:"message"`
}
