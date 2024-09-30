package database

import "gorm.io/gorm/clause"

func KeyColumnsToClause(keys []string) []clause.Column {
	result := make([]clause.Column, len(keys))

	for i, k := range keys {
		result[i] = clause.Column{Name: k}
	}

	return result
}
