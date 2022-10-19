package types

type Responsibility struct {
	ResponsibilityId   int    `json:"responsibility_id"  db:"responsibility_id"`
	ResponsibilityName string `json:"responsibility_name"  db:"responsibility_name"`
}
