package types

import "database/sql"

type Responsibility struct {
	ResponsibilityId   int    `json:"responsibility_id"  db:"responsibility_id"`
	ResponsibilityName string `json:"responsibility_name"  db:"responsibility_name"`
}

type ResponseResponsibility struct {
	Response sql.Result `json:"response"`
}
