package types

type Workstation struct {
	WorkstationId   int    `json:"workstation_id"  db:"workstation_id"`
	EmployeeId      int    `json:"employee_id"  db:"employee_id"`
	WorkstationName string `json:"workstation_name"  db:"workstation_name"`
}
