package types

// Employee структура, которая соответствует таблице Employee в БД
type Employee struct {
	EmployeeId    int    `json:"employee_id" db:"employee_id"`
	Username      string `json:"username" db:"username"`
	Password      string `json:"password" db:"password"`
	FirstName     string `json:"first_name" db:"first_name"`
	SecondName    string `json:"second_name" db:"second_name"`
	Position      int    `json:"position" db:"position"`
	SessionState  bool   `json:"session_state" db:"session_state"`
	Status        int    `json:"status" db:"status"`
	RefreshToken  string `json:"refresh_token" db:"refresh_token"`
	ExpiresAt     int    `json:"expires_at" db:"expires_at"`
	WorkstationId int    `json:"workstation_id" db:"workstation_id"`
	IsAdmin       bool   `json:"is_admin" db:"is_admin"`
}

type GetNewClientResponse struct {
	NumberTicket   int    `json:"number_ticket"`
	ServiceTicket  string `json:"service_ticket"`
	EmployeeStatus int    `json:"employee_status"`
	NumberQueue    int    `json:"number_queue"`
}

type GetEmployeeListsResponse struct {
	Data []Employee `json:"data"`
}

type ConfirmClientResponse struct {
	NumberQueue int `json:"number_queue"`
}

type EmployeeStatusResponse struct {
	EmployeeStatus int `json:"employee_status"`
}
