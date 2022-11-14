package types

type Workstation struct {
	WorkstationId   int    `json:"workstation_id"  db:"workstation_id"`
	WorkstationName string `json:"workstation_name"  db:"workstation_name"`
	EmployeeId      int    `json:"employee_id"  db:"employee_id"`
}
