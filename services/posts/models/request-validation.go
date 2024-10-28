package models

import (
	"fmt"
	"time"
)

func IsValidSortOrder(order *SortOrder) error {
	if order == nil ||
		*order == SortOrderAsc || *order == SortOrderDesc {
		return nil
	}
	return fmt.Errorf("invalid sort order: %s", *order)
}

func IsValidFilterOperator(operator FilterOperator) bool {
	switch operator {
	case OperatorEquals, OperatorNotEquals, OperatorGreaterThan,
		OperatorLessThan, OperatorContains, OperatorIn:
		return true
	default:
		return false
	}
}

func IsValidFieldOperatorCombination(field string, operator FilterOperator) bool {
	validOperators, fieldExists := AllowedFields[field]
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

func IsValidFilter(filter Filter) error {
	// Check if field is allowed
	if _, exists := AllowedFields[filter.Field]; !exists {
		return fmt.Errorf("invalid field: %s is not a filterable field", filter.Field)
	}

	// Check if operator is valid
	if !IsValidFilterOperator(filter.Operator) {
		return fmt.Errorf("invalid operator: %s", filter.Operator)
	}

	// Check if field-operator combination is valid
	if !IsValidFieldOperatorCombination(filter.Field, filter.Operator) {
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
