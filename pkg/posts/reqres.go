package posts

import (
	"fmt"
	"time"

	"github.com/Project-Fritata/fritata-backend/internal"
)

// Sorting Order
type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

func isValidSortOrder(order *SortOrder) error {
	if order == nil ||
		*order == SortOrderAsc || *order == SortOrderDesc {
		return nil
	}
	return fmt.Errorf("invalid sort order: %s", *order)
}

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

var allowedFields = map[string][]FilterOperator{
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

func isValidFilterOperator(operator FilterOperator) bool {
	switch operator {
	case OperatorEquals, OperatorNotEquals, OperatorGreaterThan,
		OperatorLessThan, OperatorContains, OperatorIn:
		return true
	default:
		return false
	}
}
func isValidFieldOperatorCombination(field string, operator FilterOperator) bool {
	validOperators, fieldExists := allowedFields[field]
	if !fieldExists {
		return false
	}

	for _, validOp := range validOperators {
		if operator == validOp {
			return true
		}
	}
	return false
}
func isValidFilter(filter Filter) error {
	// Check if field is allowed
	if _, exists := allowedFields[filter.Field]; !exists {
		return fmt.Errorf("invalid field: %s is not a filterable field", filter.Field)
	}

	// Check if operator is valid
	if !isValidFilterOperator(filter.Operator) {
		return fmt.Errorf("invalid operator: %s", filter.Operator)
	}

	// Check if field-operator combination is valid
	if !isValidFieldOperatorCombination(filter.Field, filter.Operator) {
		return fmt.Errorf("invalid operator %s for field %s", filter.Operator, filter.Field)
	}

	// Check if value is valid
	if filter.Field == "created_at" {
		if _, err := time.Parse("2006-01-02", filter.Value); err != nil {
			return fmt.Errorf("invalid value for field %s: %s", filter.Field, filter.Value)
		}
	}

	return nil
}

type GetReq struct {
	Offset int `query:"offset" validate:"min=0"`
	Limit  int `query:"limit" validate:"required,min=1,max=100"`

	SortOrder *SortOrder `query:"sort" validate:"omitempty,oneof=asc desc"`

	Filters []Filter `query:"filters"`
}

type GetRes struct {
	Post internal.Post `json:"post"`
	User internal.User `json:"user"`
}

type PostReq struct {
	Content string `json:"content"`
	Media   string `json:"media"`
}
