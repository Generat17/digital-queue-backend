package types

import "database/sql"

type Workstation struct {
	WorkstationId          int      `json:"workstation_id"  db:"workstation_id"`
	WorkstationName        string   `json:"workstation_name"  db:"workstation_name"`
	EmployeeId             int      `json:"employee_id"  db:"employee_id"`
	ResponsibilityList     []int    `json:"responsibility_list"`
	ResponsibilityListName []string `json:"responsibility_list_name"`
}

type WorkstationResponsibility struct {
	WorkstationId      int    `json:"workstation_id"  db:"workstation_id"`
	ResponsibilityId   int    `json:"responsibility_id"  db:"responsibility_id"`
	ResponsibilityName string `json:"responsibility_name"  db:"responsibility_name"`
}

type ResponseWorkstation struct {
	Response sql.Result `json:"response"`
}
