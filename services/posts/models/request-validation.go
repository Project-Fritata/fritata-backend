package models

import (
	"fmt"
	"time"
)

func IsValidSortOrder(order *SortOrder) error {
	if order == nil || *order == SortOrderAsc || *order == SortOrderDesc {
		return nil
	}
	return fmt.Errorf("invalid sort order: %s", *order)
}

func IsValidFilterOperator(operator string) bool {
	switch operator {
	case OperatorEquals, OperatorNotEquals, OperatorGreaterThan,
		OperatorLessThan, OperatorContains, OperatorIn:
		return true
	default:
		return false
	}
}

func IsValidFieldOperatorCombination(field string, operator string) bool {
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

func IsValidFilter(field string, operator string, value string) error {

	// Check if field is allowed
	if _, exists := AllowedFields[field]; !exists {
		return fmt.Errorf("invalid field: %s is not a filterable field", field)
	}

	// Check if operator is valid
	if !IsValidFilterOperator(operator) {
		return fmt.Errorf("invalid operator: %s", operator)
	}

	// Check if field-operator combination is valid
	if !IsValidFieldOperatorCombination(field, operator) {
		return fmt.Errorf("invalid operator %s for field %s", field, operator)
	}

	// Check if value is valid
	if field == "created_at" {
		if _, err := time.Parse("2006-01-02", value); err != nil {
			return fmt.Errorf("invalid value for field %s: %s", field, value)
		}
	}

	return nil
}
